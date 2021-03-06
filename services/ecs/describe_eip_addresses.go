package ecs

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

func (client *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (response *DescribeEipAddressesResponse, err error) {
	response = CreateDescribeEipAddressesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeEipAddressesWithChan(request *DescribeEipAddressesRequest) (<-chan *DescribeEipAddressesResponse, <-chan error) {
	responseChan := make(chan *DescribeEipAddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEipAddresses(request)
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

func (client *Client) DescribeEipAddressesWithCallback(request *DescribeEipAddressesRequest, callback func(response *DescribeEipAddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEipAddressesResponse
		var err error
		defer close(result)
		response, err = client.DescribeEipAddresses(request)
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

type DescribeEipAddressesRequest struct {
	*requests.RpcRequest
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	EipAddress             string           `position:"Query" name:"EipAddress"`
	Status                 string           `position:"Query" name:"Status"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
	Filter2Key             string           `position:"Query" name:"Filter.2.Key"`
	AssociatedInstanceType string           `position:"Query" name:"AssociatedInstanceType"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	AllocationId           string           `position:"Query" name:"AllocationId"`
	LockReason             string           `position:"Query" name:"LockReason"`
	Filter2Value           string           `position:"Query" name:"Filter.2.Value"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	AssociatedInstanceId   string           `position:"Query" name:"AssociatedInstanceId"`
	Filter1Value           string           `position:"Query" name:"Filter.1.Value"`
	Filter1Key             string           `position:"Query" name:"Filter.1.Key"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ChargeType             string           `position:"Query" name:"ChargeType"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
}

type DescribeEipAddressesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	EipAddresses struct {
		EipAddress []struct {
			RegionId           string `json:"RegionId" xml:"RegionId"`
			IpAddress          string `json:"IpAddress" xml:"IpAddress"`
			AllocationId       string `json:"AllocationId" xml:"AllocationId"`
			Status             string `json:"Status" xml:"Status"`
			InstanceId         string `json:"InstanceId" xml:"InstanceId"`
			Bandwidth          string `json:"Bandwidth" xml:"Bandwidth"`
			EipBandwidth       string `json:"EipBandwidth" xml:"EipBandwidth"`
			InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
			AllocationTime     string `json:"AllocationTime" xml:"AllocationTime"`
			InstanceType       string `json:"InstanceType" xml:"InstanceType"`
			ChargeType         string `json:"ChargeType" xml:"ChargeType"`
			ExpiredTime        string `json:"ExpiredTime" xml:"ExpiredTime"`
			OperationLocks     struct {
				LockReason []struct {
					LockReason string `json:"LockReason" xml:"LockReason"`
				} `json:"LockReason" xml:"LockReason"`
			} `json:"OperationLocks" xml:"OperationLocks"`
		} `json:"EipAddress" xml:"EipAddress"`
	} `json:"EipAddresses" xml:"EipAddresses"`
}

func CreateDescribeEipAddressesRequest() (request *DescribeEipAddressesRequest) {
	request = &DescribeEipAddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEipAddresses", "ecs", "openAPI")
	return
}

func CreateDescribeEipAddressesResponse() (response *DescribeEipAddressesResponse) {
	response = &DescribeEipAddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
