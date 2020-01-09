package rds

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

// CheckCreateDdrDBInstance invokes the rds.CheckCreateDdrDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/checkcreateddrdbinstance.html
func (client *Client) CheckCreateDdrDBInstance(request *CheckCreateDdrDBInstanceRequest) (response *CheckCreateDdrDBInstanceResponse, err error) {
	response = CreateCheckCreateDdrDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CheckCreateDdrDBInstanceWithChan invokes the rds.CheckCreateDdrDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/checkcreateddrdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckCreateDdrDBInstanceWithChan(request *CheckCreateDdrDBInstanceRequest) (<-chan *CheckCreateDdrDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CheckCreateDdrDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckCreateDdrDBInstance(request)
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

// CheckCreateDdrDBInstanceWithCallback invokes the rds.CheckCreateDdrDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/checkcreateddrdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckCreateDdrDBInstanceWithCallback(request *CheckCreateDdrDBInstanceRequest, callback func(response *CheckCreateDdrDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckCreateDdrDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CheckCreateDdrDBInstance(request)
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

// CheckCreateDdrDBInstanceRequest is the request struct for api CheckCreateDdrDBInstance
type CheckCreateDdrDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    requests.Integer `position:"Query" name:"DBInstanceStorage"`
	SourceDBInstanceName string           `position:"Query" name:"SourceDBInstanceName"`
	HostType             string           `position:"Query" name:"HostType"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	UserBakSetURL        string           `position:"Query" name:"UserBakSetURL"`
	Engine               string           `position:"Query" name:"Engine"`
	BackupSetRegion      string           `position:"Query" name:"BackupSetRegion"`
	BackupSetType        string           `position:"Query" name:"BackupSetType"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	BakSetName           string           `position:"Query" name:"BakSetName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupSetId          string           `position:"Query" name:"BackupSetId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	RestoreType          string           `position:"Query" name:"RestoreType"`
	SourceRegion         string           `position:"Query" name:"SourceRegion"`
}

// CheckCreateDdrDBInstanceResponse is the response struct for api CheckCreateDdrDBInstance
type CheckCreateDdrDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IsValid   string `json:"IsValid" xml:"IsValid"`
}

// CreateCheckCreateDdrDBInstanceRequest creates a request to invoke CheckCreateDdrDBInstance API
func CreateCheckCreateDdrDBInstanceRequest() (request *CheckCreateDdrDBInstanceRequest) {
	request = &CheckCreateDdrDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CheckCreateDdrDBInstance", "rds", "openAPI")
	return
}

// CreateCheckCreateDdrDBInstanceResponse creates a response to parse from CheckCreateDdrDBInstance response
func CreateCheckCreateDdrDBInstanceResponse() (response *CheckCreateDdrDBInstanceResponse) {
	response = &CheckCreateDdrDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
