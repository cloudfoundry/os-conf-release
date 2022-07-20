package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Modprobe", func() {
	It("enables the kernel modules specified", func() {
		session := boshSSH("os-conf/0", "lsmod")
		Eventually(session, 30*time.Second).Should(gbytes.Say("lp"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
