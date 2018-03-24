package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(n int, exp int) {
	var ans = Thirt(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("A Rule of Divisibility by 13", func() {
	It("should handle basic cases", func() {
		testequal(1234567, 87)
		testequal(8529, 79)
		testequal(85299258, 31)
		testequal(5634, 57)
	})
})
