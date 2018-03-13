package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Multiply", func() {
	It("test", func() {
		Expect(Kata(1, 1)).To(Equal(1))
	})
})
