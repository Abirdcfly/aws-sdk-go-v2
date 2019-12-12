// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEndpointsInput struct {
	_ struct{} `type:"structure"`

	// Filters the endpoints that are returned. You can filter endpoints on their
	// name, model, status, or the date and time that they were created. You can
	// only set one filter at a time.
	Filter *EndpointFilter `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// Displays a list of endpoint properties being retrieved by the service in
	// response to the request.
	EndpointPropertiesList []EndpointProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEndpoints = "ListEndpoints"

// ListEndpointsRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of all existing endpoints that you've created.
//
//    // Example sending a request using ListEndpointsRequest.
//    req := client.ListEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListEndpoints
func (c *Client) ListEndpointsRequest(input *ListEndpointsInput) ListEndpointsRequest {
	op := &aws.Operation{
		Name:       opListEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEndpointsInput{}
	}

	req := c.newRequest(op, input, &ListEndpointsOutput{})
	return ListEndpointsRequest{Request: req, Input: input, Copy: c.ListEndpointsRequest}
}

// ListEndpointsRequest is the request type for the
// ListEndpoints API operation.
type ListEndpointsRequest struct {
	*aws.Request
	Input *ListEndpointsInput
	Copy  func(*ListEndpointsInput) ListEndpointsRequest
}

// Send marshals and sends the ListEndpoints API request.
func (r ListEndpointsRequest) Send(ctx context.Context) (*ListEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEndpointsResponse{
		ListEndpointsOutput: r.Request.Data.(*ListEndpointsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListEndpointsResponse is the response type for the
// ListEndpoints API operation.
type ListEndpointsResponse struct {
	*ListEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEndpoints request.
func (r *ListEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
