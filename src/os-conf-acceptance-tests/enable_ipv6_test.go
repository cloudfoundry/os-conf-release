package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Enable IPV6", func() {
	It("enables ipv6 on a VM", func() {
		session := boshSSH("os-conf/0", "sudo sysctl --system")
		Eventually(session, 30*time.Second).Should(gbytes.Say("\\* Applying /etc/sysctl\\.d/10-ipv6-privacy\\.conf \\.\\.\\."))
		Eventually(session, 30*time.Second).Should(gbytes.Say("\\* Applying /etc/sysctl\\.d/61-bosh-enable_ipv6-release\\.conf \\.\\.\\."))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("os-conf/0", "sudo sysctl -a")
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.all.use_tempaddr = 0"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.default.use_tempaddr = 0"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.all.accept_ra=1"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.default.accept_ra=1"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.all.disable_ipv6=0"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.default.disable_ipv6=0"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.default.accept_redirects=1"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.conf.all.accept_redirects=1"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("net.ipv6.route.flush=0"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
