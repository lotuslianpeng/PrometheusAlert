package ocr

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

// RecognizeAccountPage invokes the ocr.RecognizeAccountPage API synchronously
// api document: https://help.aliyun.com/api/ocr/recognizeaccountpage.html
func (client *Client) RecognizeAccountPage(request *RecognizeAccountPageRequest) (response *RecognizeAccountPageResponse, err error) {
	response = CreateRecognizeAccountPageResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeAccountPageWithChan invokes the ocr.RecognizeAccountPage API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizeaccountpage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeAccountPageWithChan(request *RecognizeAccountPageRequest) (<-chan *RecognizeAccountPageResponse, <-chan error) {
	responseChan := make(chan *RecognizeAccountPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeAccountPage(request)
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

// RecognizeAccountPageWithCallback invokes the ocr.RecognizeAccountPage API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizeaccountpage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeAccountPageWithCallback(request *RecognizeAccountPageRequest, callback func(response *RecognizeAccountPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeAccountPageResponse
		var err error
		defer close(result)
		response, err = client.RecognizeAccountPage(request)
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

// RecognizeAccountPageRequest is the request struct for api RecognizeAccountPage
type RecognizeAccountPageRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// RecognizeAccountPageResponse is the response struct for api RecognizeAccountPage
type RecognizeAccountPageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeAccountPageRequest creates a request to invoke RecognizeAccountPage API
func CreateRecognizeAccountPageRequest() (request *RecognizeAccountPageRequest) {
	request = &RecognizeAccountPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeAccountPage", "ocr", "openAPI")
	return
}

// CreateRecognizeAccountPageResponse creates a response to parse from RecognizeAccountPage response
func CreateRecognizeAccountPageResponse() (response *RecognizeAccountPageResponse) {
	response = &RecognizeAccountPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
