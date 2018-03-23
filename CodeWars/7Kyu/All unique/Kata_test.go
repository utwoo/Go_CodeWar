package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("All unique", func() {
		Expect(HasUniqueChar("  nAa")).To(Equal(false))
		Expect(HasUniqueChar("abcde")).To(Equal(true))
		Expect(HasUniqueChar("++-")).To(Equal(false))
		Expect(HasUniqueChar("AaBbC")).To(Equal(true))
	})
})
