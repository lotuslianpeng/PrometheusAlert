package cloudauth

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

// OssUploadToken is a nested struct in cloudauth response
type OssUploadToken struct {
	Key      string `json:"Key" xml:"Key"`
	Bucket   string `json:"Bucket" xml:"Bucket"`
	Path     string `json:"Path" xml:"Path"`
	Token    string `json:"Token" xml:"Token"`
	EndPoint string `json:"EndPoint" xml:"EndPoint"`
	Expired  int64  `json:"Expired" xml:"Expired"`
	Secret   string `json:"Secret" xml:"Secret"`
}
