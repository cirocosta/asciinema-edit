package cast_test

import (
	"bytes"
	"github.com/cirocosta/asciinema-edit/cast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cast", func() {
	Describe("ValidateEvent", func() {
		It("fails if event is nil", func() {
			isValid, err := cast.ValidateEvent(nil)

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())
		})

		Context("regarding type", func() {
			It("fails if not specified", func() {
				isValid, err := cast.ValidateEvent(&cast.Event{
					Type: "",
				})

				Expect(err).NotTo(Succeed())
				Expect(isValid).NotTo(BeTrue())
			})

			It("fails if not `i` or `o`", func() {
				isValid, err := cast.ValidateEvent(&cast.Event{
					Type: "abc",
				})

				Expect(err).NotTo(Succeed())
				Expect(isValid).NotTo(BeTrue())
			})
		})

		It("succeeds if well specified", func() {
			isValid, err := cast.ValidateEvent(&cast.Event{
				Time: 123,
				Type: "o",
				Data: "lol",
			})

			Expect(err).To(Succeed())
			Expect(isValid).To(BeTrue())

			isValid, err = cast.ValidateEvent(&cast.Event{
				Time: 321,
				Type: "i",
				Data: "lol",
			})

			Expect(err).To(Succeed())
			Expect(isValid).To(BeTrue())
		})
	})

	Describe("ValidateHeader", func() {
		It("fails if header is nil", func() {
			isValid, err := cast.ValidateHeader(nil)

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())
		})

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

		It("succeeds if well specified", func() {
			isValid, err := cast.ValidateHeader(&cast.Header{
				Version: 2,
				Width:   123,
				Height:  321,
			})

			Expect(err).To(Succeed())
			Expect(isValid).To(BeTrue())
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
