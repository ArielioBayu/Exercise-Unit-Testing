package calculator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"exercise-unit-testing/calculator"
)

var _ = Describe("Calculator", func() {
	Describe("Add", func() {
		// blok pengujian
		It("Should add the base number with adder number", func() {
			// arrange
			calc := calculator.Calculator{Base: 3}

			// act
			calc.Add(2)

			// assertion
			Expect(calc.Result()).To(Equal(float64(5)))
		})
	})
})
