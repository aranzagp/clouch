package db_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/clouch/db"

	"github.com/manveru/faker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Db", func() {

	var (
		db     *DB
		dbname string
		host   string
	)

	BeforeEach(func() {
		fake, _ := faker.New("en")
		dbname = fake.Words(1, false)[0]
		host = "http://127.0.0.1:5984/"
		db, _ = New(host, dbname)
	})

	Describe("Test Helper methods", func() {
		It("Should return the correct URL", func() {
			Ω(db.URL()).Should(Equal(host + dbname))
		})
	})

	Describe("Create a new Database", func() {
		It("Should create and return no errors", func() {

			err := db.CreateDatabase()
			Ω(err).ShouldNot(HaveOccurred())

			dbs, err := getDatabases(host)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(dbs).Should(ContainElement(dbname))

			err = db.DeleteDatabase()

			Ω(err).ShouldNot(HaveOccurred())

		})
	})

	Describe("Create a new document", func() {
		type document struct {
			ID    string
			Revs  []string
			Hello string
		}

		It("Should return an struct with an id, and rev filled", func() {
			err := db.CreateDatabase()
			Ω(err).ShouldNot(HaveOccurred())

			dbs, err := getDatabases(host)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(dbs).Should(ContainElement(dbname))

			doc := document{Hello: "World"}
			err = db.Create(&doc)
			Ω(err).ShouldNot(HaveOccurred())

			err = db.DeleteDatabase()
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})

func getDatabases(host string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", host, "_all_dbs"), nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := []string{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil

}
