package qualitycheck

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

// ListPrecisionTask invokes the qualitycheck.ListPrecisionTask API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/listprecisiontask.html
func (client *Client) ListPrecisionTask(request *ListPrecisionTaskRequest) (response *ListPrecisionTaskResponse, err error) {
	response = CreateListPrecisionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrecisionTaskWithChan invokes the qualitycheck.ListPrecisionTask API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/listprecisiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrecisionTaskWithChan(request *ListPrecisionTaskRequest) (<-chan *ListPrecisionTaskResponse, <-chan error) {
	responseChan := make(chan *ListPrecisionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrecisionTask(request)
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

// ListPrecisionTaskWithCallback invokes the qualitycheck.ListPrecisionTask API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/listprecisiontask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrecisionTaskWithCallback(request *ListPrecisionTaskRequest, callback func(response *ListPrecisionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrecisionTaskResponse
		var err error
		defer close(result)
		response, err = client.ListPrecisionTask(request)
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

// ListPrecisionTaskRequest is the request struct for api ListPrecisionTask
type ListPrecisionTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// ListPrecisionTaskResponse is the response struct for api ListPrecisionTask
type ListPrecisionTaskResponse struct {
	*responses.BaseResponse
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	Success    bool                    `json:"Success" xml:"Success"`
	Code       string                  `json:"Code" xml:"Code"`
	Message    string                  `json:"Message" xml:"Message"`
	PageNumber int                     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                     `json:"PageSize" xml:"PageSize"`
	Count      int                     `json:"Count" xml:"Count"`
	Data       DataInListPrecisionTask `json:"Data" xml:"Data"`
}

// CreateListPrecisionTaskRequest creates a request to invoke ListPrecisionTask API
func CreateListPrecisionTaskRequest() (request *ListPrecisionTaskRequest) {
	request = &ListPrecisionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "ListPrecisionTask", "", "")
	return
}

// CreateListPrecisionTaskResponse creates a response to parse from ListPrecisionTask response
func CreateListPrecisionTaskResponse() (response *ListPrecisionTaskResponse) {
	response = &ListPrecisionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
