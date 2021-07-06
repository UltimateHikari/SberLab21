package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetVpcsList(ids ...string) (resp string, err error) {
	curId := defaultProjectId
	switch {
	case len(ids) > 1:
		return "", errors.New("too many projIds")
	case len(ids) == 1:
		curId = ids[1]
	}
	req, _ := http.NewRequest(
		"GET",
		urlVpcsList(curId),
		ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
	)
	req.Header.Add("X-Project-Id", curId)
	return doGet(req)
}

func urlVpcsList(curId string) (s string) {
	return "https://" + vpcPrefix + endpoint + "v1/" + curId + vpcSuffix
}
