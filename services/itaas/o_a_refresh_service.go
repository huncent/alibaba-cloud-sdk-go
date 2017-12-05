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

func (client *Client) OARefreshService(request *OARefreshServiceRequest) (response *OARefreshServiceResponse, err error) {
	response = CreateOARefreshServiceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OARefreshServiceWithChan(request *OARefreshServiceRequest) (<-chan *OARefreshServiceResponse, <-chan error) {
	responseChan := make(chan *OARefreshServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OARefreshService(request)
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

func (client *Client) OARefreshServiceWithCallback(request *OARefreshServiceRequest, callback func(response *OARefreshServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OARefreshServiceResponse
		var err error
		defer close(result)
		response, err = client.OARefreshService(request)
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

type OARefreshServiceRequest struct {
	*requests.RpcRequest
	Uid        string `position:"Query" name:"Uid"`
	Rid        string `position:"Query" name:"Rid"`
	AppVersion string `position:"Query" name:"AppVersion"`
	Language   string `position:"Query" name:"Language"`
	Umid       string `position:"Query" name:"Umid"`
	Cid        string `position:"Query" name:"Cid"`
}

type OARefreshServiceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	Success      bool   `json:"Success" xml:"Success"`
	ResultData   struct {
		Rid string `json:"Rid" xml:"Rid"`
		Sid string `json:"Sid" xml:"Sid"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateOARefreshServiceRequest() (request *OARefreshServiceRequest) {
	request = &OARefreshServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "OARefreshService", "", "")
	return
}

func CreateOARefreshServiceResponse() (response *OARefreshServiceResponse) {
	response = &OARefreshServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}