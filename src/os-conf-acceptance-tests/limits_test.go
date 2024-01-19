package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Limits", func() {
	BeforeEach(func() {
		if boshStemcellOS == "ubuntu-trusty" {
			Skip("Trusty Stemcells are not supported.")
		}
	})

	Context("when limits are configured", func() {
		It("sets the limits for the monit process", func() {
			session := boshSSH("os-conf/0", "pid=$(ps -e | grep monit | awk '{print $1}'); cat /proc/$pid/limits | grep 'Max open files' | awk '{print $4}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("60000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))

			session = boshSSH("os-conf/0", "pid=$(ps -e | grep monit | awk '{print $1}'); cat /proc/$pid/limits | grep 'Max open files' | awk '{print $5}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("100000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		It("sets the limits for the parent process", func() {
			session := boshSSH("os-conf/0", "pid=$(ps -e | grep monit | awk '{print $1}'); pid_parent=$(ps -o ppid:1= -p $pid); cat /proc/$pid_parent/limits | grep 'Max open files' | awk '{print $4}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("60000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))

			session = boshSSH("os-conf/0", "pid=$(ps -e | grep monit | awk '{print $1}'); pid_parent=$(ps -o ppid:1= -p $pid); cat /proc/$pid_parent/limits | grep 'Max open files' | awk '{print $5}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("100000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		It("sets the limits for the systemd process", func() {
			session := boshSSH("os-conf/0", "systemctl show | grep 'DefaultLimitNOFILESoft=' | awk -F= '{print $2}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("60000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))

			session = boshSSH("os-conf/0", "systemctl show | grep 'DefaultLimitNOFILE=' | awk -F= '{print $2}'")
			Eventually(session, 30*time.Second).Should(gbytes.Say("100000"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})
	})
})
