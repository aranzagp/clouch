package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DB struct {
	host   string
	dbname string
	url    url.URL
}

func New(host, dbname string) (*DB, error) {
	url, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	url.Path = dbname

	db := DB{
		host:   host,
		dbname: dbname,
		url:    *url,
	}

	return &db, nil
}

type response struct {
	Ok bool `json:"ok"`
}

func (db DB) CreateDatabase() error {

	req, err := http.NewRequest("PUT", db.url.String(), nil)

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

func (db DB) DeleteDatabase() error {

	req, err := http.NewRequest("DELETE", db.url.String(), nil)

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

func (db DB) URL() string {
	return db.url.String()
}
