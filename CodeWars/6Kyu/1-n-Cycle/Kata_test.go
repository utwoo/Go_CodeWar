package kata

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int, exp int) {
	var ans = cycle(n)
	if ans == exp {
		fmt.Printf("expected %v Got %v \n", exp, ans)
	} else {
		fmt.Printf("expected %v but got %v \n", exp, ans)
	}
	Expect(ans).To(Equal(exp))
}

var _ = Describe("1/n- Cycle", func() {
	It("should handle basic cases", func() {
		dotest(33, 2)
		dotest(18118, -1)
		dotest(69, 22)
		dotest(197, 98)
		dotest(22, -1)
	})
})
