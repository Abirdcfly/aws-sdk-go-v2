// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteConformancePackInput struct {
	_ struct{} `type:"structure"`

	// Name of the conformance pack you want to delete.
	//
	// ConformancePackName is a required field
	ConformancePackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConformancePackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConformancePackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConformancePackInput"}

	if s.ConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConformancePackName"))
	}
	if s.ConformancePackName != nil && len(*s.ConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConformancePackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteConformancePackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConformancePackOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConformancePack = "DeleteConformancePack"

// DeleteConformancePackRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified conformance pack and all the AWS Config rules, remediation
// actions, and all evaluation results within that conformance pack.
//
// AWS Config sets the conformance pack to DELETE_IN_PROGRESS until the deletion
// is complete. You cannot update a conformance pack while it is in this state.
//
//    // Example sending a request using DeleteConformancePackRequest.
//    req := client.DeleteConformancePackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConformancePack
func (c *Client) DeleteConformancePackRequest(input *DeleteConformancePackInput) DeleteConformancePackRequest {
	op := &aws.Operation{
		Name:       opDeleteConformancePack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConformancePackInput{}
	}

	req := c.newRequest(op, input, &DeleteConformancePackOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteConformancePackRequest{Request: req, Input: input, Copy: c.DeleteConformancePackRequest}
}

// DeleteConformancePackRequest is the request type for the
// DeleteConformancePack API operation.
type DeleteConformancePackRequest struct {
	*aws.Request
	Input *DeleteConformancePackInput
	Copy  func(*DeleteConformancePackInput) DeleteConformancePackRequest
}

// Send marshals and sends the DeleteConformancePack API request.
func (r DeleteConformancePackRequest) Send(ctx context.Context) (*DeleteConformancePackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConformancePackResponse{
		DeleteConformancePackOutput: r.Request.Data.(*DeleteConformancePackOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConformancePackResponse is the response type for the
// DeleteConformancePack API operation.
type DeleteConformancePackResponse struct {
	*DeleteConformancePackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConformancePack request.
func (r *DeleteConformancePackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
