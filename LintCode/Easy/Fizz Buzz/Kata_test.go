package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("Fizz Buzz", func() {
		Expect(fizzBuzz(16)).Should(Equal(
			[]string{
				"1", "2", "fizz",
				"4", "buzz", "fizz",
				"7", "8", "fizz",
				"buzz", "11", "fizz",
				"13", "14", "fizz buzz", "16"}))
	})
})
