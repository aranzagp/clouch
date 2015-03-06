package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	Ok bool `json:"ok"`
}

func CreateDatabase(dbname, host string) error {

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", host, dbname), nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Cannot create database")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if !res.Ok {
		return errors.New("Cannot create database")
	}

	return nil
}

func DeleteDatabase(dbname, host string) error {
	fmt.Print(fmt.Sprintf("URL %s", host))
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", host, dbname), nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	res := response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if !res.Ok {
		return errors.New("Cannot delete database")
	}

	return nil
}
