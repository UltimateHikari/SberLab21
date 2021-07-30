package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetProjList(id string) (resp string, err error) {
	req, _ := http.NewRequest(
		"GET",
		urlVpcsList(id),
		ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
	)
	req.Header.Add("X-Project-Id", id)
	return doGet(req)
}

func urlProjList(curId string) (s string) {
	return "https://" + iamPrefix + endpoint + "v3/users/" + curId + iamSuffix
}
