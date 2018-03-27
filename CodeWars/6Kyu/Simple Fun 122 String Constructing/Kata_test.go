package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Fun #122: String Constructing", func() {
	It("should pass static tests", func() {
		Expect(StringConstructing("aba", "abbabba")).To(Equal(9))
		Expect(StringConstructing("aba", "abaa")).To(Equal(4))
		Expect(StringConstructing("aba", "a")).To(Equal(3))
		Expect(StringConstructing("a", "a")).To(Equal(1))
		Expect(StringConstructing("a", "aaa")).To(Equal(3))
		Expect(StringConstructing("abcdefgh", "hgfedcba")).To(Equal(64))
		Expect(StringConstructing(
			"sxdfcgdgxdfgdxxf",
			"xgdfsxgdxgfsgdfxgfdgfgdfgdgsdfxgfdxgdfxgdgdfgdfxgsdxgdfxgfgsxfgdfgsdfxgdfxgsgsfgxsgsdxgsdfxgsgsdfxgsdfx",
		)).To(Equal(288))
		Expect(StringConstructing("bbaabcbcbc", "bbcccbabcc")).To(Equal(12))
	})
})
