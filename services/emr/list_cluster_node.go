package emr

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

func (client *Client) ListClusterNode(request *ListClusterNodeRequest) (response *ListClusterNodeResponse, err error) {
	response = CreateListClusterNodeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListClusterNodeWithChan(request *ListClusterNodeRequest) (<-chan *ListClusterNodeResponse, <-chan error) {
	responseChan := make(chan *ListClusterNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterNode(request)
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

func (client *Client) ListClusterNodeWithCallback(request *ListClusterNodeRequest, callback func(response *ListClusterNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterNodeResponse
		var err error
		defer close(result)
		response, err = client.ListClusterNode(request)
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

type ListClusterNodeRequest struct {
	*requests.RpcRequest
	PageSize        string `position:"Query" name:"PageSize"`
	PageNumber      string `position:"Query" name:"PageNumber"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListClusterNodeResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ClusterNodeList []struct {
		NodeId string `json:"NodeId" xml:"NodeId"`
		NodeIp string `json:"NodeIp" xml:"NodeIp"`
		Status string `json:"Status" xml:"Status"`
	} `json:"ClusterNodeList" xml:"ClusterNodeList"`
}

func CreateListClusterNodeRequest() (request *ListClusterNodeRequest) {
	request = &ListClusterNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterNode", "", "")
	return
}

func CreateListClusterNodeResponse() (response *ListClusterNodeResponse) {
	response = &ListClusterNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}