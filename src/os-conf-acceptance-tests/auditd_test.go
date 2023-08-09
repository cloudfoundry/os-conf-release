package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("AuditD", func() {
	It("allows users to set auditd rules", func() {
		session := boshSSH("os-conf/0", "sudo grep auditd_test_rule /etc/audit/rules.d/audit.rules")
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))
	})
})
