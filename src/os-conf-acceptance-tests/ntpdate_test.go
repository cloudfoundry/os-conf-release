package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("NTPDate", func() {
	BeforeEach(func() {
		if boshStemcellOS != "ubuntu-trusty" {
			Skip("Only Trusty Stemcells use ntpdate")
		}
	})

	It("allows using an unprivileged_port", func() {
		session := boshSSH("os-conf/0", "sudo cat /var/vcap/bosh/bin/sync-time")
		Eventually(session, 30*time.Second).Should(gbytes.Say("/usr/sbin/ntpdate -u"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
