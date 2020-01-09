package openanalytics

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

// EndPointInfo is a nested struct in openanalytics response
type EndPointInfo struct {
	EndPointID      string `json:"endPointID" xml:"endPointID"`
	DomainURL       string `json:"domainURL" xml:"domainURL"`
	Host            string `json:"host" xml:"host"`
	Port            string `json:"port" xml:"port"`
	NetworkType     string `json:"networkType" xml:"networkType"`
	VpcID           string `json:"vpcID" xml:"vpcID"`
	VSwitch         string `json:"vSwitch" xml:"vSwitch"`
	Zone            string `json:"zone" xml:"zone"`
	AllowIP         string `json:"allowIP" xml:"allowIP"`
	CloudInstanceID string `json:"cloudInstanceID" xml:"cloudInstanceID"`
}
