package db

type response struct {
	Ok bool `json:"ok"`
}

type documentResponse struct {
	Ok  bool   `json:"ok"`
	ID  string `json:"id"`
	Rev string `json:"rev"`
}
