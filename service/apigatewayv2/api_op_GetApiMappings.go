// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetApiMappingsInput struct {
	_ struct{} `type:"structure"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetApiMappingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApiMappingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApiMappingsInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiMappingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApiMappingsOutput struct {
	_ struct{} `type:"structure"`

	Items []ApiMapping `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetApiMappingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiMappingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetApiMappings = "GetApiMappings"

// GetApiMappingsRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets API mappings.
//
//    // Example sending a request using GetApiMappingsRequest.
//    req := client.GetApiMappingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetApiMappings
func (c *Client) GetApiMappingsRequest(input *GetApiMappingsInput) GetApiMappingsRequest {
	op := &aws.Operation{
		Name:       opGetApiMappings,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/domainnames/{domainName}/apimappings",
	}

	if input == nil {
		input = &GetApiMappingsInput{}
	}

	req := c.newRequest(op, input, &GetApiMappingsOutput{})
	return GetApiMappingsRequest{Request: req, Input: input, Copy: c.GetApiMappingsRequest}
}

// GetApiMappingsRequest is the request type for the
// GetApiMappings API operation.
type GetApiMappingsRequest struct {
	*aws.Request
	Input *GetApiMappingsInput
	Copy  func(*GetApiMappingsInput) GetApiMappingsRequest
}

// Send marshals and sends the GetApiMappings API request.
func (r GetApiMappingsRequest) Send(ctx context.Context) (*GetApiMappingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApiMappingsResponse{
		GetApiMappingsOutput: r.Request.Data.(*GetApiMappingsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApiMappingsResponse is the response type for the
// GetApiMappings API operation.
type GetApiMappingsResponse struct {
	*GetApiMappingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApiMappings request.
func (r *GetApiMappingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
