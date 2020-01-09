package domain

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryDomainSuffix invokes the domain.QueryDomainSuffix API synchronously
// api document: https://help.aliyun.com/api/domain/querydomainsuffix.html
func (client *Client) QueryDomainSuffix(request *QueryDomainSuffixRequest) (response *QueryDomainSuffixResponse, err error) {
	response = CreateQueryDomainSuffixResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainSuffixWithChan invokes the domain.QueryDomainSuffix API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomainsuffix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainSuffixWithChan(request *QueryDomainSuffixRequest) (<-chan *QueryDomainSuffixResponse, <-chan error) {
	responseChan := make(chan *QueryDomainSuffixResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainSuffix(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryDomainSuffixWithCallback invokes the domain.QueryDomainSuffix API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomainsuffix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainSuffixWithCallback(request *QueryDomainSuffixRequest, callback func(response *QueryDomainSuffixResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainSuffixResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainSuffix(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryDomainSuffixRequest is the request struct for api QueryDomainSuffix
type QueryDomainSuffixRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDomainSuffixResponse is the response struct for api QueryDomainSuffix
type QueryDomainSuffixResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	SuffixList SuffixList `json:"SuffixList" xml:"SuffixList"`
}

// CreateQueryDomainSuffixRequest creates a request to invoke QueryDomainSuffix API
func CreateQueryDomainSuffixRequest() (request *QueryDomainSuffixRequest) {
	request = &QueryDomainSuffixRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainSuffix", "", "")
	return
}

// CreateQueryDomainSuffixResponse creates a response to parse from QueryDomainSuffix response
func CreateQueryDomainSuffixResponse() (response *QueryDomainSuffixResponse) {
	response = &QueryDomainSuffixResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
