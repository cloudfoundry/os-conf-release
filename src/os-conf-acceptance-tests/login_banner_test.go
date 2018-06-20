package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("LoginBanner", func() {
	It("sets the login banner", func() {
		session := boshSSH("os-conf/0", "exit 0")
		Eventually(session, 30*time.Second).Should(gbytes.Say("Jim was here."))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
