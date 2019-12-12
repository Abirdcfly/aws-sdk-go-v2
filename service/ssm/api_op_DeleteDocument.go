// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDocumentInput struct {
	_ struct{} `type:"structure"`

	// The version of the document that you want to delete. If not provided, all
	// versions of the document are deleted.
	DocumentVersion *string `type:"string"`

	// Some SSM document types require that you specify a Force flag before you
	// can delete the document. For example, you must specify a Force flag to delete
	// a document of type ApplicationConfigurationSchema. You can restrict access
	// to the Force flag in an AWS Identity and Access Management (IAM) policy.
	Force *bool `type:"boolean"`

	// The name of the document.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The version name of the document that you want to delete. If not provided,
	// all versions of the document are deleted.
	VersionName *string `type:"string"`
}

// String returns the string representation
func (s DeleteDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDocumentInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDocumentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDocument = "DeleteDocument"

// DeleteDocumentRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Deletes the Systems Manager document and all instance associations to the
// document.
//
// Before you delete the document, we recommend that you use DeleteAssociation
// to disassociate all instances that are associated with the document.
//
//    // Example sending a request using DeleteDocumentRequest.
//    req := client.DeleteDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteDocument
func (c *Client) DeleteDocumentRequest(input *DeleteDocumentInput) DeleteDocumentRequest {
	op := &aws.Operation{
		Name:       opDeleteDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDocumentInput{}
	}

	req := c.newRequest(op, input, &DeleteDocumentOutput{})
	return DeleteDocumentRequest{Request: req, Input: input, Copy: c.DeleteDocumentRequest}
}

// DeleteDocumentRequest is the request type for the
// DeleteDocument API operation.
type DeleteDocumentRequest struct {
	*aws.Request
	Input *DeleteDocumentInput
	Copy  func(*DeleteDocumentInput) DeleteDocumentRequest
}

// Send marshals and sends the DeleteDocument API request.
func (r DeleteDocumentRequest) Send(ctx context.Context) (*DeleteDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDocumentResponse{
		DeleteDocumentOutput: r.Request.Data.(*DeleteDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDocumentResponse is the response type for the
// DeleteDocument API operation.
type DeleteDocumentResponse struct {
	*DeleteDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDocument request.
func (r *DeleteDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
