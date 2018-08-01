package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("UserAdd", func() {
	It("allows users to add users to the VM", func() {
		By("adding the users to /etc/passwd", func() {
			session := boshSSH("os-conf/0", "sudo cat /etc/passwd")
			Eventually(session, 30*time.Second).Should(gbytes.Say(`test-user-password:x:\d+:\d+::/home/test-user-password:/bin/rbash`))
			Eventually(session, 30*time.Second).Should(gbytes.Say(`test-user-key:x:\d+:\d+::/home/test-user-key:/bin/bash`))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		By("adding a password for the password user", func() {
			session := boshSSH("os-conf/0", "sudo cat /etc/shadow")
			Eventually(session, 30*time.Second).Should(gbytes.Say(`test-user-password:\$6\$kMBogqsbx\$70Y2m/mwYR8vKZqR9RD2UUPoWz8mJoBiH8IAbvH2v6LCjxJgB3kDtwR8QttqtI/WSqCsFy4qXZaKPM64sZMwK\.:\d+:1:99999:7:::`))
			Eventually(session, 30*time.Second).Should(gbytes.Say(`test-user-key::\d+:1:99999:7:::`))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		By("adding an authorized key for the key user", func() {
			session := boshSSH("os-conf/0", "sudo cat /home/test-user-key/.ssh/authorized_keys")
			Eventually(session, 30*time.Second).Should(gbytes.Say("test-user-public-key"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))

			session = boshSSH("os-conf/0", "sudo ls -larth /home/test-user-password/.ssh")
			Consistently(session, 10*time.Second).ShouldNot(gbytes.Say("authorized_keys"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		By("adding them to the bosh_sshers group", func() {
			session := boshSSH("os-conf/0", "sudo grep bosh_sshers /etc/group")
			Eventually(session, 30*time.Second).Should(gbytes.Say("test-user-password,test-user-key"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})

		By("adding the apropriate user to the bosh_sudoers group", func() {
			session := boshSSH("os-conf/0", "sudo grep bosh_sudoers /etc/group")
			Consistently(session, 10*time.Second).ShouldNot(gbytes.Say("test-user-password"))
			Eventually(session, 30*time.Second).Should(gbytes.Say("test-user-key"))
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
		})
	})
})
