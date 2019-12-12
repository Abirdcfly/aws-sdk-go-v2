// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketCorsInput struct {
	_ struct{} `type:"structure" payload:"CORSConfiguration"`

	// Specifies the bucket impacted by the corsconfiguration.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Describes the cross-origin access configuration for objects in an Amazon
	// S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev//cors.html) in the Amazon
	// Simple Storage Service Developer Guide.
	//
	// CORSConfiguration is a required field
	CORSConfiguration *CORSConfiguration `locationName:"CORSConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketCorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketCorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketCorsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.CORSConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("CORSConfiguration"))
	}
	if s.CORSConfiguration != nil {
		if err := s.CORSConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CORSConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketCorsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketCorsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CORSConfiguration != nil {
		v := s.CORSConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CORSConfiguration", v, metadata)
	}
	return nil
}

type PutBucketCorsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketCorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketCorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketCors = "PutBucketCors"

// PutBucketCorsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the cors configuration for your bucket. If the configuration exists,
// Amazon S3 replaces it.
//
// To use this operation, you must be allowed to perform the s3:PutBucketCORS
// action. By default, the bucket owner has this permission and can grant it
// to others.
//
// You set this configuration on a bucket so that the bucket can service cross-origin
// requests. For example, you might want to enable a request whose origin is
// http://www.example.com to access your Amazon S3 bucket at my.example.bucket.com
// by using the browser's XMLHttpRequest capability.
//
// To enable cross-origin resource sharing (CORS) on a bucket, you add the cors
// subresource to the bucket. The cors subresource is an XML document in which
// you configure rules that identify origins and the HTTP methods that can be
// executed on your bucket. The document is limited to 64 KB in size.
//
// When Amazon S3 receives a cross-origin request (or a pre-flight OPTIONS request)
// against a bucket, it evaluates the cors configuration on the bucket and uses
// the first CORSRule rule that matches the incoming browser request to enable
// a cross-origin request. For a rule to match, the following conditions must
// be met:
//
//    * The request's Origin header must match AllowedOrigin elements.
//
//    * The request method (for example, GET, PUT, HEAD, and so on) or the Access-Control-Request-Method
//    header in case of a pre-flight OPTIONS request must be one of the AllowedMethod
//    elements.
//
//    * Every header specified in the Access-Control-Request-Headers request
//    header of a pre-flight request must match an AllowedHeader element.
//
// For more information about CORS, go to Enabling Cross-Origin Resource Sharing
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon
// Simple Storage Service Developer Guide.
//
// Related Resources
//
//    * GetBucketCors
//
//    * DeleteBucketCors
//
//    * RESTOPTIONSobject
//
//    // Example sending a request using PutBucketCorsRequest.
//    req := client.PutBucketCorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketCors
func (c *Client) PutBucketCorsRequest(input *PutBucketCorsInput) PutBucketCorsRequest {
	op := &aws.Operation{
		Name:       opPutBucketCors,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &PutBucketCorsInput{}
	}

	req := c.newRequest(op, input, &PutBucketCorsOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketCorsRequest{Request: req, Input: input, Copy: c.PutBucketCorsRequest}
}

// PutBucketCorsRequest is the request type for the
// PutBucketCors API operation.
type PutBucketCorsRequest struct {
	*aws.Request
	Input *PutBucketCorsInput
	Copy  func(*PutBucketCorsInput) PutBucketCorsRequest
}

// Send marshals and sends the PutBucketCors API request.
func (r PutBucketCorsRequest) Send(ctx context.Context) (*PutBucketCorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketCorsResponse{
		PutBucketCorsOutput: r.Request.Data.(*PutBucketCorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketCorsResponse is the response type for the
// PutBucketCors API operation.
type PutBucketCorsResponse struct {
	*PutBucketCorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketCors request.
func (r *PutBucketCorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
