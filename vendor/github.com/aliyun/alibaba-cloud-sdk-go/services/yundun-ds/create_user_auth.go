package yundun_ds

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

// CreateUserAuth invokes the yundun_ds.CreateUserAuth API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/createuserauth.html
func (client *Client) CreateUserAuth(request *CreateUserAuthRequest) (response *CreateUserAuthResponse, err error) {
	response = CreateCreateUserAuthResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserAuthWithChan invokes the yundun_ds.CreateUserAuth API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/createuserauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserAuthWithChan(request *CreateUserAuthRequest) (<-chan *CreateUserAuthResponse, <-chan error) {
	responseChan := make(chan *CreateUserAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserAuth(request)
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

// CreateUserAuthWithCallback invokes the yundun_ds.CreateUserAuth API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/createuserauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserAuthWithCallback(request *CreateUserAuthRequest, callback func(response *CreateUserAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserAuthResponse
		var err error
		defer close(result)
		response, err = client.CreateUserAuth(request)
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

// CreateUserAuthRequest is the request struct for api CreateUserAuth
type CreateUserAuthRequest struct {
	*requests.RpcRequest
	AccountId       requests.Integer `position:"Query" name:"AccountId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	AccessKey       string           `position:"Query" name:"AccessKey"`
	AccessKeySecret string           `position:"Query" name:"AccessKeySecret"`
	Lang            string           `position:"Query" name:"Lang"`
}

// CreateUserAuthResponse is the response struct for api CreateUserAuth
type CreateUserAuthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateUserAuthRequest creates a request to invoke CreateUserAuth API
func CreateCreateUserAuthRequest() (request *CreateUserAuthRequest) {
	request = &CreateUserAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "CreateUserAuth", "sddp", "openAPI")
	return
}

// CreateCreateUserAuthResponse creates a response to parse from CreateUserAuth response
func CreateCreateUserAuthResponse() (response *CreateUserAuthResponse) {
	response = &CreateUserAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
