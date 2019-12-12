// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteBucketInput struct {
	_ struct{} `type:"structure"`

	// Specifies the bucket being deleted.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteBucketOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucket = "DeleteBucket"

// DeleteBucketRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the bucket. All objects (including all object versions and delete
// markers) in the bucket must be deleted before the bucket itself can be deleted.
//
// Related Resources
//
//    *
//
//    *
//
//    // Example sending a request using DeleteBucketRequest.
//    req := client.DeleteBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucket
func (c *Client) DeleteBucketRequest(input *DeleteBucketInput) DeleteBucketRequest {
	op := &aws.Operation{
		Name:       opDeleteBucket,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &DeleteBucketInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketRequest{Request: req, Input: input, Copy: c.DeleteBucketRequest}
}

// DeleteBucketRequest is the request type for the
// DeleteBucket API operation.
type DeleteBucketRequest struct {
	*aws.Request
	Input *DeleteBucketInput
	Copy  func(*DeleteBucketInput) DeleteBucketRequest
}

// Send marshals and sends the DeleteBucket API request.
func (r DeleteBucketRequest) Send(ctx context.Context) (*DeleteBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketResponse{
		DeleteBucketOutput: r.Request.Data.(*DeleteBucketOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketResponse is the response type for the
// DeleteBucket API operation.
type DeleteBucketResponse struct {
	*DeleteBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucket request.
func (r *DeleteBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
