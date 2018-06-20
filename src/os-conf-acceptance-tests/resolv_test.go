package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Resolv", func() {
	It("allows users to set options and serach domain", func() {
		session := boshSSH("os-conf/0", "sudo cat /etc/resolv.conf")
		Eventually(session, 30*time.Second).Should(gbytes.Say("search pivotal.io"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("options rotate"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("options timeout:1"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
