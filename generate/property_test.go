package main_test

import (
	"encoding/json"

	"github.com/grahamjenson/goformation/cloudformation/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goformation Code Generator", func() {

	Context("with a single generated custom property", func() {

		Context("specified as a Go struct", func() {

			property := &resources.AWSServerlessFunction_S3Location{
				Bucket:  "test-bucket",
				Key:     "test-key",
				Version: 123,
			}
			expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

			result, err := json.Marshal(property)
			It("should marshal to JSON successfully", func() {
				Expect(result).To(Equal(expected))
				Expect(err).To(BeNil())
			})

		})
		Context("specified as JSON", func() {

			property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)
			expected := &resources.AWSServerlessFunction_S3Location{
				Bucket:  "test-bucket",
				Key:     "test-key",
				Version: 123,
			}

			result := &resources.AWSServerlessFunction_S3Location{}
			err := json.Unmarshal(property, result)
			It("should unmarshal to a Go struct successfully", func() {
				Expect(result).To(Equal(expected))
				Expect(err).To(BeNil())
			})

		})

	})

	Context("with a polymorphic property", func() {

		Context("with a primitive value", func() {

			Context("specified as a Go struct", func() {

				value := "test-primitive-value"
				property := &resources.AWSServerlessFunction_CodeUri{
					String: &value,
				}

				expected := []byte(`"test-primitive-value"`)
				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`"test-primitive-value"`)
				value := "test-primitive-value"
				expected := &resources.AWSServerlessFunction_CodeUri{
					String: &value,
				}

				result := &resources.AWSServerlessFunction_CodeUri{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

		Context("with a custom type value", func() {

			Context("specified as a Go struct", func() {

				property := &resources.AWSServerlessFunction_CodeUri{
					S3Location: &resources.AWSServerlessFunction_S3Location{
						Bucket:  "test-bucket",
						Key:     "test-key",
						Version: 123,
					},
				}

				expected := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				result, err := json.Marshal(property)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("specified as JSON", func() {

				property := []byte(`{"Bucket":"test-bucket","Key":"test-key","Version":123}`)

				expected := &resources.AWSServerlessFunction_CodeUri{
					S3Location: &resources.AWSServerlessFunction_S3Location{
						Bucket:  "test-bucket",
						Key:     "test-key",
						Version: 123,
					},
				}

				result := &resources.AWSServerlessFunction_CodeUri{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

})
