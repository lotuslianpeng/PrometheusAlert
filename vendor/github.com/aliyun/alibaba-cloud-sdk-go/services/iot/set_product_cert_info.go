package iot

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

// SetProductCertInfo invokes the iot.SetProductCertInfo API synchronously
// api document: https://help.aliyun.com/api/iot/setproductcertinfo.html
func (client *Client) SetProductCertInfo(request *SetProductCertInfoRequest) (response *SetProductCertInfoResponse, err error) {
	response = CreateSetProductCertInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SetProductCertInfoWithChan invokes the iot.SetProductCertInfo API asynchronously
// api document: https://help.aliyun.com/api/iot/setproductcertinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetProductCertInfoWithChan(request *SetProductCertInfoRequest) (<-chan *SetProductCertInfoResponse, <-chan error) {
	responseChan := make(chan *SetProductCertInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetProductCertInfo(request)
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

// SetProductCertInfoWithCallback invokes the iot.SetProductCertInfo API asynchronously
// api document: https://help.aliyun.com/api/iot/setproductcertinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetProductCertInfoWithCallback(request *SetProductCertInfoRequest, callback func(response *SetProductCertInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetProductCertInfoResponse
		var err error
		defer close(result)
		response, err = client.SetProductCertInfo(request)
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

// SetProductCertInfoRequest is the request struct for api SetProductCertInfo
type SetProductCertInfoRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	IssueModel    requests.Integer `position:"Query" name:"IssueModel"`
}

// SetProductCertInfoResponse is the response struct for api SetProductCertInfo
type SetProductCertInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateSetProductCertInfoRequest creates a request to invoke SetProductCertInfo API
func CreateSetProductCertInfoRequest() (request *SetProductCertInfoRequest) {
	request = &SetProductCertInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SetProductCertInfo", "iot", "openAPI")
	return
}

// CreateSetProductCertInfoResponse creates a response to parse from SetProductCertInfo response
func CreateSetProductCertInfoResponse() (response *SetProductCertInfoResponse) {
	response = &SetProductCertInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
