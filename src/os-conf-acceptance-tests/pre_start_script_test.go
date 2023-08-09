package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("PreStartScript", func() {
	It("allows users to execute an arbitrary script in pre-start", func() {
		session := boshSSH("os-conf/0", "sudo cat /var/vcap/sys/log/pre-start-script/stdout.log")
		Eventually(session, 30*time.Second).Should(gbytes.Say("Arbitrary pre start script executed."))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
