package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func PreauthCookie(w http.ResponseWriter, r *http.Request) {

	url := "https://ai.fakeopen.com/auth/preauth"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, string(body))

}
