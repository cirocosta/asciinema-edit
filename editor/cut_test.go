package editor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/editor"
)

var _ = Describe("Cut", func() {
	Describe("parameter validation", func() {
		var data = &cast.Cast{
			EventStream: []*cast.Event{},
		}

		Context("with nil cast", func() {
			It("fails", func() {
				err := editor.Cut(nil, 1, 2)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with an empty event stream", func() {

			It("errors", func() {
				err := editor.Cut(data, 1, 2)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with `from` > `to`", func() {
			It("fails", func() {
				err := editor.Cut(data, 3, 2)
				Expect(err).ToNot(Succeed())
			})
		})
	})

	Context("with non-empty event stream", func() {
		var (
			err                                error
			data                               *cast.Cast
			initialNumberOfEvents              int
			event1, event1_2, event1_6, event2 *cast.Event
		)

		BeforeEach(func() {
			event1 = &cast.Event{
				Time: 1,
				Data: "event1",
			}
			event1_2 = &cast.Event{
				Time: 1.2,
				Data: "event1_2",
			}
			event1_6 = &cast.Event{
				Time: 1.6,
				Data: "event1_6",
			}
			event2 = &cast.Event{
				Time: 2,
				Data: "event2",
			}

			data = &cast.Cast{
				EventStream: []*cast.Event{
					event1,
					event1_2,
					event1_6,
					event2,
				},
			}

			initialNumberOfEvents = len(data.EventStream)
		})

		Context("with `from` not found", func() {
			It("fails", func() {
				err = editor.Cut(data, 1.1, 2)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with `to` not found", func() {
			It("fails", func() {
				err = editor.Cut(data, 2, 3.3)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("cutting a single frame when `from` == `to`", func() {
			JustBeforeEach(func() {
				err = editor.Cut(data, 1.2, 1.2)
				Expect(err).To(Succeed())
			})

			It("removes the frame", func() {
				Expect(data.EventStream).To(ContainElement(event1))
				Expect(data.EventStream).ToNot(ContainElement(event1_2))
				Expect(data.EventStream).To(ContainElement(event1_6))
				Expect(data.EventStream).To(ContainElement(event2))

				Expect(len(data.EventStream)).
					To(Equal(initialNumberOfEvents - 1))
			})

			It("adjusts the remaining time stamps", func() {
				Expect(event1.Time).To(Equal(float64(1)))
				Expect(event1_6.Time).To(Equal(float64(1.2)))
				Expect(event2.Time).To(Equal(float64(1.6)))
			})
		})

		Context("cutting range without bounds included", func() {
			JustBeforeEach(func() {
				err = editor.Cut(data, 1.2, 1.6)
				Expect(err).To(Succeed())
			})

			It("removes the frame", func() {
				Expect(data.EventStream).To(ContainElement(event1))
				Expect(data.EventStream).ToNot(ContainElement(event1_2))
				Expect(data.EventStream).ToNot(ContainElement(event1_6))
				Expect(data.EventStream).To(ContainElement(event2))

				Expect(len(data.EventStream)).
					To(Equal(initialNumberOfEvents - 2))
			})

			It("adjusts the remaining time stamps", func() {
				Expect(event1.Time).To(Equal(float64(1)))
				Expect(event2.Time).To(Equal(float64(1.2)))
			})
		})

		It("cuts frames in range containing last element", func() {
			err = editor.Cut(data, 1.2, 2)
			Expect(err).To(Succeed())

			Expect(data.EventStream).To(ContainElement(event1))
			Expect(data.EventStream).ToNot(ContainElement(event1_2))
			Expect(data.EventStream).ToNot(ContainElement(event1_6))
			Expect(data.EventStream).ToNot(ContainElement(event2))

			Expect(len(data.EventStream)).
				To(Equal(initialNumberOfEvents - 3))
		})
	})
})
