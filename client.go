/**
 * Created by elvizlai on 2016/2/16 17:03
 * Copyright © PubCloud
 */

package AliMNS

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"encoding/xml"
	"errors"
)

var httpClient = &http.Client{}

type aliClient struct {
	id     string
	secret string
	url    string
}

func NewClient(accessKeyId, accessKeySecret, url string) aliClient {
	return aliClient{id: accessKeyId, secret: accessKeySecret, url: url}
}

//send req to AliMNS
func (c aliClient) request(method string, path string, dataBytes []byte) (resp *http.Response, err error) {
	var req *http.Request

	if req, err = http.NewRequest(method, c.url + path, bytes.NewBuffer(dataBytes)); err != nil {
		return nil, err
	}

	headers := map[string]string{}
	headers[_MQ_VERSION] = version
	headers[_CONTENT_TYPE] = "application/xml"
	headers[_DATE] = time.Now().UTC().Format(http.TimeFormat)

	if dataBytes != nil {
		headers[_CONTENT_MD5] = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", md5.Sum(dataBytes))))
	}

	sign, err := signature(c.secret, method, headers, path)
	if err != nil {
		return nil, err
	}

	headers[_AUTHORIZATION] = "MNS " + c.id + ":" + sign

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	return httpClient.Do(req)
}

//handle resp from AliMNS
func (c aliClient) respHandler(method string, path string, dataBytes []byte, v interface{}) (err error) {
	resp, err := c.request(method, path, dataBytes)
	if err != nil {
		return err
	}

	//error
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK &&resp.StatusCode != http.StatusNoContent {
		decoder := xml.NewDecoder(resp.Body)
		e := ErrorResponse{}
		decoder.Decode(&e)
		return errors.New(e.Message)
	}

	if v != nil {
		decoder := xml.NewDecoder(resp.Body)
		return decoder.Decode(v)
	}

	return
}