package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("Trailing Zeros", func() {
		Expect(trailingZeros(11)).To(Equal(int64(2)))
		Expect(trailingZeros(1333)).To(Equal(int64(331)))
	})
})
