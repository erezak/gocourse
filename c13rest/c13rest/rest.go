// A package for rest CRUD
//
// Part of the gocourse
package c13rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Rest client
//
// Can only be created by the NewRestClient constructor
type RestClient struct {
	hostUrl  string
	basePath string
}

// Constructor for RestClient
// Parameters:
//
// hostUrl string - url including scheme
// basePath string - basePath for all resources
func NewRestClient(hostUrl string, basePath string) RestClient{
	return RestClient{hostUrl: hostUrl,
		basePath: basePath}
}


func (r *RestClient) Create(resourceType string, resourceData []byte) (statusCode int,
	status string,
	bodyResponse []byte,
	err error) {
	statusCode, status, bodyResponse, err = r.doRequest("POST", r.constructResourcePath(resourceType), resourceData)
	return
}

func (r *RestClient) Delete(resourceType string, resourceId string) (statusCode int,
	status string,
	bodyResponse []byte,
	err error) {
	statusCode, status, bodyResponse, err = r.doRequest("DELETE", r.constructResourcePath(resourceType, resourceId), nil)
	return
}

func (r *RestClient) doRequest(method string, resourcePath string, resourceData []byte) (statusCode int,
	status string,
	bodyResponse []byte,
	err error) {
	var reader io.Reader
	if resourceData != nil {
		reader = bytes.NewReader(resourceData)
	}
	request, err := http.NewRequest(method, resourcePath, reader)
	if err != nil {
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	statusCode = response.StatusCode
	status = response.Status
	bodyResponse, err = ioutil.ReadAll(response.Body)
	return
}

func (r *RestClient) constructResourcePath(resources ...string) (resourcePath string) {
	resourcePath = r.hostUrl
	// Check if hostUrl and basePath end with /
	if !strings.HasSuffix(resourcePath, "/") {
		resourcePath += "/"
	}

	resourcePath += r.basePath

	for _, resource := range resources {
		if !strings.HasSuffix(resourcePath, "/") {
			resourcePath += "/"
		}
		resourcePath += resource
	}

	return
}
