package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubWebhookRequestBody struct {
	Action     string
	Number     int64
	Repository string
	Sender     string
}

// implements http.Handler
type GithubWebhookHandler struct{}

func (h *GithubWebhookHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		obj := new(GithubWebhookRequestBody)
		err = json.Unmarshal(body, obj)
		if err != nil {
			panic(err)
		}
		fmt.Println("Action", obj.Action)
		fmt.Println("Number", obj.Number)
		fmt.Println("Sender", obj.Sender)
		fmt.Println("Repository", obj.Repository)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Ok"))
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Not OK"))
}

func main() {
	handler := new(GithubWebhookHandler)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

}
