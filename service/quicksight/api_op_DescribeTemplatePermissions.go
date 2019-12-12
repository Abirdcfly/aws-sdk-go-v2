// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeTemplatePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the template that you're describing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTemplatePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTemplatePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTemplatePermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeTemplatePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeTemplatePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of resource permissions to be set on the template.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon Resource Name (ARN) of the template.
	TemplateArn *string `type:"string"`

	// The ID for the template.
	TemplateId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTemplatePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeTemplatePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateArn != nil {
		v := *s.TemplateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeTemplatePermissions = "DescribeTemplatePermissions"

// DescribeTemplatePermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes read and write permissions on a template.
//
//    // Example sending a request using DescribeTemplatePermissionsRequest.
//    req := client.DescribeTemplatePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeTemplatePermissions
func (c *Client) DescribeTemplatePermissionsRequest(input *DescribeTemplatePermissionsInput) DescribeTemplatePermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeTemplatePermissions,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/permissions",
	}

	if input == nil {
		input = &DescribeTemplatePermissionsInput{}
	}

	req := c.newRequest(op, input, &DescribeTemplatePermissionsOutput{})
	return DescribeTemplatePermissionsRequest{Request: req, Input: input, Copy: c.DescribeTemplatePermissionsRequest}
}

// DescribeTemplatePermissionsRequest is the request type for the
// DescribeTemplatePermissions API operation.
type DescribeTemplatePermissionsRequest struct {
	*aws.Request
	Input *DescribeTemplatePermissionsInput
	Copy  func(*DescribeTemplatePermissionsInput) DescribeTemplatePermissionsRequest
}

// Send marshals and sends the DescribeTemplatePermissions API request.
func (r DescribeTemplatePermissionsRequest) Send(ctx context.Context) (*DescribeTemplatePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTemplatePermissionsResponse{
		DescribeTemplatePermissionsOutput: r.Request.Data.(*DescribeTemplatePermissionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTemplatePermissionsResponse is the response type for the
// DescribeTemplatePermissions API operation.
type DescribeTemplatePermissionsResponse struct {
	*DescribeTemplatePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTemplatePermissions request.
func (r *DescribeTemplatePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
