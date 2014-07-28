package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// exit code = 0 if ok
// exit code > 0 if failed
// ... print messages to stdout

type ResponsePayload struct {
	Message string
	Code    int
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("wrong number of arguments")
	}

	refname := string(os.Args[1])
	oldrev := string(os.Args[2])
	newrev := string(os.Args[3])

	var u url.URL
	u.Scheme = "http"
	u.Host = "api:80"

	q := url.Values{}
	q.Set("refname", refname)
	q.Set("oldrev", oldrev)
	q.Set("newrev", newrev)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	jsonResp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var rp ResponsePayload
	err = json.Unmarshal(jsonResp, &rp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(rp.Message)
	os.Exit(rp.Code)
}
