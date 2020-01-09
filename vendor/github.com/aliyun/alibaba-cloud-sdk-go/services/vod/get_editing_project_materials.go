package vod

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

// GetEditingProjectMaterials invokes the vod.GetEditingProjectMaterials API synchronously
// api document: https://help.aliyun.com/api/vod/geteditingprojectmaterials.html
func (client *Client) GetEditingProjectMaterials(request *GetEditingProjectMaterialsRequest) (response *GetEditingProjectMaterialsResponse, err error) {
	response = CreateGetEditingProjectMaterialsResponse()
	err = client.DoAction(request, response)
	return
}

// GetEditingProjectMaterialsWithChan invokes the vod.GetEditingProjectMaterials API asynchronously
// api document: https://help.aliyun.com/api/vod/geteditingprojectmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEditingProjectMaterialsWithChan(request *GetEditingProjectMaterialsRequest) (<-chan *GetEditingProjectMaterialsResponse, <-chan error) {
	responseChan := make(chan *GetEditingProjectMaterialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEditingProjectMaterials(request)
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

// GetEditingProjectMaterialsWithCallback invokes the vod.GetEditingProjectMaterials API asynchronously
// api document: https://help.aliyun.com/api/vod/geteditingprojectmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEditingProjectMaterialsWithCallback(request *GetEditingProjectMaterialsRequest, callback func(response *GetEditingProjectMaterialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEditingProjectMaterialsResponse
		var err error
		defer close(result)
		response, err = client.GetEditingProjectMaterials(request)
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

// GetEditingProjectMaterialsRequest is the request struct for api GetEditingProjectMaterials
type GetEditingProjectMaterialsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Type                 string `position:"Query" name:"Type"`
	MaterialType         string `position:"Query" name:"MaterialType"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// GetEditingProjectMaterialsResponse is the response struct for api GetEditingProjectMaterials
type GetEditingProjectMaterialsResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	MaterialList MaterialList `json:"MaterialList" xml:"MaterialList"`
}

// CreateGetEditingProjectMaterialsRequest creates a request to invoke GetEditingProjectMaterials API
func CreateGetEditingProjectMaterialsRequest() (request *GetEditingProjectMaterialsRequest) {
	request = &GetEditingProjectMaterialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetEditingProjectMaterials", "vod", "openAPI")
	return
}

// CreateGetEditingProjectMaterialsResponse creates a response to parse from GetEditingProjectMaterials response
func CreateGetEditingProjectMaterialsResponse() (response *GetEditingProjectMaterialsResponse) {
	response = &GetEditingProjectMaterialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
