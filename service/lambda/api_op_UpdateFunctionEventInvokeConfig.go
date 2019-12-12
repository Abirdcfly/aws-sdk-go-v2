// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFunctionEventInvokeConfigInput struct {
	_ struct{} `type:"structure"`

	// A destination for events after they have been sent to a function for processing.
	//
	// Destinations
	//
	//    * Function - The Amazon Resource Name (ARN) of a Lambda function.
	//
	//    * Queue - The ARN of an SQS queue.
	//
	//    * Topic - The ARN of an SNS topic.
	//
	//    * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *DestinationConfig `type:"structure"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int64 `min:"60" type:"integer"`

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int64 `type:"integer"`

	// A version number or alias name.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateFunctionEventInvokeConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFunctionEventInvokeConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFunctionEventInvokeConfigInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.MaximumEventAgeInSeconds != nil && *s.MaximumEventAgeInSeconds < 60 {
		invalidParams.Add(aws.NewErrParamMinValue("MaximumEventAgeInSeconds", 60))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionEventInvokeConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DestinationConfig != nil {
		v := s.DestinationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DestinationConfig", v, metadata)
	}
	if s.MaximumEventAgeInSeconds != nil {
		v := *s.MaximumEventAgeInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumEventAgeInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.MaximumRetryAttempts != nil {
		v := *s.MaximumRetryAttempts

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumRetryAttempts", protocol.Int64Value(v), metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateFunctionEventInvokeConfigOutput struct {
	_ struct{} `type:"structure"`

	// A destination for events after they have been sent to a function for processing.
	//
	// Destinations
	//
	//    * Function - The Amazon Resource Name (ARN) of a Lambda function.
	//
	//    * Queue - The ARN of an SQS queue.
	//
	//    * Topic - The ARN of an SNS topic.
	//
	//    * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *DestinationConfig `type:"structure"`

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string `type:"string"`

	// The date and time that the configuration was last updated.
	LastModified *time.Time `type:"timestamp"`

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int64 `min:"60" type:"integer"`

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int64 `type:"integer"`
}

// String returns the string representation
func (s UpdateFunctionEventInvokeConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionEventInvokeConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DestinationConfig != nil {
		v := s.DestinationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DestinationConfig", v, metadata)
	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.MaximumEventAgeInSeconds != nil {
		v := *s.MaximumEventAgeInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumEventAgeInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.MaximumRetryAttempts != nil {
		v := *s.MaximumRetryAttempts

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaximumRetryAttempts", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opUpdateFunctionEventInvokeConfig = "UpdateFunctionEventInvokeConfig"

// UpdateFunctionEventInvokeConfigRequest returns a request value for making API operation for
// AWS Lambda.
//
// Updates the configuration for asynchronous invocation for a function, version,
// or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
//
//    // Example sending a request using UpdateFunctionEventInvokeConfigRequest.
//    req := client.UpdateFunctionEventInvokeConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateFunctionEventInvokeConfig
func (c *Client) UpdateFunctionEventInvokeConfigRequest(input *UpdateFunctionEventInvokeConfigInput) UpdateFunctionEventInvokeConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateFunctionEventInvokeConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-09-25/functions/{FunctionName}/event-invoke-config",
	}

	if input == nil {
		input = &UpdateFunctionEventInvokeConfigInput{}
	}

	req := c.newRequest(op, input, &UpdateFunctionEventInvokeConfigOutput{})
	return UpdateFunctionEventInvokeConfigRequest{Request: req, Input: input, Copy: c.UpdateFunctionEventInvokeConfigRequest}
}

// UpdateFunctionEventInvokeConfigRequest is the request type for the
// UpdateFunctionEventInvokeConfig API operation.
type UpdateFunctionEventInvokeConfigRequest struct {
	*aws.Request
	Input *UpdateFunctionEventInvokeConfigInput
	Copy  func(*UpdateFunctionEventInvokeConfigInput) UpdateFunctionEventInvokeConfigRequest
}

// Send marshals and sends the UpdateFunctionEventInvokeConfig API request.
func (r UpdateFunctionEventInvokeConfigRequest) Send(ctx context.Context) (*UpdateFunctionEventInvokeConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFunctionEventInvokeConfigResponse{
		UpdateFunctionEventInvokeConfigOutput: r.Request.Data.(*UpdateFunctionEventInvokeConfigOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFunctionEventInvokeConfigResponse is the response type for the
// UpdateFunctionEventInvokeConfig API operation.
type UpdateFunctionEventInvokeConfigResponse struct {
	*UpdateFunctionEventInvokeConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFunctionEventInvokeConfig request.
func (r *UpdateFunctionEventInvokeConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
