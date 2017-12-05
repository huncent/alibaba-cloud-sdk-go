package aas

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

func (client *Client) GetShortTermAccessKeyForAccount(request *GetShortTermAccessKeyForAccountRequest) (response *GetShortTermAccessKeyForAccountResponse, err error) {
	response = CreateGetShortTermAccessKeyForAccountResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetShortTermAccessKeyForAccountWithChan(request *GetShortTermAccessKeyForAccountRequest) (<-chan *GetShortTermAccessKeyForAccountResponse, <-chan error) {
	responseChan := make(chan *GetShortTermAccessKeyForAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetShortTermAccessKeyForAccount(request)
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

func (client *Client) GetShortTermAccessKeyForAccountWithCallback(request *GetShortTermAccessKeyForAccountRequest, callback func(response *GetShortTermAccessKeyForAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetShortTermAccessKeyForAccountResponse
		var err error
		defer close(result)
		response, err = client.GetShortTermAccessKeyForAccount(request)
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

type GetShortTermAccessKeyForAccountRequest struct {
	*requests.RpcRequest
	ExpireTime   string `position:"Query" name:"ExpireTime"`
	IsMfaPresent string `position:"Query" name:"IsMfaPresent"`
	PK           string `position:"Query" name:"PK"`
}

type GetShortTermAccessKeyForAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PK        string `json:"PK" xml:"PK"`
	AccessKey struct {
		CreateTime      string `json:"CreateTime" xml:"CreateTime"`
		AccessKeyId     string `json:"AccessKeyId" xml:"AccessKeyId"`
		AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
		AccessKeyStatus string `json:"AccessKeyStatus" xml:"AccessKeyStatus"`
		AccessKeyType   string `json:"AccessKeyType" xml:"AccessKeyType"`
		ExpireTime      string `json:"ExpireTime" xml:"ExpireTime"`
	} `json:"AccessKey" xml:"AccessKey"`
}

func CreateGetShortTermAccessKeyForAccountRequest() (request *GetShortTermAccessKeyForAccountRequest) {
	request = &GetShortTermAccessKeyForAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Aas", "2015-07-01", "GetShortTermAccessKeyForAccount", "", "")
	return
}

func CreateGetShortTermAccessKeyForAccountResponse() (response *GetShortTermAccessKeyForAccountResponse) {
	response = &GetShortTermAccessKeyForAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}