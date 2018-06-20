package os_conf_acceptance_tests_test

import (
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("DisableAgent", func() {
	XIt("kills the agent after a timeout", func() {
		Eventually(func() string {
			cmd := exec.Command(
				boshBinaryPath,
				"-d",
				deploymentName,
				"vms",
			)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))

			return string(session.Out.Contents())
		}, 60*time.Second).Should(MatchRegexp("unresponsive-agent/[\\w-]+\\s+unresponsive agent"))
	})
})
