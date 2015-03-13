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

			Context("Test with a struct with a struct inside", func() {
				testA := TestA{
					ID:    "1",
					Revs:  []string{"1234"},
					Hello: "World",
				}

				testC := TestC{
					ID:    "1234",
					Revs:  []string{"1234"},
					TestA: testA,
				}
				It("Should return a nested map with correct values", func() {
					test, err := Do(&testC)
					Ω(err).Should(BeNil())

					var revs interface{} = []string{"1234"}
					var id interface{} = "1"
					var id2 interface{} = "1234"

					Ω(test).Should(HaveKeyWithValue("ID", id2))
					Ω(test).Should(HaveKeyWithValue("Revs", revs))

					Ω(test).Should(HaveKey("TestA"))
					Ω(test["TestA"]).Should(HaveKeyWithValue("Revs", revs))
					Ω(test["TestA"]).Should(HaveKeyWithValue("ID", id))

				})
			})

			Context("Test with a struct with a pointer struct inside", func() {
				testA := TestA{
					ID:    "1",
					Revs:  []string{"1234"},
					Hello: "World",
				}

				testD := TestD{
					ID:    "1234",
					Revs:  []string{"1234"},
					TestA: &testA,
				}
				It("Should return a nested map with correct values", func() {
					test, err := Do(&testD)
					Ω(err).Should(BeNil())

					var revs interface{} = []string{"1234"}
					var id interface{} = "1"
					var id2 interface{} = "1234"

					Ω(test).Should(HaveKeyWithValue("ID", id2))
					Ω(test).Should(HaveKeyWithValue("Revs", revs))

					Ω(test).Should(HaveKey("TestA"))
					Ω(test["TestA"]).Should(HaveKeyWithValue("Revs", revs))
					Ω(test["TestA"]).Should(HaveKeyWithValue("ID", id))

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

type TestC struct {
	ID    string
	Revs  []string
	TestA TestA
}

type TestD struct {
	ID    string
	Revs  []string
	TestA *TestA
}
