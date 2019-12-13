// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteAssetInput struct {
	_ struct{} `type:"structure"`

	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"AssetId" type:"string" required:"true"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// RevisionId is a required field
	RevisionId *string `location:"uri" locationName:"RevisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssetInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AssetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteAssetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAssetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAsset = "DeleteAsset"

// DeleteAssetRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation deletes an asset.
//
//    // Example sending a request using DeleteAssetRequest.
//    req := client.DeleteAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/DeleteAsset
func (c *Client) DeleteAssetRequest(input *DeleteAssetInput) DeleteAssetRequest {
	op := &aws.Operation{
		Name:       opDeleteAsset,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/data-sets/{DataSetId}/revisions/{RevisionId}/assets/{AssetId}",
	}

	if input == nil {
		input = &DeleteAssetInput{}
	}

	req := c.newRequest(op, input, &DeleteAssetOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAssetRequest{Request: req, Input: input, Copy: c.DeleteAssetRequest}
}

// DeleteAssetRequest is the request type for the
// DeleteAsset API operation.
type DeleteAssetRequest struct {
	*aws.Request
	Input *DeleteAssetInput
	Copy  func(*DeleteAssetInput) DeleteAssetRequest
}

// Send marshals and sends the DeleteAsset API request.
func (r DeleteAssetRequest) Send(ctx context.Context) (*DeleteAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssetResponse{
		DeleteAssetOutput: r.Request.Data.(*DeleteAssetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssetResponse is the response type for the
// DeleteAsset API operation.
type DeleteAssetResponse struct {
	*DeleteAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAsset request.
func (r *DeleteAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}