package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("HardenSshd", func() {
	It("hardens the sshd", func() {
		By("disabeling port forwarding", func() {
			session := boshSSH("os-conf/0", "sudo sshd -T | sort")
			Eventually(session, 30*time.Second).Should(gbytes.Say("allowstreamlocalforwarding no"))
			Eventually(session, 30*time.Second).Should(gbytes.Say("allowtcpforwarding no"))
			Eventually(session, 30*time.Second).Should(gbytes.Say("gatewayports no"))
			Eventually(session, 30*time.Second).Should(gbytes.Say("permittunnel no"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})
	})
})
