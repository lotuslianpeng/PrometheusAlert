package sddp

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

// Instance is a nested struct in sddp response
type Instance struct {
	Id                int    `json:"Id" xml:"Id"`
	Name              string `json:"Name" xml:"Name"`
	Owner             string `json:"Owner" xml:"Owner"`
	CreationTime      int    `json:"CreationTime" xml:"CreationTime"`
	ProductId         string `json:"ProductId" xml:"ProductId"`
	ProductCode       string `json:"ProductCode" xml:"ProductCode"`
	Protection        bool   `json:"Protection" xml:"Protection"`
	Labelsec          int    `json:"Labelsec" xml:"Labelsec"`
	OdpsRiskLevelName string `json:"OdpsRiskLevelName" xml:"OdpsRiskLevelName"`
	Sensitive         bool   `json:"Sensitive" xml:"Sensitive"`
	RiskLevelId       int    `json:"RiskLevelId" xml:"RiskLevelId"`
	RiskLevelName     string `json:"RiskLevelName" xml:"RiskLevelName"`
	RuleName          string `json:"RuleName" xml:"RuleName"`
	DepartName        string `json:"DepartName" xml:"DepartName"`
	TotalCount        int    `json:"TotalCount" xml:"TotalCount"`
	SensitiveCount    int    `json:"SensitiveCount" xml:"SensitiveCount"`
	Acl               string `json:"Acl" xml:"Acl"`
}
