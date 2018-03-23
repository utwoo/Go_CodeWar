package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fizz Buzz Cuckoo Clock", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(FizzBuzzCuckooClock("13:34")).To(Equal("tick"))
		Expect(FizzBuzzCuckooClock("21:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
		Expect(FizzBuzzCuckooClock("11:15")).To(Equal("Fizz Buzz"))
		Expect(FizzBuzzCuckooClock("03:03")).To(Equal("Fizz"))
		Expect(FizzBuzzCuckooClock("14:30")).To(Equal("Cuckoo"))
		Expect(FizzBuzzCuckooClock("08:55")).To(Equal("Buzz"))
		Expect(FizzBuzzCuckooClock("00:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
		Expect(FizzBuzzCuckooClock("12:00")).To(Equal("Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"))
	})
})
