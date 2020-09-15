package v_utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
const requestHttpTimeout = 2 * time.Minute


func (slf *evoHttpClient) Do(request *http.Request) (*http.Response, error) {
	return slf.client.Do(request)
}

func getHttpClient(isSkipSecurityChecking bool, timeOut *time.Duration) (client *httpclient.Client) {
	httpTimeOut := requestHttpTimeout
	if timeOut != nil {
		httpTimeOut = *timeOut
	}

	if !isSkipSecurityChecking {
		client = httpclient.NewClient(httpclient.WithHTTPTimeout(httpTimeOut))
	} else {
		client = httpclient.NewClient(
			httpclient.WithHTTPClient(&evoHttpClient{
				client: http.Client{
					Timeout: httpTimeOut,
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
						},
					},
				},
			}))
	}
	return
}



func (*VS) HttpGet(url string, header map[string]string, params map[string]string, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error) {
	client := getHttpClient(isSkipSecurityChecking, timeOut)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	if params != nil {
		q := request.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}

	request.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	return responseBody, response.StatusCode, nil
}


func (slf *VS) HttpPost(url string, header map[string]string, body map[string]interface{}, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error) {
	data, err := json.Marshal(body)
	if err  != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	return slf.HttpPostData(url, header, data, isSkipSecurityChecking, timeOut)
}


func (*VS) HttpPostData(url string, header map[string]string, data []byte, isSkipSecurityChecking bool, timeOut *time.Duration) ([]byte, int, error) {
	client := getHttpClient(isSkipSecurityChecking, timeOut)

	payload := bytes.NewBuffer(data)
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	request.Header.Set("Content-Type", "application/json")
	if header != nil {
		for key, val := range header {
			request.Header.Add(key, val)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.WithStack(err)
		return nil, -1, err
	}

	return responseBody, response.StatusCode, nil
}
