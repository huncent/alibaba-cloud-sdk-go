package cdn

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

func (client *Client) DescribeLiveStreamRoomUserNumber(request *DescribeLiveStreamRoomUserNumberRequest) (response *DescribeLiveStreamRoomUserNumberResponse, err error) {
	response = CreateDescribeLiveStreamRoomUserNumberResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRoomUserNumberWithChan(request *DescribeLiveStreamRoomUserNumberRequest) (<-chan *DescribeLiveStreamRoomUserNumberResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRoomUserNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRoomUserNumber(request)
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

func (client *Client) DescribeLiveStreamRoomUserNumberWithCallback(request *DescribeLiveStreamRoomUserNumberRequest, callback func(response *DescribeLiveStreamRoomUserNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRoomUserNumberResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRoomUserNumber(request)
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

type DescribeLiveStreamRoomUserNumberRequest struct {
	*requests.RpcRequest
	HlsSwitch     string           `position:"Query" name:"HlsSwitch"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamRoomUserNumberResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TotalUserNumber int    `json:"TotalUserNumber" xml:"TotalUserNumber"`
	OnlineUserInfo  struct {
		LiveStreamOnlineUserNumInfo []struct {
			StreamUrl  string `json:"StreamUrl" xml:"StreamUrl"`
			UserNumber int    `json:"UserNumber" xml:"UserNumber"`
			Time       string `json:"Time" xml:"Time"`
		} `json:"LiveStreamOnlineUserNumInfo" xml:"LiveStreamOnlineUserNumInfo"`
	} `json:"OnlineUserInfo" xml:"OnlineUserInfo"`
}

func CreateDescribeLiveStreamRoomUserNumberRequest() (request *DescribeLiveStreamRoomUserNumberRequest) {
	request = &DescribeLiveStreamRoomUserNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRoomUserNumber", "", "")
	return
}

func CreateDescribeLiveStreamRoomUserNumberResponse() (response *DescribeLiveStreamRoomUserNumberResponse) {
	response = &DescribeLiveStreamRoomUserNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
