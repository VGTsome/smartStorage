package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Param map[string]interface{}

type httpRequest struct {
	Url        string
	Params     Param
	Method     string
	Status     string
	StatusCode int
	Header     map[string][]string
	Body       string
	Response   *http.Response
	Proto      string
	Host       string
	URL        *url.URL
}

var req *httpRequest

func (r *httpRequest) get(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return r.do(url, params, "GET", isJSON)
}

func (r *httpRequest) post(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return r.do(url, params, "POST", isJSON)
}

func (r *httpRequest) put(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return r.do(url, params, "PUT", isJSON)
}

func (r *httpRequest) delete(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return r.do(url, params, "DELETE", isJSON)
}

func (r *httpRequest) do(url string, params Param, method string, isJSON bool) (rest *httpRequest, err error) {
	var reqs *http.Request
	if isJSON {
		jpar, _ := json.Marshal(params)
		reqs, err = http.NewRequest(method, url, bytes.NewReader(jpar))
	} else {
		reqs, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}
	// 添加header
	reqs.Header.Add("Authorization", "token_value")
	reqs.Header.Add("Content-Type", "text/plain; charset=UTF-8")
	reqs.Header.Add("User-Agent", "Go-http-client/1.14")
	reqs.Header.Add("Transfer-Encoding", "chunked")
	reqs.Header.Add("Accept-Encoding", "gzip, deflate")
	reqs.Header.Add("x-token", "xuwang11223344")
	// 设置参数
	q := reqs.URL.Query()
	for k, v := range params {
		q.Add(k, fmt.Sprint(v))
	}
	if !isJSON {
		reqs.URL.RawQuery = q.Encode()
	}
	res, err := http.DefaultClient.Do(reqs)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var body string
	if res.StatusCode == 200 {
		switch res.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(res.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(res.Body)
			body = string(bodyByte)
		}
	} else {
		bodyByte, _ := ioutil.ReadAll(res.Body)
		body = string(bodyByte)
	}
	rest = &httpRequest{
		Url:        url,
		Params:     params,
		Method:     method,
		Body:       body,
		Header:     reqs.Header,
		Response:   res,
		Proto:      reqs.Proto,
		Host:       reqs.Host,
		URL:        reqs.URL,
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}
	return rest, nil
}

func (r *httpRequest) Dump() {
	fmt.Println("----------------------------------------------------")
	fmt.Println(r.Method, r.Proto)
	fmt.Println("Host", ":", r.Host)
	fmt.Println("URL", ":", r.URL)
	fmt.Println("RawQuery", ":", r.URL.RawQuery)
	for key, val := range r.Header {
		fmt.Println(key, ":", val)
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println("Status", ":", r.Status)
	for key, val := range r.Response.Header {
		fmt.Println(key, ":", val)
	}
}

func Get(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return req.get(url, params, isJSON)
}

func Post(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return req.post(url, params, isJSON)
}

func Put(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return req.put(url, params, isJSON)
}

func Delete(url string, params Param, isJSON bool) (rest *httpRequest, err error) {
	return req.delete(url, params, isJSON)
}
func GetData(urlstr string) (body string, err error) {
	client := &http.Client{}

	resp, err := client.Get(urlstr)

	defer resp.Body.Close()
	bodybyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = "error"
	}
	body = string(bodybyte)
	return body, err
}
