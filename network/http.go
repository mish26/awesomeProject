package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
/*	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

 */

	base, _ := url.Parse("http://example.com/ssgaga")
	reference, _ := url.Parse("/test?A=1&B=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)


	request, _ := http.NewRequest(http.MethodGet, endpoint,nil)
	// postの場合は、bodyに値が必要
	// request, _ := http.NewRequest(http.MethodPost, endpoint,bytes.NewBuffer([]byte("pw")))
	request.Header.Add("if-None-Match", `E/"dhshs"`)
	query := request.URL.Query()
	fmt.Println(query)
	query.Add("C","tagaga&%")
	fmt.Println(query)
	fmt.Println(query.Encode())
	request.URL.RawQuery = query.Encode()
	fmt.Println(request)

	var client *http.Client = &http.Client{}
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

