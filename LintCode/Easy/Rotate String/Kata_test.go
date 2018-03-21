package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("Rotate String", func() {
		str := "abcdefg"
		rotateString(&str, 10)
		Expect(str).Should(Equal("efgabcd"))
	})
})
