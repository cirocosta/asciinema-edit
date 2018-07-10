package transformer_test

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DummyTransformation struct{}

func (t *DummyTransformation) Transform(c *cast.Cast) (err error) {
	return
}

var _ = Describe("Transformer", func() {
	Describe("New", func() {
		Context("with nil transform", func() {
			It("fails", func() {
				_, err := transformer.New(nil, "", "")
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with transformation", func() {
			var (
				transformation transformer.Transformation
				tempDir        string
				err            error
			)

			BeforeEach(func() {
				transformation = &DummyTransformation{}
				tempDir, err = ioutil.TempDir("", "")
				Expect(err).To(Succeed())
			})

			AfterEach(func() {
				os.RemoveAll(tempDir)
			})

			Context("having empty input and output", func() {
				It("succeeds", func() {
					_, err = transformer.New(transformation, "", "")
					Expect(err).To(Succeed())
				})
			})

			Context("with input specified", func() {
				var input string

				It("fails if it doesnt exist", func() {
					input = path.Join(tempDir, "inexistent")

					_, err = transformer.New(transformation, input, "")
					Expect(err).ToNot(Succeed())
				})

				It("fails if is a directory", func() {
					input = tempDir
					_, err = transformer.New(transformation, input, "")
					Expect(err).ToNot(Succeed())
				})

				It("succeeds if file exists", func() {
					inputFile, err := ioutil.TempFile(tempDir, "")
					Expect(err).To(Succeed())

					input = inputFile.Name()

					_, err = transformer.New(transformation,
						inputFile.Name(), "")
					Expect(err).To(Succeed())
				})
			})

			Context("with output specified", func() {
				var output string

				It("fails if directory doesnt exist", func() {
					output = "/inexistent/directory/file.txt"

					_, err = transformer.New(transformation, "", output)
					Expect(err).ToNot(Succeed())
				})

				It("creates file if it doesnt exist in existing directory", func() {
					output = path.Join(tempDir, "output-file")

					_, err = transformer.New(transformation, "", output)
					Expect(err).To(Succeed())

					_, err = os.Stat(output)
					Expect(err).To(Succeed())
				})

				It("succeeds if file exists", func() {
					outputFile, err := ioutil.TempFile(tempDir, "")
					Expect(err).To(Succeed())

					output = outputFile.Name()

					_, err = transformer.New(transformation, "", output)
					Expect(err).To(Succeed())
				})
			})
		})
	})

	Describe("transform", func() {
		var (
			trans  *transformer.Transformer
			input  string
			output = "/dev/null"
			err    error
		)

		JustBeforeEach(func() {
			trans, err = transformer.New(
				&DummyTransformation{},
				input,
				output)
			Expect(err).To(Succeed())
		})

		AfterEach(func() {
			os.Remove(input)
		})

		Context("with malformed input", func() {
			BeforeEach(func() {
				input, err = createTempFileWithContent("malformed")
				Expect(err).To(Succeed())
			})

			It("fails", func() {
				err = trans.Transform()
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with malformed event stream", func() {
			BeforeEach(func() {
				input, err = createTempFileWithContent(`{"version": 2, "width": 123, "height": 123}
[1, "o", "aaa"]
[3, "o", "ccc"]
[2, "o", "bbb"]`)
				Expect(err).To(Succeed())
			})

			It("fails", func() {
				err = trans.Transform()
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with well formed event stream", func() {
			BeforeEach(func() {
				input, err = createTempFileWithContent(`{"version": 2, "width": 123, "height": 123}
[1, "o", "aaa"]
[2, "o", "bbb"]
[3, "o", "ccc"]`)
				Expect(err).To(Succeed())
			})

			It("succeeds", func() {
				err = trans.Transform()
				Expect(err).To(Succeed())
			})
		})
	})
})

func createTempFileWithContent(content string) (res string, err error) {
	var file *os.File

	file, err = ioutil.TempFile("", "")
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return
	}

	res = file.Name()
	return
}
