// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTrialInput struct {
	_ struct{} `type:"structure"`

	// The name of the trial to describe.
	//
	// TrialName is a required field
	TrialName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTrialInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrialInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrialInput"}

	if s.TrialName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialName"))
	}
	if s.TrialName != nil && len(*s.TrialName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTrialOutput struct {
	_ struct{} `type:"structure"`

	// Who created the trial.
	CreatedBy *UserContext `type:"structure"`

	// When the trial was created.
	CreationTime *time.Time `type:"timestamp"`

	// The name of the trial as displayed. If DisplayName isn't specified, TrialName
	// is displayed.
	DisplayName *string `min:"1" type:"string"`

	// The name of the experiment the trial is part of.
	ExperimentName *string `min:"1" type:"string"`

	// Who last modified the trial.
	LastModifiedBy *UserContext `type:"structure"`

	// When the trial was last modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the source and, optionally, the job type.
	Source *TrialSource `type:"structure"`

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string `type:"string"`

	// The name of the trial.
	TrialName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTrialOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrial = "DescribeTrial"

// DescribeTrialRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Provides a list of a trial's properties.
//
//    // Example sending a request using DescribeTrialRequest.
//    req := client.DescribeTrialRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeTrial
func (c *Client) DescribeTrialRequest(input *DescribeTrialInput) DescribeTrialRequest {
	op := &aws.Operation{
		Name:       opDescribeTrial,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrialInput{}
	}

	req := c.newRequest(op, input, &DescribeTrialOutput{})
	return DescribeTrialRequest{Request: req, Input: input, Copy: c.DescribeTrialRequest}
}

// DescribeTrialRequest is the request type for the
// DescribeTrial API operation.
type DescribeTrialRequest struct {
	*aws.Request
	Input *DescribeTrialInput
	Copy  func(*DescribeTrialInput) DescribeTrialRequest
}

// Send marshals and sends the DescribeTrial API request.
func (r DescribeTrialRequest) Send(ctx context.Context) (*DescribeTrialResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrialResponse{
		DescribeTrialOutput: r.Request.Data.(*DescribeTrialOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrialResponse is the response type for the
// DescribeTrial API operation.
type DescribeTrialResponse struct {
	*DescribeTrialOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrial request.
func (r *DescribeTrialResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
