package aegis

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

// RollbackQuaraFile invokes the aegis.RollbackQuaraFile API synchronously
// api document: https://help.aliyun.com/api/aegis/rollbackquarafile.html
func (client *Client) RollbackQuaraFile(request *RollbackQuaraFileRequest) (response *RollbackQuaraFileResponse, err error) {
	response = CreateRollbackQuaraFileResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackQuaraFileWithChan invokes the aegis.RollbackQuaraFile API asynchronously
// api document: https://help.aliyun.com/api/aegis/rollbackquarafile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackQuaraFileWithChan(request *RollbackQuaraFileRequest) (<-chan *RollbackQuaraFileResponse, <-chan error) {
	responseChan := make(chan *RollbackQuaraFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackQuaraFile(request)
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

// RollbackQuaraFileWithCallback invokes the aegis.RollbackQuaraFile API asynchronously
// api document: https://help.aliyun.com/api/aegis/rollbackquarafile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackQuaraFileWithCallback(request *RollbackQuaraFileRequest, callback func(response *RollbackQuaraFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackQuaraFileResponse
		var err error
		defer close(result)
		response, err = client.RollbackQuaraFile(request)
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

// RollbackQuaraFileRequest is the request struct for api RollbackQuaraFile
type RollbackQuaraFileRequest struct {
	*requests.RpcRequest
	SourceIp  string `position:"Query" name:"SourceIp"`
	EventType string `position:"Query" name:"EventType"`
	Tag       string `position:"Query" name:"Tag"`
	Uuid      string `position:"Query" name:"Uuid"`
	EventName string `position:"Query" name:"EventName"`
}

// RollbackQuaraFileResponse is the response struct for api RollbackQuaraFile
type RollbackQuaraFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateRollbackQuaraFileRequest creates a request to invoke RollbackQuaraFile API
func CreateRollbackQuaraFileRequest() (request *RollbackQuaraFileRequest) {
	request = &RollbackQuaraFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "RollbackQuaraFile", "vipaegis", "openAPI")
	return
}

// CreateRollbackQuaraFileResponse creates a response to parse from RollbackQuaraFile response
func CreateRollbackQuaraFileResponse() (response *RollbackQuaraFileResponse) {
	response = &RollbackQuaraFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
