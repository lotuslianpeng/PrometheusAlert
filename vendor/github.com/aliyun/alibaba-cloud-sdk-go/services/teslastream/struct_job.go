package teslastream

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

// Job is a nested struct in teslastream response
type Job struct {
	Text           string   `json:"Text" xml:"Text"`
	TpsOut         int      `json:"TpsOut" xml:"TpsOut"`
	CpuUsed        float64  `json:"CpuUsed" xml:"CpuUsed"`
	Priority       string   `json:"Priority" xml:"Priority"`
	TpsIn          int      `json:"TpsIn" xml:"TpsIn"`
	Nickname       string   `json:"Nickname" xml:"Nickname"`
	CpuRequest     int      `json:"CpuRequest" xml:"CpuRequest"`
	MemUsed        float64  `json:"MemUsed" xml:"MemUsed"`
	PluginRelation string   `json:"PluginRelation" xml:"PluginRelation"`
	JobUniqKey     string   `json:"JobUniqKey" xml:"JobUniqKey"`
	Delay          int      `json:"Delay" xml:"Delay"`
	MemRequest     int      `json:"MemRequest" xml:"MemRequest"`
	Type           string   `json:"Type" xml:"Type"`
	Parents        []string `json:"Parents" xml:"Parents"`
	Childrens      []string `json:"Childrens" xml:"Childrens"`
}
