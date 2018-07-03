package cast_test

import (
	"bytes"
	"github.com/cirocosta/asciinema-edit/cast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	validHeader = cast.Header{
		Version: 2,
		Width:   123,
		Height:  123,
	}
	validEvent1 = cast.Event{
		Time: 1,
		Type: "o",
		Data: "1",
	}
	validEvent2 = cast.Event{
		Time: 2,
		Type: "o",
		Data: "2",
	}
	validEvent3 = cast.Event{
		Time: 3,
		Type: "o",
		Data: "3",
	}
	invalidEvent4 = cast.Event{
		Time: 4,
		Type: "something-wrong",
		Data: "wrong",
	}
)

var _ = Describe("Cast", func() {
	Describe("Validate", func() {
		Context("with a nil cast", func() {
			It("fails", func() {
				isValid, err := cast.Validate(nil)

				Expect(err).NotTo(Succeed())
				Expect(isValid).NotTo(BeTrue())
			})
		})

	})

	Describe("ValidateEvent", func() {
		Context("with nil event", func() {
			It("fails", func() {
				isValid, err := cast.ValidateEvent(nil)

				Expect(err).NotTo(Succeed())
				Expect(isValid).NotTo(BeTrue())
			})
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
		Context("with nil header", func() {
			It("fails", func() {
				isValid, err := cast.ValidateHeader(nil)

				Expect(err).NotTo(Succeed())
				Expect(isValid).NotTo(BeTrue())
			})
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

	Describe("ValidateEventStream", func() {
		Context("with empty stream", func() {
			var data = cast.Cast{
				Header:      validHeader,
				EventStream: []*cast.Event{},
			}

			It("is valid", func() {
				isValid, err := cast.Validate(&data)

				Expect(err).To(Succeed())
				Expect(isValid).To(BeTrue())
			})
		})

		It("fails if not sorted by time", func() {
			var data = cast.Cast{
				Header: validHeader,
				EventStream: []*cast.Event{
					&validEvent2,
					&validEvent1,
					&validEvent3,
				},
			}

			isValid, err := cast.Validate(&data)

			Expect(err).NotTo(Succeed())
			Expect(isValid).NotTo(BeTrue())

		})

		It("succeeds if sorted and valid", func() {
			var data = cast.Cast{
				Header: validHeader,
				EventStream: []*cast.Event{
					&validEvent1,
					&validEvent2,
					&validEvent3,
				},
			}

			isValid, err := cast.Validate(&data)

			Expect(err).To(Succeed())
			Expect(isValid).To(BeTrue())
		})

		It("fails if there's an invalid event", func() {
			var data = cast.Cast{
				Header: validHeader,
				EventStream: []*cast.Event{
					&validEvent1,
					&invalidEvent4,
				},
			}

			isValid, err := cast.Validate(&data)

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
