package utils

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Get(uri string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)

	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return nil, err
	}
	body := resp.Body()

	return body, nil

}

func Post(uri string, payload string, content_type string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType(content_type)

	req.SetRequestURI(uri)

	req.SetBody([]byte(payload))

	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return nil, err
	}

	body := resp.Body()

	// fmt.Printf("Response body is: %s", body)

	return body, nil

}
