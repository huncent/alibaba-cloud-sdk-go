package ros

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

func (client *Client) WaitConditions(request *WaitConditionsRequest) (response *WaitConditionsResponse, err error) {
	response = CreateWaitConditionsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) WaitConditionsWithChan(request *WaitConditionsRequest) (<-chan *WaitConditionsResponse, <-chan error) {
	responseChan := make(chan *WaitConditionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WaitConditions(request)
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

func (client *Client) WaitConditionsWithCallback(request *WaitConditionsRequest, callback func(response *WaitConditionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WaitConditionsResponse
		var err error
		defer close(result)
		response, err = client.WaitConditions(request)
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

type WaitConditionsRequest struct {
	*requests.RoaRequest
	Expire    string `position:"Query" name:"expire"`
	Resource  string `position:"Query" name:"resource"`
	Stackname string `position:"Query" name:"stackname"`
	Signature string `position:"Query" name:"signature"`
	Stackid   string `position:"Query" name:"stackid"`
}

type WaitConditionsResponse struct {
	*responses.BaseResponse
}

func CreateWaitConditionsRequest() (request *WaitConditionsRequest) {
	request = &WaitConditionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "WaitConditions", "/waitcondition", "", "")
	return
}

func CreateWaitConditionsResponse() (response *WaitConditionsResponse) {
	response = &WaitConditionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}