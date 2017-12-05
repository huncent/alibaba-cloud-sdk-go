package yundun

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

func (client *Client) Summary(request *SummaryRequest) (response *SummaryResponse, err error) {
	response = CreateSummaryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SummaryWithChan(request *SummaryRequest) (<-chan *SummaryResponse, <-chan error) {
	responseChan := make(chan *SummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Summary(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) SummaryWithCallback(request *SummaryRequest, callback func(response *SummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SummaryResponse
		var err error
		defer close(result)
		response, err = client.Summary(request)
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

type SummaryRequest struct {
	*requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
	JstOwnerId  string `position:"Query" name:"JstOwnerId"`
}

type SummaryResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Status            int64  `json:"Status" xml:"Status"`
	AbnormalHostCount int64  `json:"AbnormalHostCount" xml:"AbnormalHostCount"`
	Ddos              struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"Ddos" xml:"Ddos"`
	BruteForce struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"BruteForce" xml:"BruteForce"`
	Webshell struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"Webshell" xml:"Webshell"`
	RemoteLogin struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"RemoteLogin" xml:"RemoteLogin"`
	WebAttack struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"WebAttack" xml:"WebAttack"`
	WebLeak struct {
		Count     int64 `json:"Count" xml:"Count"`
		HostCount int64 `json:"HostCount" xml:"HostCount"`
	} `json:"WebLeak" xml:"WebLeak"`
}

func CreateSummaryRequest() (request *SummaryRequest) {
	request = &SummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun", "2015-04-16", "Summary", "", "")
	return
}

func CreateSummaryResponse() (response *SummaryResponse) {
	response = &SummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}