package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type CallAPI struct {
	URL string
}

func (c CallAPI) CallAPIPostJson(values interface{}) interface{} {
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Println(err)
	}

	resp, err := http.Post(c.URL, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func (c CallAPI) CallAPIGetJson() interface{} {
	resp, err := http.Get(c.URL)

	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func (c CallAPI) CallAPIPatchJson(values interface{}) interface{} {
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, c.URL, bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func (c CallAPI) CallAPIDeleteJson(values interface{}) interface{} {
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, c.URL, bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
