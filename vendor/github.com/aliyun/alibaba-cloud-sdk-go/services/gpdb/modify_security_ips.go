package gpdb

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

// ModifySecurityIps invokes the gpdb.ModifySecurityIps API synchronously
// api document: https://help.aliyun.com/api/gpdb/modifysecurityips.html
func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (response *ModifySecurityIpsResponse, err error) {
	response = CreateModifySecurityIpsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityIpsWithChan invokes the gpdb.ModifySecurityIps API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifysecurityips.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityIpsWithChan(request *ModifySecurityIpsRequest) (<-chan *ModifySecurityIpsResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityIps(request)
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

// ModifySecurityIpsWithCallback invokes the gpdb.ModifySecurityIps API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifysecurityips.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityIpsWithCallback(request *ModifySecurityIpsRequest, callback func(response *ModifySecurityIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityIpsResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityIps(request)
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

// ModifySecurityIpsRequest is the request struct for api ModifySecurityIps
type ModifySecurityIpsRequest struct {
	*requests.RpcRequest
	SecurityIPList             string `position:"Query" name:"SecurityIPList"`
	DBInstanceIPArrayName      string `position:"Query" name:"DBInstanceIPArrayName"`
	DBInstanceIPArrayAttribute string `position:"Query" name:"DBInstanceIPArrayAttribute"`
	DBInstanceId               string `position:"Query" name:"DBInstanceId"`
}

// ModifySecurityIpsResponse is the response struct for api ModifySecurityIps
type ModifySecurityIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityIpsRequest creates a request to invoke ModifySecurityIps API
func CreateModifySecurityIpsRequest() (request *ModifySecurityIpsRequest) {
	request = &ModifySecurityIpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifySecurityIps", "gpdb", "openAPI")
	return
}

// CreateModifySecurityIpsResponse creates a response to parse from ModifySecurityIps response
func CreateModifySecurityIpsResponse() (response *ModifySecurityIpsResponse) {
	response = &ModifySecurityIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
