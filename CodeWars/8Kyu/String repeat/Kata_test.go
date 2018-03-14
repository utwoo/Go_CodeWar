package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("should repeat correctly", func() {
		Expect(Kata(4, "a")).To(Equal("aaaa"))
		Expect(Kata(3, "hello ")).To(Equal("hello hello hello "))
		Expect(Kata(2, "abc")).To(Equal("abcabc"))
	})
})
