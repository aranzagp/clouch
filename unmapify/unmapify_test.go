package unmapify_test

import (
	ts "github.com/thetonymaster/clouch/teststructs"
	. "github.com/thetonymaster/clouch/unmapify"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmapify", func() {

	Describe("Unmap a map", func() {

		Context("Unmap a simple map", func() {
			testA := ts.TestA{}
			mp := map[string]interface{}{
				"Hello": "World",
			}
			It("Should fill the struct with correct values", func() {
				err := Do(&testA, mp)
				Ω(err).Should(BeNil())

				Ω(testA.Hello).Should(Equal("World"))
			})
		})

		Context("Unmap with tags", func() {
			testF := ts.TestF{}
			mp := map[string]interface{}{
				"hello": "World",
			}
			It("Should fill the fields with the correct values", func() {
				err := Do(&testF, mp)
				Ω(err).Should(BeNil())

				Ω(testF.Hello).Should(Equal("World"))
			})
		})

	})

})
