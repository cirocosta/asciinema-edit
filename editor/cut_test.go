package editor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/editor"
)

var (
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
)

var _ = Describe("Cut", func() {
	Context("with nil cast", func() {
		It("fails", func() {
			err := editor.Cut(nil, 1, 2)
			Expect(err).ToNot(Succeed())
		})
	})

	Context("with an empty event stream", func() {
		var data = &cast.Cast{
			EventStream: []*cast.Event{},
		}

		It("errors", func() {
			err := editor.Cut(data, 1, 2)
			Expect(err).ToNot(Succeed())
		})
	})

	Context("with non-empty event stream", func() {
		var (
			err                   error
			data                  *cast.Cast
			initialNumberOfEvents int
		)

		BeforeEach(func() {
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

		It("fails if `from` > `to`", func() {
			err = editor.Cut(data, 3, 2)
			Expect(err).ToNot(Succeed())
		})

		It("cuts single frame if `from` == `to`", func() {
			err = editor.Cut(data, 1.2, 1.2)
			Expect(err).To(Succeed())

			Expect(data.EventStream).To(ContainElement(event1))
			Expect(data.EventStream).ToNot(ContainElement(event1_2))
			Expect(data.EventStream).To(ContainElement(event1_6))
			Expect(data.EventStream).To(ContainElement(event2))

			Expect(len(data.EventStream)).To(Equal(initialNumberOfEvents - 1))
		})

		It("cuts frames in range", func() {
			err = editor.Cut(data, 1.2, 2)
			Expect(err).To(Succeed())

			Expect(data.EventStream).To(ContainElement(event1))
			Expect(data.EventStream).ToNot(ContainElement(event1_2))
			Expect(data.EventStream).ToNot(ContainElement(event1_6))
			Expect(data.EventStream).ToNot(ContainElement(event2))

			Expect(len(data.EventStream)).To(Equal(initialNumberOfEvents - 3))
		})
	})
})
