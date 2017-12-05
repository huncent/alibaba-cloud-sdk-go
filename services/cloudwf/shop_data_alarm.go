package cloudwf

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

func (client *Client) ShopDataAlarm(request *ShopDataAlarmRequest) (response *ShopDataAlarmResponse, err error) {
	response = CreateShopDataAlarmResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ShopDataAlarmWithChan(request *ShopDataAlarmRequest) (<-chan *ShopDataAlarmResponse, <-chan error) {
	responseChan := make(chan *ShopDataAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopDataAlarm(request)
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

func (client *Client) ShopDataAlarmWithCallback(request *ShopDataAlarmRequest, callback func(response *ShopDataAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopDataAlarmResponse
		var err error
		defer close(result)
		response, err = client.ShopDataAlarm(request)
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

type ShopDataAlarmRequest struct {
	*requests.RpcRequest
	Warn      string `position:"Query" name:"Warn"`
	WarnEmail string `position:"Query" name:"WarnEmail"`
	Sid       string `position:"Query" name:"Sid"`
	WarnPhone string `position:"Query" name:"WarnPhone"`
	CloseWarn string `position:"Query" name:"CloseWarn"`
}

type ShopDataAlarmResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateShopDataAlarmRequest() (request *ShopDataAlarmRequest) {
	request = &ShopDataAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDataAlarm", "", "")
	return
}

func CreateShopDataAlarmResponse() (response *ShopDataAlarmResponse) {
	response = &ShopDataAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}