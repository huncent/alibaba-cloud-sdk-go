package domain

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

func (client *Client) SaveTaskForSubmittingDomainNameCredentialByTemplateId(request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) (response *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, err error) {
	response = CreateSaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveTaskForSubmittingDomainNameCredentialByTemplateIdWithChan(request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) (<-chan *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, <-chan error) {
	responseChan := make(chan *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveTaskForSubmittingDomainNameCredentialByTemplateId(request)
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

func (client *Client) SaveTaskForSubmittingDomainNameCredentialByTemplateIdWithCallback(request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest, callback func(response *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
		var err error
		defer close(result)
		response, err = client.SaveTaskForSubmittingDomainNameCredentialByTemplateId(request)
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

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest struct {
	*requests.RpcRequest
	DomainName        string `position:"Query" name:"DomainName"`
	SaleId            string `position:"Query" name:"SaleId"`
	ContactTemplateId string `position:"Query" name:"ContactTemplateId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

func CreateSaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest() (request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) {
	request = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForSubmittingDomainNameCredentialByTemplateId", "", "")
	return
}

func CreateSaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse() (response *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse) {
	response = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}