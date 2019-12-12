// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutOrganizationConformancePackInput struct {
	_ struct{} `type:"structure"`

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []ConformancePackInputParameter `type:"list"`

	// Location of an Amazon S3 bucket where AWS Config can deliver evaluation results.
	// AWS Config stores intermediate files while processing conformance pack template.
	//
	// The delivery bucket name should start with awsconfigconforms. For example:
	// "Resource": "arn:aws:s3:::your_bucket_name/*". For more information, see
	// Permissions for cross account bucket access (https://docs.aws.amazon.com/config/latest/developerguide/conformance-pack-organization-apis.html).
	//
	// DeliveryS3Bucket is a required field
	DeliveryS3Bucket *string `min:"3" type:"string" required:"true"`

	// The prefix for the Amazon S3 bucket.
	DeliveryS3KeyPrefix *string `min:"1" type:"string"`

	// A list of AWS accounts to be excluded from an organization conformance pack
	// while deploying a conformance pack.
	ExcludedAccounts []string `type:"list"`

	// Name of the organization conformance pack you want to create.
	//
	// OrganizationConformancePackName is a required field
	OrganizationConformancePackName *string `min:"1" type:"string" required:"true"`

	// A string containing full conformance pack template body. Structure containing
	// the template body with a minimum length of 1 byte and a maximum length of
	// 51,200 bytes.
	TemplateBody *string `min:"1" type:"string"`

	// Location of file containing the template body. The uri must point to the
	// conformance pack template (max size: 300 KB).
	//
	// You must have access to read Amazon S3 bucket.
	TemplateS3Uri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutOrganizationConformancePackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutOrganizationConformancePackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutOrganizationConformancePackInput"}

	if s.DeliveryS3Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryS3Bucket"))
	}
	if s.DeliveryS3Bucket != nil && len(*s.DeliveryS3Bucket) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryS3Bucket", 3))
	}
	if s.DeliveryS3KeyPrefix != nil && len(*s.DeliveryS3KeyPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryS3KeyPrefix", 1))
	}

	if s.OrganizationConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConformancePackName"))
	}
	if s.OrganizationConformancePackName != nil && len(*s.OrganizationConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConformancePackName", 1))
	}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateS3Uri != nil && len(*s.TemplateS3Uri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateS3Uri", 1))
	}
	if s.ConformancePackInputParameters != nil {
		for i, v := range s.ConformancePackInputParameters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ConformancePackInputParameters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutOrganizationConformancePackOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the organization conformance pack.
	OrganizationConformancePackArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutOrganizationConformancePackOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutOrganizationConformancePack = "PutOrganizationConformancePack"

// PutOrganizationConformancePackRequest returns a request value for making API operation for
// AWS Config.
//
// Deploys conformance packs across member accounts in an AWS Organization.
//
// This API enables organization service access for config-multiaccountsetup.amazonaws.com
// through the EnableAWSServiceAccess action and creates a service linked role
// AWSServiceRoleForConfigMultiAccountSetup in the master account of your organization.
// The service linked role is created only when the role does not exist in the
// master account. AWS Config verifies the existence of role with GetRole action.
//
// You must specify either the TemplateS3Uri or the TemplateBody parameter,
// but not both. If you provide both AWS Config uses the TemplateS3Uri parameter
// and ignores the TemplateBody parameter.
//
// AWS Config sets the state of a conformance pack to CREATE_IN_PROGRESS and
// UPDATE_IN_PROGRESS until the confomance pack is created or updated. You cannot
// update a conformance pack while it is in this state.
//
// You can create 6 conformance packs with 25 AWS Config rules in each pack.
//
//    // Example sending a request using PutOrganizationConformancePackRequest.
//    req := client.PutOrganizationConformancePackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutOrganizationConformancePack
func (c *Client) PutOrganizationConformancePackRequest(input *PutOrganizationConformancePackInput) PutOrganizationConformancePackRequest {
	op := &aws.Operation{
		Name:       opPutOrganizationConformancePack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutOrganizationConformancePackInput{}
	}

	req := c.newRequest(op, input, &PutOrganizationConformancePackOutput{})
	return PutOrganizationConformancePackRequest{Request: req, Input: input, Copy: c.PutOrganizationConformancePackRequest}
}

// PutOrganizationConformancePackRequest is the request type for the
// PutOrganizationConformancePack API operation.
type PutOrganizationConformancePackRequest struct {
	*aws.Request
	Input *PutOrganizationConformancePackInput
	Copy  func(*PutOrganizationConformancePackInput) PutOrganizationConformancePackRequest
}

// Send marshals and sends the PutOrganizationConformancePack API request.
func (r PutOrganizationConformancePackRequest) Send(ctx context.Context) (*PutOrganizationConformancePackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutOrganizationConformancePackResponse{
		PutOrganizationConformancePackOutput: r.Request.Data.(*PutOrganizationConformancePackOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutOrganizationConformancePackResponse is the response type for the
// PutOrganizationConformancePack API operation.
type PutOrganizationConformancePackResponse struct {
	*PutOrganizationConformancePackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutOrganizationConformancePack request.
func (r *PutOrganizationConformancePackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
