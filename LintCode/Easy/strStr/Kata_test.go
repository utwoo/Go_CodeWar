package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("strStr", func() {
		Expect(strStr("abcdabcdefg", "bcd")).To(Equal(1))
		Expect(strStr("abcdabcdefg", "def")).To(Equal(7))
	})
})
