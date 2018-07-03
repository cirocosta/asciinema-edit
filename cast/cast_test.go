package cast_test

import (
	"bytes"
	"github.com/cirocosta/asciinema-edit/cast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cast", func() {
	Describe("ValidateHeader", func() {
		It("fails if version is not 2", func() {
			isValid, err := cast.ValidateHeader(&cast.Header{
				Version: 123,
			})

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())
		})

		It("fails if width is not greater than zero ", func() {
			isValid, err := cast.ValidateHeader(&cast.Header{
				Version: 2,
				Width:   0,
			})

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())
		})

		It("fails if height is not greater than zero ", func() {
			isValid, err := cast.ValidateHeader(&cast.Header{
				Version: 2,
				Width:   10,
				Height:  0,
			})

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())
		})
	})

	Describe("Encode", func() {
		Context("with nil writer", func() {
			It("fails", func() {
				err := cast.Encode(nil, nil)
				Expect(err).NotTo(Succeed())
			})
		})

		Context("with nil cast", func() {
			var buf bytes.Buffer

			It("fails", func() {
				err := cast.Encode(&buf, nil)
				Expect(err).NotTo(Succeed())
			})
		})
	})
})
