// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This tests how timestamps are serialized, including using the default format of
// date-time and various @timestampFormat trait values.
func (c *Client) XmlTimestamps(ctx context.Context, params *XmlTimestampsInput, optFns ...func(*Options)) (*XmlTimestampsOutput, error) {
	if params == nil {
		params = &XmlTimestampsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlTimestamps", params, optFns, addOperationXmlTimestampsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlTimestampsInput struct {
}

type XmlTimestampsOutput struct {
	DateTime *time.Time

	EpochSeconds *time.Time

	HttpDate *time.Time

	Normal *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlTimestampsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpXmlTimestamps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpXmlTimestamps{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlTimestamps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opXmlTimestamps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlTimestamps",
	}
}