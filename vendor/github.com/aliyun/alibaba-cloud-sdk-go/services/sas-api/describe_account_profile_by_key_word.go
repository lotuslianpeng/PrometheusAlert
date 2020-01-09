package sas_api

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

// DescribeAccountProfileByKeyWord invokes the sas_api.DescribeAccountProfileByKeyWord API synchronously
// api document: https://help.aliyun.com/api/sas-api/describeaccountprofilebykeyword.html
func (client *Client) DescribeAccountProfileByKeyWord(request *DescribeAccountProfileByKeyWordRequest) (response *DescribeAccountProfileByKeyWordResponse, err error) {
	response = CreateDescribeAccountProfileByKeyWordResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccountProfileByKeyWordWithChan invokes the sas_api.DescribeAccountProfileByKeyWord API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describeaccountprofilebykeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccountProfileByKeyWordWithChan(request *DescribeAccountProfileByKeyWordRequest) (<-chan *DescribeAccountProfileByKeyWordResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountProfileByKeyWordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccountProfileByKeyWord(request)
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

// DescribeAccountProfileByKeyWordWithCallback invokes the sas_api.DescribeAccountProfileByKeyWord API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describeaccountprofilebykeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccountProfileByKeyWordWithCallback(request *DescribeAccountProfileByKeyWordRequest, callback func(response *DescribeAccountProfileByKeyWordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountProfileByKeyWordResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccountProfileByKeyWord(request)
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

// DescribeAccountProfileByKeyWordRequest is the request struct for api DescribeAccountProfileByKeyWord
type DescribeAccountProfileByKeyWordRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Keyword  string `position:"Query" name:"Keyword"`
}

// DescribeAccountProfileByKeyWordResponse is the response struct for api DescribeAccountProfileByKeyWord
type DescribeAccountProfileByKeyWordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Ip        string `json:"Ip" xml:"Ip"`
	IpInfo    string `json:"IpInfo" xml:"IpInfo"`
	Phone     string `json:"Phone" xml:"Phone"`
	PhoneInfo string `json:"PhoneInfo" xml:"PhoneInfo"`
}

// CreateDescribeAccountProfileByKeyWordRequest creates a request to invoke DescribeAccountProfileByKeyWord API
func CreateDescribeAccountProfileByKeyWordRequest() (request *DescribeAccountProfileByKeyWordRequest) {
	request = &DescribeAccountProfileByKeyWordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas-api", "2017-07-05", "DescribeAccountProfileByKeyWord", "sas-api", "openAPI")
	return
}

// CreateDescribeAccountProfileByKeyWordResponse creates a response to parse from DescribeAccountProfileByKeyWord response
func CreateDescribeAccountProfileByKeyWordResponse() (response *DescribeAccountProfileByKeyWordResponse) {
	response = &DescribeAccountProfileByKeyWordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
