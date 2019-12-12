// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateCustomerGatewayInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the customer gateway. For more information,
	// see Resources Defined by Amazon EC2 (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	//
	// CustomerGatewayArn is a required field
	CustomerGatewayArn *string `location:"uri" locationName:"customerGatewayArn" type:"string" required:"true"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateCustomerGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateCustomerGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateCustomerGatewayInput"}

	if s.CustomerGatewayArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerGatewayArn"))
	}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateCustomerGatewayInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CustomerGatewayArn != nil {
		v := *s.CustomerGatewayArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "customerGatewayArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateCustomerGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the customer gateway association.
	CustomerGatewayAssociation *CustomerGatewayAssociation `type:"structure"`
}

// String returns the string representation
func (s DisassociateCustomerGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateCustomerGatewayOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CustomerGatewayAssociation != nil {
		v := s.CustomerGatewayAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CustomerGatewayAssociation", v, metadata)
	}
	return nil
}

const opDisassociateCustomerGateway = "DisassociateCustomerGateway"

// DisassociateCustomerGatewayRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Disassociates a customer gateway from a device and a link.
//
//    // Example sending a request using DisassociateCustomerGatewayRequest.
//    req := client.DisassociateCustomerGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/DisassociateCustomerGateway
func (c *Client) DisassociateCustomerGatewayRequest(input *DisassociateCustomerGatewayInput) DisassociateCustomerGatewayRequest {
	op := &aws.Operation{
		Name:       opDisassociateCustomerGateway,
		HTTPMethod: "DELETE",
		HTTPPath:   "/global-networks/{globalNetworkId}/customer-gateway-associations/{customerGatewayArn}",
	}

	if input == nil {
		input = &DisassociateCustomerGatewayInput{}
	}

	req := c.newRequest(op, input, &DisassociateCustomerGatewayOutput{})
	return DisassociateCustomerGatewayRequest{Request: req, Input: input, Copy: c.DisassociateCustomerGatewayRequest}
}

// DisassociateCustomerGatewayRequest is the request type for the
// DisassociateCustomerGateway API operation.
type DisassociateCustomerGatewayRequest struct {
	*aws.Request
	Input *DisassociateCustomerGatewayInput
	Copy  func(*DisassociateCustomerGatewayInput) DisassociateCustomerGatewayRequest
}

// Send marshals and sends the DisassociateCustomerGateway API request.
func (r DisassociateCustomerGatewayRequest) Send(ctx context.Context) (*DisassociateCustomerGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateCustomerGatewayResponse{
		DisassociateCustomerGatewayOutput: r.Request.Data.(*DisassociateCustomerGatewayOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateCustomerGatewayResponse is the response type for the
// DisassociateCustomerGateway API operation.
type DisassociateCustomerGatewayResponse struct {
	*DisassociateCustomerGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateCustomerGateway request.
func (r *DisassociateCustomerGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
