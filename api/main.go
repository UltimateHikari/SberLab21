package main

import (
	"api/apigw/core"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	url        = "https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/0ce61dbdd30024ef2f4dc006072bb596/vpcs?limit=2"
	defaulturl = "https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/0b5a73ddd98027372f2ec00668b88856/vpcs?limit=2"
)

func main() {
	s := initSigner()
	//fmt.Println(s.Key)
	req, _ := http.NewRequest(
		"GET",
		defaulturl,
		ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
	)

	req.Header.Add("content-type", "application/json")

	s.Sign(req)
	//fmt.Println(req.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}

func initSigner() (s core.Signer) {
	f, _ := os.Open("../../keys")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var tokens [2]string
	for i := 0; scanner.Scan(); i++ {
		tokens[i] = scanner.Text()
		//fmt.Println(tokens[i])
	}

	s = core.Signer{
		Key:    tokens[0],
		Secret: tokens[1],
	}
	return s
}
