package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Find the unique number", func() {
	It("should work for some basic cases", func() {
		Expect(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
		Expect(FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})
})
