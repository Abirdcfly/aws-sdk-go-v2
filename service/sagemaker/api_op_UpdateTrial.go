// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTrialInput struct {
	_ struct{} `type:"structure"`

	// The name of the trial as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialName is displayed.
	DisplayName *string `min:"1" type:"string"`

	// The name of the trial to update.
	//
	// TrialName is a required field
	TrialName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTrialInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrialInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrialInput"}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

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

type UpdateTrialOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string `type:"string"`
}

// String returns the string representation
func (s UpdateTrialOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTrial = "UpdateTrial"

// UpdateTrialRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates the display name of a trial.
//
//    // Example sending a request using UpdateTrialRequest.
//    req := client.UpdateTrialRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateTrial
func (c *Client) UpdateTrialRequest(input *UpdateTrialInput) UpdateTrialRequest {
	op := &aws.Operation{
		Name:       opUpdateTrial,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTrialInput{}
	}

	req := c.newRequest(op, input, &UpdateTrialOutput{})
	return UpdateTrialRequest{Request: req, Input: input, Copy: c.UpdateTrialRequest}
}

// UpdateTrialRequest is the request type for the
// UpdateTrial API operation.
type UpdateTrialRequest struct {
	*aws.Request
	Input *UpdateTrialInput
	Copy  func(*UpdateTrialInput) UpdateTrialRequest
}

// Send marshals and sends the UpdateTrial API request.
func (r UpdateTrialRequest) Send(ctx context.Context) (*UpdateTrialResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrialResponse{
		UpdateTrialOutput: r.Request.Data.(*UpdateTrialOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrialResponse is the response type for the
// UpdateTrial API operation.
type UpdateTrialResponse struct {
	*UpdateTrialOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrial request.
func (r *UpdateTrialResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
