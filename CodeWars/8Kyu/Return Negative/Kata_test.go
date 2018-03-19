package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Return Negative", func() {
	It("test", func() {
		Expect(Kata(42)).To(Equal(-42))
		Expect(Kata(-9)).To(Equal(-9))
		Expect(Kata(0)).To(Equal(0))
	})
})
