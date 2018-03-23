package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp []int) {
	var ans = Game(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Playing on a chessboard", func() {
	It("should handle basic cases", func() {
		dotest(0, []int{0})
		dotest(1, []int{1, 2})
		dotest(8, []int{32})
		dotest(40, []int{800})
	})
})
