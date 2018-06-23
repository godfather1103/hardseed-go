package utils

import (
	"net/http"
	"io/ioutil"
	"io"
	"net/url"
)

type HttpClientParam struct {
	Proxy url.URL
}


func GetHttpClientParam(proxy string)  HttpClientParam{
	httpClientParam := HttpClientParam{}
	if len(proxy) > 0 {
		urli := url.URL{}
		urlproxy, _ := urli.Parse(proxy)
		httpClientParam.Proxy = *urlproxy
	}
	return httpClientParam
}

var client *http.Client

func initClient(clientParam HttpClientParam){
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(&clientParam.Proxy),
			},
		}
	}
}

func GetHttpInfo(httpUrl string, method string,param io.Reader,clientParam HttpClientParam) (string, int) {

	if client == nil {
		initClient(clientParam)
	}

	request, err := http.NewRequest(method, httpUrl, param)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.18 Safari/537.36")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	result := string(body)
	code := response.StatusCode
	return result, code
}

func GetHttpInfoHeader(httpUrl string, method string,param io.Reader,header map[string]string,clientParam HttpClientParam) (string, int) {

	if client == nil {
		initClient(clientParam)
	}

	request, err := http.NewRequest(method, httpUrl, param)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.18 Safari/537.36")

	for key,value := range header {
		request.Header.Set(key,value)
	}

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	result := string(body)
	code := response.StatusCode
	return result, code
}

func GetMultipart(httpUrl string, params map[string]string,contentType string,clientParam HttpClientParam)  (string, int){
	if client == nil {
		initClient(clientParam)
	}

	vals := url.Values{}
	for key,value:=range params{
		if len(key)>0 && len(value)>0 {
			vals.Set(key,value)
		}
	}
	response, _ := http.PostForm(httpUrl,vals)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	result := string(body)
	code := response.StatusCode
	return result, code
}