package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Sysctl", func() {
	It("allows users to set arbitrary systcl values", func() {
		session := boshSSH("os-conf/0", "sudo cat /etc/sysctl.d/72-bosh-os-conf-sysctl.conf")
		Eventually(session, 30*time.Second).Should(gbytes.Say("vm.swappiness=10"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.core.somaxconn=1024"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("os-conf/0", "sudo sysctl vm.swappiness")
		Eventually(session, 30*time.Second).Should(gbytes.Say("vm.swappiness = 10"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("os-conf/0", "sudo sysctl net.core.somaxconn")
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.core.somaxconn = 1024"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
