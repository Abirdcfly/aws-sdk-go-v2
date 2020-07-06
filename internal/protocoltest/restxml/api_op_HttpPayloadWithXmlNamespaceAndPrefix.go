// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML namespace.
func (c *Client) HttpPayloadWithXmlNamespaceAndPrefix(ctx context.Context, params *HttpPayloadWithXmlNamespaceAndPrefixInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNamespaceAndPrefixOutput, error) {
	stack := middleware.NewStack("HttpPayloadWithXmlNamespaceAndPrefix", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespaceAndPrefix(options.Region), middleware.Before)
	addawsRestxml_serdeOpHttpPayloadWithXmlNamespaceAndPrefixMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "HttpPayloadWithXmlNamespaceAndPrefix",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadWithXmlNamespaceAndPrefixOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNamespaceAndPrefixInput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix
}

type HttpPayloadWithXmlNamespaceAndPrefixOutput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpPayloadWithXmlNamespaceAndPrefixMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlNamespaceAndPrefix{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlNamespaceAndPrefix{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespaceAndPrefix(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "HttpPayloadWithXmlNamespaceAndPrefix",
	}
}
