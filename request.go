package fluidcurrency

import (
	"io/ioutil"
	"net/http"
	"time"
)

// type request interface {
// 	Get(URL string) ([]byte, error)
// }
//
// func newRequest() request {
// 	return newHTTPRequest()
// }
//
// type httpRequest struct {
// 	netClient *http.Client
// }
//
// func newHTTPRequest() *httpRequest {
// 	client := &http.Client{
// 		Timeout: time.Second * 10,
// 	}
// 	return &httpRequest{
// 		netClient: client,
// 	}
// }

// Get main get function with a url
func Get(URL string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// return &httpRequest{
	// 	netClient: client,
	// }
	response, err := client.Get(URL)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
