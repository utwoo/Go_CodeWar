package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Piano Kata, Part 1", func() {
	It("should test that the solution returns the correct value", func() {
		// Expect(BlackOrWhiteKey(1)).To(Equal("white"))
		// Expect(BlackOrWhiteKey(5)).To(Equal("black"))
		// Expect(BlackOrWhiteKey(12)).To(Equal("black"))
		// Expect(BlackOrWhiteKey(42)).To(Equal("white"))
		Expect(BlackOrWhiteKey(87)).To(Equal("white"))
		Expect(BlackOrWhiteKey(88)).To(Equal("white"))
		Expect(BlackOrWhiteKey(89)).To(Equal("white"))
		// Expect(BlackOrWhiteKey(90)).To(Equal("black"))
		// Expect(BlackOrWhiteKey(91)).To(Equal("white"))
		// Expect(BlackOrWhiteKey(92)).To(Equal("white"))
		// Expect(BlackOrWhiteKey(100)).To(Equal("black"))
		// Expect(BlackOrWhiteKey(111)).To(Equal("white"))
		// Expect(BlackOrWhiteKey(200)).To(Equal("black"))
		// Expect(BlackOrWhiteKey(2017)).To(Equal("white"))
	})
})
