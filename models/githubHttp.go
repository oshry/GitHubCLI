/*
Copyright © 2022 Oshry Levy
Everything we need for new http request
*/
package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Secure struct {
	Login    string
	Password string
	Token    string
}

type HttpObj struct {
	Url string
	Secure
}

func (r *RunCmd) GithubApiPrepareReq() {
	newReq, err := http.NewRequest(
		http.MethodGet,
		r.Url,
		nil,
	)
	if err != nil {
		log.Printf("Could not request repository from gitub. %v", err)
	}

	if len(r.Secure.Login)+len(r.Secure.Password) > 0 {
		newReq.SetBasicAuth(r.Secure.Login, r.Secure.Password)
	} else if len(r.Secure.Token) > 0 {
		newReq.Header.Add("Authorization", fmt.Sprintf("token %s", r.Secure.Token))
	}

	// headers
	newReq.Header.Add("Accept", "application/vnd.github.v3+json")
	r.HttpReq = newReq
}

func (r *RunCmd) HttpExecuteReq() []byte {
	response, err := http.DefaultClient.Do(r.HttpReq)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
