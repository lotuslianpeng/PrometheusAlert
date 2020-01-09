package cloudphoto

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

// CreatePhoto invokes the cloudphoto.CreatePhoto API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphoto.html
func (client *Client) CreatePhoto(request *CreatePhotoRequest) (response *CreatePhotoResponse, err error) {
	response = CreateCreatePhotoResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePhotoWithChan invokes the cloudphoto.CreatePhoto API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphoto.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhotoWithChan(request *CreatePhotoRequest) (<-chan *CreatePhotoResponse, <-chan error) {
	responseChan := make(chan *CreatePhotoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhoto(request)
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

// CreatePhotoWithCallback invokes the cloudphoto.CreatePhoto API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphoto.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhotoWithCallback(request *CreatePhotoRequest, callback func(response *CreatePhotoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhotoResponse
		var err error
		defer close(result)
		response, err = client.CreatePhoto(request)
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

// CreatePhotoRequest is the request struct for api CreatePhoto
type CreatePhotoRequest struct {
	*requests.RpcRequest
	TakenAt         requests.Integer `position:"Query" name:"TakenAt"`
	PhotoTitle      string           `position:"Query" name:"PhotoTitle"`
	LibraryId       string           `position:"Query" name:"LibraryId"`
	ShareExpireTime requests.Integer `position:"Query" name:"ShareExpireTime"`
	StoreName       string           `position:"Query" name:"StoreName"`
	UploadType      string           `position:"Query" name:"UploadType"`
	Remark          string           `position:"Query" name:"Remark"`
	SessionId       string           `position:"Query" name:"SessionId"`
	Staging         string           `position:"Query" name:"Staging"`
	FileId          string           `position:"Query" name:"FileId"`
}

// CreatePhotoResponse is the response struct for api CreatePhoto
type CreatePhotoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Photo     Photo  `json:"Photo" xml:"Photo"`
}

// CreateCreatePhotoRequest creates a request to invoke CreatePhoto API
func CreateCreatePhotoRequest() (request *CreatePhotoRequest) {
	request = &CreatePhotoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhoto", "cloudphoto", "openAPI")
	return
}

// CreateCreatePhotoResponse creates a response to parse from CreatePhoto response
func CreateCreatePhotoResponse() (response *CreatePhotoResponse) {
	response = &CreatePhotoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
