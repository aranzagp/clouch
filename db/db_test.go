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
	Describe("Create a new Database", func() {
		It("Should create and return no errors", func() {
			fake, _ := faker.New("en")
			name := fake.Words(1, false)[0]
			host := "http://127.0.0.1:5984/"
			err := CreateDatabase(name, host)
			Ω(err).ShouldNot(HaveOccurred())

			dbs, err := getDatabases(host)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(dbs).Should(ContainElement(name))

			err = DeleteDatabase(name, host)

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
