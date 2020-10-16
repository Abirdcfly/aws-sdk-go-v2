// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples adds headers to the input of a request and response by prefix.
func (c *Client) HttpPrefixHeaders(ctx context.Context, params *HttpPrefixHeadersInput, optFns ...func(*Options)) (*HttpPrefixHeadersOutput, error) {
	if params == nil {
		params = &HttpPrefixHeadersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPrefixHeaders", params, optFns, addOperationHttpPrefixHeadersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPrefixHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPrefixHeadersInput struct {
	Foo *string

	FooMap map[string]*string
}

type HttpPrefixHeadersOutput struct {
	Foo *string

	FooMap map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpPrefixHeadersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpPrefixHeaders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPrefixHeaders{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPrefixHeaders(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opHttpPrefixHeaders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPrefixHeaders",
	}
}