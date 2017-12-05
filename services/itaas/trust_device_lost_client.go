package itaas

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

func (client *Client) TrustDeviceLostClient(request *TrustDeviceLostClientRequest) (response *TrustDeviceLostClientResponse, err error) {
	response = CreateTrustDeviceLostClientResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) TrustDeviceLostClientWithChan(request *TrustDeviceLostClientRequest) (<-chan *TrustDeviceLostClientResponse, <-chan error) {
	responseChan := make(chan *TrustDeviceLostClientResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TrustDeviceLostClient(request)
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

func (client *Client) TrustDeviceLostClientWithCallback(request *TrustDeviceLostClientRequest, callback func(response *TrustDeviceLostClientResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TrustDeviceLostClientResponse
		var err error
		defer close(result)
		response, err = client.TrustDeviceLostClient(request)
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

type TrustDeviceLostClientRequest struct {
	*requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	OperateType   string `position:"Query" name:"OperateType"`
	Sid           string `position:"Query" name:"Sid"`
	AuthType      string `position:"Query" name:"AuthType"`
	OsType        string `position:"Query" name:"OsType"`
	AppVersion    string `position:"Query" name:"AppVersion"`
	Cid           string `position:"Query" name:"Cid"`
	AppName       string `position:"Query" name:"AppName"`
	QueryUmid     string `position:"Query" name:"QueryUmid"`
	Language      string `position:"Query" name:"Language"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Umid          string `position:"Query" name:"Umid"`
}

type TrustDeviceLostClientResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
		AdRefreshTokenId string `json:"AdRefreshTokenId" xml:"AdRefreshTokenId"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateTrustDeviceLostClientRequest() (request *TrustDeviceLostClientRequest) {
	request = &TrustDeviceLostClientRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "TrustDeviceLostClient", "", "")
	return
}

func CreateTrustDeviceLostClientResponse() (response *TrustDeviceLostClientResponse) {
	response = &TrustDeviceLostClientResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}