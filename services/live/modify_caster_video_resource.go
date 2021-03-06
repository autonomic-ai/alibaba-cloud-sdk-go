package live

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

// ModifyCasterVideoResource invokes the live.ModifyCasterVideoResource API synchronously
// api document: https://help.aliyun.com/api/live/modifycastervideoresource.html
func (client *Client) ModifyCasterVideoResource(request *ModifyCasterVideoResourceRequest) (response *ModifyCasterVideoResourceResponse, err error) {
	response = CreateModifyCasterVideoResourceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCasterVideoResourceWithChan invokes the live.ModifyCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/modifycastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCasterVideoResourceWithChan(request *ModifyCasterVideoResourceRequest) (<-chan *ModifyCasterVideoResourceResponse, <-chan error) {
	responseChan := make(chan *ModifyCasterVideoResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCasterVideoResource(request)
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

// ModifyCasterVideoResourceWithCallback invokes the live.ModifyCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/modifycastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCasterVideoResourceWithCallback(request *ModifyCasterVideoResourceRequest, callback func(response *ModifyCasterVideoResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCasterVideoResourceResponse
		var err error
		defer close(result)
		response, err = client.ModifyCasterVideoResource(request)
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

// ModifyCasterVideoResourceRequest is the request struct for api ModifyCasterVideoResource
type ModifyCasterVideoResourceRequest struct {
	*requests.RpcRequest
	ResourceId    string           `position:"Query" name:"ResourceId"`
	BeginOffset   requests.Integer `position:"Query" name:"BeginOffset"`
	VodUrl        string           `position:"Query" name:"VodUrl"`
	LiveStreamUrl string           `position:"Query" name:"LiveStreamUrl"`
	CasterId      string           `position:"Query" name:"CasterId"`
	EndOffset     requests.Integer `position:"Query" name:"EndOffset"`
	ResourceName  string           `position:"Query" name:"ResourceName"`
	RepeatNum     requests.Integer `position:"Query" name:"RepeatNum"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	MaterialId    string           `position:"Query" name:"MaterialId"`
}

// ModifyCasterVideoResourceResponse is the response struct for api ModifyCasterVideoResource
type ModifyCasterVideoResourceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	CasterId   string `json:"CasterId" xml:"CasterId"`
	ResourceId string `json:"ResourceId" xml:"ResourceId"`
}

// CreateModifyCasterVideoResourceRequest creates a request to invoke ModifyCasterVideoResource API
func CreateModifyCasterVideoResourceRequest() (request *ModifyCasterVideoResourceRequest) {
	request = &ModifyCasterVideoResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyCasterVideoResource", "live", "openAPI")
	return
}

// CreateModifyCasterVideoResourceResponse creates a response to parse from ModifyCasterVideoResource response
func CreateModifyCasterVideoResourceResponse() (response *ModifyCasterVideoResourceResponse) {
	response = &ModifyCasterVideoResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
