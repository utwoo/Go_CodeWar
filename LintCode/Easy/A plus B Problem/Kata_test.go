package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("A + B Problem", func() {
		Expect(aplusb(1, 1)).To(Equal(2))
	})
})
