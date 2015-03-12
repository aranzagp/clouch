package mapify_test

import (
	. "github.com/clouch/mapify"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mapify", func() {
	Describe("Convert a struct to a map", func() {
		Context("Without pointer values", func() {
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
			Context("Test with a struct with pointer values", func() {
				bar := "Bar"
				num := 1
				test := TestB{
					ID:    "1",
					Revs:  []string{"1234"},
					Hello: "World",
					Foo:   &bar,
					Num:   &num,
				}
				It("Should return a map of of the necessaty values", func() {
					testA, err := Do(&test)
					Ω(err).Should(BeNil())

					var world interface{} = "World"
					var revs interface{} = []string{"1234"}
					var id interface{} = "1"
					var foo interface{} = bar
					var num interface{} = 1

					Ω(testA).Should(HaveKeyWithValue("ID", id))
					Ω(testA).Should(HaveKeyWithValue("Hello", world))
					Ω(testA).Should(HaveKeyWithValue("Revs", revs))
					Ω(testA).Should(HaveKeyWithValue("Foo", foo))
					Ω(testA).Should(HaveKeyWithValue("Num", num))
				})
			})

			Context("Test with a struct with nil pointer values", func() {
				bar := "Bar"
				test := TestB{
					ID:    "1",
					Revs:  []string{"1234"},
					Hello: "World",
					Foo:   &bar,
				}
				It("Should return a map of of the necessaty values", func() {
					testA, err := Do(&test)
					Ω(err).Should(BeNil())

					var world interface{} = "World"
					var revs interface{} = []string{"1234"}
					var id interface{} = "1"
					var foo interface{} = bar

					Ω(testA).Should(HaveKeyWithValue("ID", id))
					Ω(testA).Should(HaveKeyWithValue("Hello", world))
					Ω(testA).Should(HaveKeyWithValue("Revs", revs))
					Ω(testA).Should(HaveKeyWithValue("Foo", foo))
					Ω(testA).Should(HaveKey("Num"))
					Ω(testA["Num"]).Should(BeNil())
				})
			})

			Context("Test with a struct with nil and not-nil pointer values", func() {
				It("Should return a map of of the necessaty values", func() {
				})
			})

		})
	})
})

type TestA struct {
	ID    string
	Revs  []string
	Hello string
}

type TestB struct {
	ID    string
	Revs  []string
	Hello string
	Foo   *string
	Num   *int
}
