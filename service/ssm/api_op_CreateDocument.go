// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDocumentInput struct {
	_ struct{} `type:"structure"`

	// A list of key and value pairs that describe attachments to a version of a
	// document.
	Attachments []AttachmentsSource `type:"list"`

	// A valid JSON or YAML string.
	//
	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	// Specify the document format for the request. The document format can be either
	// JSON or YAML. JSON is the default format.
	DocumentFormat DocumentFormat `type:"string" enum:"true"`

	// The type of document to create. Valid document types include: Command, Policy,
	// Automation, Session, and Package.
	DocumentType DocumentType `type:"string" enum:"true"`

	// A name for the Systems Manager document.
	//
	// Do not use the following to begin the names of documents you create. They
	// are reserved by AWS for use as document prefixes:
	//
	//    * aws
	//
	//    * amazon
	//
	//    * amzn
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// A list of SSM documents required by a document. For example, an ApplicationConfiguration
	// document requires an ApplicationConfigurationSchema document.
	Requires []DocumentRequires `min:"1" type:"list"`

	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	// For example, you might want to tag an SSM document to identify the types
	// of targets or the environment where it will run. In this case, you could
	// specify the following key name/value pairs:
	//
	//    * Key=OS,Value=Windows
	//
	//    * Key=Environment,Value=Production
	//
	// To add tags to an existing SSM document, use the AddTagsToResource action.
	Tags []Tag `type:"list"`

	// Specify a target type to define the kinds of resources the document can run
	// on. For example, to run a document on EC2 instances, specify the following
	// value: /AWS::EC2::Instance. If you specify a value of '/' the document can
	// run on all types of resources. If you don't specify a value, the document
	// can't run on any resources. For a list of valid resource types, see AWS Resource
	// Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// in the AWS CloudFormation User Guide.
	TargetType *string `type:"string"`

	// An optional field specifying the version of the artifact you are creating
	// with the document. For example, "Release 12, Update 6". This value is unique
	// across all versions of a document, and cannot be changed.
	VersionName *string `type:"string"`
}

// String returns the string representation
func (s CreateDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDocumentInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if s.Content != nil && len(*s.Content) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Content", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Requires != nil && len(s.Requires) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Requires", 1))
	}
	if s.Attachments != nil {
		for i, v := range s.Attachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Requires != nil {
		for i, v := range s.Requires {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Requires", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDocumentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Systems Manager document.
	DocumentDescription *DocumentDescription `type:"structure"`
}

// String returns the string representation
func (s CreateDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDocument = "CreateDocument"

// CreateDocumentRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Creates a Systems Manager document.
//
// After you create a document, you can use CreateAssociation to associate it
// with one or more running instances.
//
//    // Example sending a request using CreateDocumentRequest.
//    req := client.CreateDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateDocument
func (c *Client) CreateDocumentRequest(input *CreateDocumentInput) CreateDocumentRequest {
	op := &aws.Operation{
		Name:       opCreateDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDocumentInput{}
	}

	req := c.newRequest(op, input, &CreateDocumentOutput{})
	return CreateDocumentRequest{Request: req, Input: input, Copy: c.CreateDocumentRequest}
}

// CreateDocumentRequest is the request type for the
// CreateDocument API operation.
type CreateDocumentRequest struct {
	*aws.Request
	Input *CreateDocumentInput
	Copy  func(*CreateDocumentInput) CreateDocumentRequest
}

// Send marshals and sends the CreateDocument API request.
func (r CreateDocumentRequest) Send(ctx context.Context) (*CreateDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDocumentResponse{
		CreateDocumentOutput: r.Request.Data.(*CreateDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDocumentResponse is the response type for the
// CreateDocument API operation.
type CreateDocumentResponse struct {
	*CreateDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDocument request.
func (r *CreateDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
