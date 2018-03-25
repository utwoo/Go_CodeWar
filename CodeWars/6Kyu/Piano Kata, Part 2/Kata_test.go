package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Piano Kata, Part 2", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(WhichNote(1)).To(Equal("A"))
		Expect(WhichNote(5)).To(Equal("C#"))
		Expect(WhichNote(12)).To(Equal("G#"))
		Expect(WhichNote(42)).To(Equal("D"))
		Expect(WhichNote(88)).To(Equal("C"))
		Expect(WhichNote(89)).To(Equal("A"))
		Expect(WhichNote(92)).To(Equal("C"))
		Expect(WhichNote(100)).To(Equal("G#"))
		Expect(WhichNote(111)).To(Equal("G"))
		Expect(WhichNote(200)).To(Equal("G#"))
		Expect(WhichNote(2017)).To(Equal("F"))
	})
})
