package mapify_test

import (
	. "github.com/clouch/mapify"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mapify", func() {
	Describe("Convert a struct to a map", func() {
		Context("Without clouch and json tags", func() {
			test := TestA{
				ID:    "1",
				Revs:  []string{"1234"},
				Hello: "World",
			}
			It("Should return a map of of the necessaty values", func() {
				testA, err := Do(&test)
				Ω(err).Should(BeNil())

				var world interface{} = "World"
				var revs interface{} = []string{"1234"}
				var id interface{} = "1"

				Ω(testA).Should(HaveKeyWithValue("ID", id))
				Ω(testA).Should(HaveKeyWithValue("Hello", world))
				Ω(testA).Should(HaveKeyWithValue("Revs", revs))
			})
		})
	})
})

type TestA struct {
	ID    string
	Revs  []string
	Hello string
}
