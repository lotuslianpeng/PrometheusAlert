package mts

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

// DeleteMcuJob invokes the mts.DeleteMcuJob API synchronously
// api document: https://help.aliyun.com/api/mts/deletemcujob.html
func (client *Client) DeleteMcuJob(request *DeleteMcuJobRequest) (response *DeleteMcuJobResponse, err error) {
	response = CreateDeleteMcuJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMcuJobWithChan invokes the mts.DeleteMcuJob API asynchronously
// api document: https://help.aliyun.com/api/mts/deletemcujob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMcuJobWithChan(request *DeleteMcuJobRequest) (<-chan *DeleteMcuJobResponse, <-chan error) {
	responseChan := make(chan *DeleteMcuJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMcuJob(request)
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

// DeleteMcuJobWithCallback invokes the mts.DeleteMcuJob API asynchronously
// api document: https://help.aliyun.com/api/mts/deletemcujob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMcuJobWithCallback(request *DeleteMcuJobRequest, callback func(response *DeleteMcuJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMcuJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteMcuJob(request)
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

// DeleteMcuJobRequest is the request struct for api DeleteMcuJob
type DeleteMcuJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteMcuJobResponse is the response struct for api DeleteMcuJob
type DeleteMcuJobResponse struct {
	*responses.BaseResponse
	RequestId      string                       `json:"RequestId" xml:"RequestId"`
	NonExistJobIds NonExistJobIdsInDeleteMcuJob `json:"NonExistJobIds" xml:"NonExistJobIds"`
	DeletedJobIds  DeletedJobIds                `json:"DeletedJobIds" xml:"DeletedJobIds"`
}

// CreateDeleteMcuJobRequest creates a request to invoke DeleteMcuJob API
func CreateDeleteMcuJobRequest() (request *DeleteMcuJobRequest) {
	request = &DeleteMcuJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DeleteMcuJob", "mts", "openAPI")
	return
}

// CreateDeleteMcuJobResponse creates a response to parse from DeleteMcuJob response
func CreateDeleteMcuJobResponse() (response *DeleteMcuJobResponse) {
	response = &DeleteMcuJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
