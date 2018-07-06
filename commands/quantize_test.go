package commands_test

import (
	"math"

	"github.com/cirocosta/asciinema-edit/commands"
	"github.com/cirocosta/asciinema-edit/editor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseQuantizeRange", func() {
	Context("having empty input", func() {
		var input = ""

		It("fails", func() {
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})
	})

	Context("having invalid chars", func() {
		var input string

		It("fails if non-decimal", func() {
			input = "1a"
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})

		It("fails if starts w/ non-numeric", func() {
			input = "a"
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})

		It("fails if ends w/ non-numeric", func() {
			input = "1,a"
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})

		It("fails with leading comma", func() {
			input = ",1"
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})

		It("fails with trailing comma", func() {
			input = "1,"
			_, err := commands.ParseQuantizeRange(input)
			Expect(err).NotTo(Succeed())
		})
	})

	Context("having a valid entry", func() {
		var (
			input  string
			qRange editor.QuantizeRange
			err    error
		)

		Context("that is unbounded", func() {
			BeforeEach(func() {
				input = "1.2"
			})

			JustBeforeEach(func() {
				qRange, err = commands.ParseQuantizeRange(input)
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})

			It("starts with lower bound", func() {
				Expect(qRange.From).To(Equal(float64(1.2)))
			})

			It("ends with the closes to infinit (max float)", func() {
				Expect(qRange.To).To(Equal(math.MaxFloat64))
			})

			Context("with negative lower bound", func() {
				BeforeEach(func() {
					input = "-1.2"
				})

				It("fails", func() {
					Expect(err).ToNot(Succeed())
				})
			})
		})

		Context("that is bounded with `from > to`", func() {
			BeforeEach(func() {
				input = "2,1"
				qRange, err = commands.ParseQuantizeRange(input)
			})

			It("fails", func() {
				Expect(err).ToNot(Succeed())
			})
		})

		Context("that is bounded with `from == to`", func() {
			BeforeEach(func() {
				input = "2,2"
				qRange, err = commands.ParseQuantizeRange(input)
			})

			It("fails", func() {
				Expect(err).ToNot(Succeed())
			})
		})
	})
})
