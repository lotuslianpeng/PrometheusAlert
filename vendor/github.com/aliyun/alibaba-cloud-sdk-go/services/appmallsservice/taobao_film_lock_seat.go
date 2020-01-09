package appmallsservice

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

// TaobaoFilmLockSeat invokes the appmallsservice.TaobaoFilmLockSeat API synchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmlockseat.html
func (client *Client) TaobaoFilmLockSeat(request *TaobaoFilmLockSeatRequest) (response *TaobaoFilmLockSeatResponse, err error) {
	response = CreateTaobaoFilmLockSeatResponse()
	err = client.DoAction(request, response)
	return
}

// TaobaoFilmLockSeatWithChan invokes the appmallsservice.TaobaoFilmLockSeat API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmlockseat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmLockSeatWithChan(request *TaobaoFilmLockSeatRequest) (<-chan *TaobaoFilmLockSeatResponse, <-chan error) {
	responseChan := make(chan *TaobaoFilmLockSeatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaobaoFilmLockSeat(request)
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

// TaobaoFilmLockSeatWithCallback invokes the appmallsservice.TaobaoFilmLockSeat API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmlockseat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmLockSeatWithCallback(request *TaobaoFilmLockSeatRequest, callback func(response *TaobaoFilmLockSeatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaobaoFilmLockSeatResponse
		var err error
		defer close(result)
		response, err = client.TaobaoFilmLockSeat(request)
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

// TaobaoFilmLockSeatRequest is the request struct for api TaobaoFilmLockSeat
type TaobaoFilmLockSeatRequest struct {
	*requests.RpcRequest
	ScheduleId requests.Integer `position:"Query" name:"ScheduleId"`
	SeatIds    string           `position:"Query" name:"SeatIds"`
	SeatNames  string           `position:"Query" name:"SeatNames"`
	Mobile     string           `position:"Query" name:"Mobile"`
	ExtUserId  string           `position:"Query" name:"ExtUserId"`
	ParamsJson string           `position:"Query" name:"ParamsJson"`
}

// TaobaoFilmLockSeatResponse is the response struct for api TaobaoFilmLockSeat
type TaobaoFilmLockSeatResponse struct {
	*responses.BaseResponse
	ErrorCode  string     `json:"ErrorCode" xml:"ErrorCode"`
	Msg        string     `json:"Msg" xml:"Msg"`
	SubCode    string     `json:"SubCode" xml:"SubCode"`
	SubMsg     string     `json:"SubMsg" xml:"SubMsg"`
	LogsId     string     `json:"LogsId" xml:"LogsId"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	SeatLocked SeatLocked `json:"SeatLocked" xml:"SeatLocked"`
}

// CreateTaobaoFilmLockSeatRequest creates a request to invoke TaobaoFilmLockSeat API
func CreateTaobaoFilmLockSeatRequest() (request *TaobaoFilmLockSeatRequest) {
	request = &TaobaoFilmLockSeatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AppMallsService", "2018-02-24", "TaobaoFilmLockSeat", "", "")
	return
}

// CreateTaobaoFilmLockSeatResponse creates a response to parse from TaobaoFilmLockSeat response
func CreateTaobaoFilmLockSeatResponse() (response *TaobaoFilmLockSeatResponse) {
	response = &TaobaoFilmLockSeatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
