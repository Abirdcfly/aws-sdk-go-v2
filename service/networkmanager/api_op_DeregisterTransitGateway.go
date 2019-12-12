// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeregisterTransitGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the transit gateway.
	//
	// TransitGatewayArn is a required field
	TransitGatewayArn *string `location:"uri" locationName:"transitGatewayArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterTransitGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterTransitGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterTransitGatewayInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}

	if s.TransitGatewayArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeregisterTransitGatewayInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TransitGatewayArn != nil {
		v := *s.TransitGatewayArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "transitGatewayArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeregisterTransitGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The transit gateway registration information.
	TransitGatewayRegistration *TransitGatewayRegistration `type:"structure"`
}

// String returns the string representation
func (s DeregisterTransitGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeregisterTransitGatewayOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TransitGatewayRegistration != nil {
		v := s.TransitGatewayRegistration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TransitGatewayRegistration", v, metadata)
	}
	return nil
}

const opDeregisterTransitGateway = "DeregisterTransitGateway"

// DeregisterTransitGatewayRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Deregisters a transit gateway from your global network. This action does
// not delete your transit gateway, or modify any of its attachments. This action
// removes any customer gateway associations.
//
//    // Example sending a request using DeregisterTransitGatewayRequest.
//    req := client.DeregisterTransitGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/DeregisterTransitGateway
func (c *Client) DeregisterTransitGatewayRequest(input *DeregisterTransitGatewayInput) DeregisterTransitGatewayRequest {
	op := &aws.Operation{
		Name:       opDeregisterTransitGateway,
		HTTPMethod: "DELETE",
		HTTPPath:   "/global-networks/{globalNetworkId}/transit-gateway-registrations/{transitGatewayArn}",
	}

	if input == nil {
		input = &DeregisterTransitGatewayInput{}
	}

	req := c.newRequest(op, input, &DeregisterTransitGatewayOutput{})
	return DeregisterTransitGatewayRequest{Request: req, Input: input, Copy: c.DeregisterTransitGatewayRequest}
}

// DeregisterTransitGatewayRequest is the request type for the
// DeregisterTransitGateway API operation.
type DeregisterTransitGatewayRequest struct {
	*aws.Request
	Input *DeregisterTransitGatewayInput
	Copy  func(*DeregisterTransitGatewayInput) DeregisterTransitGatewayRequest
}

// Send marshals and sends the DeregisterTransitGateway API request.
func (r DeregisterTransitGatewayRequest) Send(ctx context.Context) (*DeregisterTransitGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTransitGatewayResponse{
		DeregisterTransitGatewayOutput: r.Request.Data.(*DeregisterTransitGatewayOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTransitGatewayResponse is the response type for the
// DeregisterTransitGateway API operation.
type DeregisterTransitGatewayResponse struct {
	*DeregisterTransitGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTransitGateway request.
func (r *DeregisterTransitGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
