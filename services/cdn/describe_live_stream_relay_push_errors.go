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

// DescribeLiveStreamRelayPushErrors invokes the cdn.DescribeLiveStreamRelayPushErrors API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamrelaypusherrors.html
func (client *Client) DescribeLiveStreamRelayPushErrors(request *DescribeLiveStreamRelayPushErrorsRequest) (response *DescribeLiveStreamRelayPushErrorsResponse, err error) {
	response = CreateDescribeLiveStreamRelayPushErrorsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamRelayPushErrorsWithChan invokes the cdn.DescribeLiveStreamRelayPushErrors API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamrelaypusherrors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRelayPushErrorsWithChan(request *DescribeLiveStreamRelayPushErrorsRequest) (<-chan *DescribeLiveStreamRelayPushErrorsResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRelayPushErrorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRelayPushErrors(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeLiveStreamRelayPushErrorsWithCallback invokes the cdn.DescribeLiveStreamRelayPushErrors API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamrelaypusherrors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRelayPushErrorsWithCallback(request *DescribeLiveStreamRelayPushErrorsRequest, callback func(response *DescribeLiveStreamRelayPushErrorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRelayPushErrorsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRelayPushErrors(request)
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

// DescribeLiveStreamRelayPushErrorsRequest is the request struct for api DescribeLiveStreamRelayPushErrors
type DescribeLiveStreamRelayPushErrorsRequest struct {
	*requests.RpcRequest
	RelayDomain   string           `position:"Query" name:"RelayDomain"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamRelayPushErrorsResponse is the response struct for api DescribeLiveStreamRelayPushErrors
type DescribeLiveStreamRelayPushErrorsResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	RelayPushErrorsModelList RelayPushErrorsModelList `json:"RelayPushErrorsModelList" xml:"RelayPushErrorsModelList"`
}

// CreateDescribeLiveStreamRelayPushErrorsRequest creates a request to invoke DescribeLiveStreamRelayPushErrors API
func CreateDescribeLiveStreamRelayPushErrorsRequest() (request *DescribeLiveStreamRelayPushErrorsRequest) {
	request = &DescribeLiveStreamRelayPushErrorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushErrors", "", "")
	return
}

// CreateDescribeLiveStreamRelayPushErrorsResponse creates a response to parse from DescribeLiveStreamRelayPushErrors response
func CreateDescribeLiveStreamRelayPushErrorsResponse() (response *DescribeLiveStreamRelayPushErrorsResponse) {
	response = &DescribeLiveStreamRelayPushErrorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
