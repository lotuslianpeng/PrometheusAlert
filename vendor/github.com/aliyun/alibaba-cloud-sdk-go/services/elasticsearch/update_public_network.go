package elasticsearch

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

// UpdatePublicNetwork invokes the elasticsearch.UpdatePublicNetwork API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/updatepublicnetwork.html
func (client *Client) UpdatePublicNetwork(request *UpdatePublicNetworkRequest) (response *UpdatePublicNetworkResponse, err error) {
	response = CreateUpdatePublicNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePublicNetworkWithChan invokes the elasticsearch.UpdatePublicNetwork API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/updatepublicnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePublicNetworkWithChan(request *UpdatePublicNetworkRequest) (<-chan *UpdatePublicNetworkResponse, <-chan error) {
	responseChan := make(chan *UpdatePublicNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePublicNetwork(request)
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

// UpdatePublicNetworkWithCallback invokes the elasticsearch.UpdatePublicNetwork API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/updatepublicnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePublicNetworkWithCallback(request *UpdatePublicNetworkRequest, callback func(response *UpdatePublicNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePublicNetworkResponse
		var err error
		defer close(result)
		response, err = client.UpdatePublicNetwork(request)
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

// UpdatePublicNetworkRequest is the request struct for api UpdatePublicNetwork
type UpdatePublicNetworkRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// UpdatePublicNetworkResponse is the response struct for api UpdatePublicNetwork
type UpdatePublicNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdatePublicNetworkRequest creates a request to invoke UpdatePublicNetwork API
func CreateUpdatePublicNetworkRequest() (request *UpdatePublicNetworkRequest) {
	request = &UpdatePublicNetworkRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdatePublicNetwork", "/openapi/instances/[InstanceId]/public-network", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePublicNetworkResponse creates a response to parse from UpdatePublicNetwork response
func CreateUpdatePublicNetworkResponse() (response *UpdatePublicNetworkResponse) {
	response = &UpdatePublicNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
