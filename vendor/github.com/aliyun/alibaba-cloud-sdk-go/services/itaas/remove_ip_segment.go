package itaas

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

// RemoveIPSegment invokes the itaas.RemoveIPSegment API synchronously
// api document: https://help.aliyun.com/api/itaas/removeipsegment.html
func (client *Client) RemoveIPSegment(request *RemoveIPSegmentRequest) (response *RemoveIPSegmentResponse, err error) {
	response = CreateRemoveIPSegmentResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveIPSegmentWithChan invokes the itaas.RemoveIPSegment API asynchronously
// api document: https://help.aliyun.com/api/itaas/removeipsegment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveIPSegmentWithChan(request *RemoveIPSegmentRequest) (<-chan *RemoveIPSegmentResponse, <-chan error) {
	responseChan := make(chan *RemoveIPSegmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveIPSegment(request)
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

// RemoveIPSegmentWithCallback invokes the itaas.RemoveIPSegment API asynchronously
// api document: https://help.aliyun.com/api/itaas/removeipsegment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveIPSegmentWithCallback(request *RemoveIPSegmentRequest, callback func(response *RemoveIPSegmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveIPSegmentResponse
		var err error
		defer close(result)
		response, err = client.RemoveIPSegment(request)
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

// RemoveIPSegmentRequest is the request struct for api RemoveIPSegment
type RemoveIPSegmentRequest struct {
	*requests.RpcRequest
	Clientappid string `position:"Query" name:"Clientappid"`
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Uuid        string `position:"Query" name:"Uuid"`
	Operator    string `position:"Query" name:"Operator"`
}

// RemoveIPSegmentResponse is the response struct for api RemoveIPSegment
type RemoveIPSegmentResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	ErrorCode int                        `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string                     `json:"ErrorMsg" xml:"ErrorMsg"`
	Success   bool                       `json:"Success" xml:"Success"`
	ErrorList ErrorListInRemoveIPSegment `json:"ErrorList" xml:"ErrorList"`
}

// CreateRemoveIPSegmentRequest creates a request to invoke RemoveIPSegment API
func CreateRemoveIPSegmentRequest() (request *RemoveIPSegmentRequest) {
	request = &RemoveIPSegmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-05", "RemoveIPSegment", "itaas", "openAPI")
	return
}

// CreateRemoveIPSegmentResponse creates a response to parse from RemoveIPSegment response
func CreateRemoveIPSegmentResponse() (response *RemoveIPSegmentResponse) {
	response = &RemoveIPSegmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
