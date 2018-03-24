package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cat Kata, Part 1", func() {
	It("should test that the solution returns the correct value", func() {
		// Only one cat is in the yard
		Expect(PeacefulYard([]string{
			"------------",
			"------------",
			"-L----------",
			"------------",
			"------------",
			"------------"}, 10)).To(Equal(true))

		// There are two cats in the yard, and they are closer than the minimum distance
		Expect(PeacefulYard([]string{
			"------------",
			"---M--------",
			"------------",
			"------------",
			"-------R----",
			"------------"}, 6)).To(Equal(false))

		// All three cats are in the yard, all further apart than or equal to the minimum distance
		Expect(PeacefulYard([]string{
			"-----------L",
			"--R---------",
			"------------",
			"------------",
			"------------",
			"--M---------"}, 4)).To(Equal(true))
	})
})
