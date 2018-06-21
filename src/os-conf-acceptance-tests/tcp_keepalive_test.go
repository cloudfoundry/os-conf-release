package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("TCPKeepAlive", func() {
	It("allows users to set tcp keepalive properties", func() {
		session := boshSSH("os-conf/0", "sudo cat /etc/sysctl.d/70-bosh-tcp-keepalive.conf")
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_time=10"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_intvl=11"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_probes=12"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("os-conf/0", "sudo sysctl -a")
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_time = 10"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_intvl = 11"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv4.tcp_keepalive_probes = 12"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
