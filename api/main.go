package main

import (
	"api/apigw/core"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var signer core.Signer

const (
	defaultProjectId = "0b5a73ddd98027372f2ec00668b88856"
	vpcPrefix        = "vpc"
	iamPrefix        = "iam"
	endpoint         = ".ru-moscow-1.hc.sbercloud.ru/"
	vpcSuffix        = "/vpcs?limit=2"
	defaulturl       = "https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/0b5a73ddd98027372f2ec00668b88856/vpcs?limit=2"
	iamurl           = "https://iam.ru-moscow-1.hc.sbercloud.ru/v3/users/0ce43a5b788024e71f03c0060aaf6125/projects"
)

func main() {
	signer = initSigner()
	resp, err := getVpcsList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func doGet(req *http.Request) (s string, err error) {
	signer.Sign(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
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
