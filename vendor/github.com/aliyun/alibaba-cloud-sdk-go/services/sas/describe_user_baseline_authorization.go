package sas

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

// DescribeUserBaselineAuthorization invokes the sas.DescribeUserBaselineAuthorization API synchronously
// api document: https://help.aliyun.com/api/sas/describeuserbaselineauthorization.html
func (client *Client) DescribeUserBaselineAuthorization(request *DescribeUserBaselineAuthorizationRequest) (response *DescribeUserBaselineAuthorizationResponse, err error) {
	response = CreateDescribeUserBaselineAuthorizationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserBaselineAuthorizationWithChan invokes the sas.DescribeUserBaselineAuthorization API asynchronously
// api document: https://help.aliyun.com/api/sas/describeuserbaselineauthorization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserBaselineAuthorizationWithChan(request *DescribeUserBaselineAuthorizationRequest) (<-chan *DescribeUserBaselineAuthorizationResponse, <-chan error) {
	responseChan := make(chan *DescribeUserBaselineAuthorizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserBaselineAuthorization(request)
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

// DescribeUserBaselineAuthorizationWithCallback invokes the sas.DescribeUserBaselineAuthorization API asynchronously
// api document: https://help.aliyun.com/api/sas/describeuserbaselineauthorization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserBaselineAuthorizationWithCallback(request *DescribeUserBaselineAuthorizationRequest, callback func(response *DescribeUserBaselineAuthorizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserBaselineAuthorizationResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserBaselineAuthorization(request)
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

// DescribeUserBaselineAuthorizationRequest is the request struct for api DescribeUserBaselineAuthorization
type DescribeUserBaselineAuthorizationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
}

// DescribeUserBaselineAuthorizationResponse is the response struct for api DescribeUserBaselineAuthorization
type DescribeUserBaselineAuthorizationResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	UserBaselineAuthorization UserBaselineAuthorization `json:"UserBaselineAuthorization" xml:"UserBaselineAuthorization"`
}

// CreateDescribeUserBaselineAuthorizationRequest creates a request to invoke DescribeUserBaselineAuthorization API
func CreateDescribeUserBaselineAuthorizationRequest() (request *DescribeUserBaselineAuthorizationRequest) {
	request = &DescribeUserBaselineAuthorizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeUserBaselineAuthorization", "sas", "openAPI")
	return
}

// CreateDescribeUserBaselineAuthorizationResponse creates a response to parse from DescribeUserBaselineAuthorization response
func CreateDescribeUserBaselineAuthorizationResponse() (response *DescribeUserBaselineAuthorizationResponse) {
	response = &DescribeUserBaselineAuthorizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
