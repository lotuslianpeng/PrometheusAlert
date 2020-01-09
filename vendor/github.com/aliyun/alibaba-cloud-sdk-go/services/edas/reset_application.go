package edas

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

// ResetApplication invokes the edas.ResetApplication API synchronously
// api document: https://help.aliyun.com/api/edas/resetapplication.html
func (client *Client) ResetApplication(request *ResetApplicationRequest) (response *ResetApplicationResponse, err error) {
	response = CreateResetApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// ResetApplicationWithChan invokes the edas.ResetApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/resetapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApplicationWithChan(request *ResetApplicationRequest) (<-chan *ResetApplicationResponse, <-chan error) {
	responseChan := make(chan *ResetApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetApplication(request)
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

// ResetApplicationWithCallback invokes the edas.ResetApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/resetapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApplicationWithCallback(request *ResetApplicationRequest, callback func(response *ResetApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetApplicationResponse
		var err error
		defer close(result)
		response, err = client.ResetApplication(request)
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

// ResetApplicationRequest is the request struct for api ResetApplication
type ResetApplicationRequest struct {
	*requests.RoaRequest
	AppId   string `position:"Query" name:"AppId"`
	EccInfo string `position:"Query" name:"EccInfo"`
}

// ResetApplicationResponse is the response struct for api ResetApplication
type ResetApplicationResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateResetApplicationRequest creates a request to invoke ResetApplication API
func CreateResetApplicationRequest() (request *ResetApplicationRequest) {
	request = &ResetApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ResetApplication", "/pop/v5/changeorder/co_reset", "", "")
	request.Method = requests.POST
	return
}

// CreateResetApplicationResponse creates a response to parse from ResetApplication response
func CreateResetApplicationResponse() (response *ResetApplicationResponse) {
	response = &ResetApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
