package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("Merge Two Sorted Arrays", func() {
		Expect(mergeSortedArray([]int{1, 2, 3, 4}, []int{2, 4, 5, 6})).Should(Equal([]int{1, 2, 2, 3, 4, 4, 5, 6}))
	})
})
