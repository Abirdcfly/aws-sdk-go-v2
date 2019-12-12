// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreatePackagingGroupInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePackagingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePackagingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePackagingGroupInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePackagingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreatePackagingGroupOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	DomainName *string `locationName:"domainName" type:"string"`

	Id *string `locationName:"id" type:"string"`
}

// String returns the string representation
func (s CreatePackagingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePackagingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreatePackagingGroup = "CreatePackagingGroup"

// CreatePackagingGroupRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Creates a new MediaPackage VOD PackagingGroup resource.
//
//    // Example sending a request using CreatePackagingGroupRequest.
//    req := client.CreatePackagingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/CreatePackagingGroup
func (c *Client) CreatePackagingGroupRequest(input *CreatePackagingGroupInput) CreatePackagingGroupRequest {
	op := &aws.Operation{
		Name:       opCreatePackagingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/packaging_groups",
	}

	if input == nil {
		input = &CreatePackagingGroupInput{}
	}

	req := c.newRequest(op, input, &CreatePackagingGroupOutput{})
	return CreatePackagingGroupRequest{Request: req, Input: input, Copy: c.CreatePackagingGroupRequest}
}

// CreatePackagingGroupRequest is the request type for the
// CreatePackagingGroup API operation.
type CreatePackagingGroupRequest struct {
	*aws.Request
	Input *CreatePackagingGroupInput
	Copy  func(*CreatePackagingGroupInput) CreatePackagingGroupRequest
}

// Send marshals and sends the CreatePackagingGroup API request.
func (r CreatePackagingGroupRequest) Send(ctx context.Context) (*CreatePackagingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePackagingGroupResponse{
		CreatePackagingGroupOutput: r.Request.Data.(*CreatePackagingGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePackagingGroupResponse is the response type for the
// CreatePackagingGroup API operation.
type CreatePackagingGroupResponse struct {
	*CreatePackagingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePackagingGroup request.
func (r *CreatePackagingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
