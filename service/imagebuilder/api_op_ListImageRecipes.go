// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListImageRecipesInput struct {
	_ struct{} `type:"structure"`

	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum items to return in a request.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token to specify where to start paginating. This is the NextToken from
	// a previously truncated response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The owner defines whose image recipes you wish to list. By default this request
	// will only show image recipes owned by your account. You may use this field
	// to specify if you wish to view image recipes owned by yourself, Amazon, or
	// those image recipes that have been shared with you by other customers.
	Owner Ownership `locationName:"owner" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListImageRecipesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListImageRecipesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListImageRecipesInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListImageRecipesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Owner) > 0 {
		v := s.Owner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "owner", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListImageRecipesOutput struct {
	_ struct{} `type:"structure"`

	// The list of image pipelines.
	ImageRecipeSummaryList []ImageRecipeSummary `locationName:"imageRecipeSummaryList" type:"list"`

	// The next token used for paginated responses. When this is not empty then
	// there are additional elements that the service that not include in this request.
	// Use this token with the next request to retrieve additional object.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s ListImageRecipesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListImageRecipesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImageRecipeSummaryList != nil {
		v := s.ImageRecipeSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "imageRecipeSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListImageRecipes = "ListImageRecipes"

// ListImageRecipesRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Returns a list of image recipes.
//
//    // Example sending a request using ListImageRecipesRequest.
//    req := client.ListImageRecipesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/ListImageRecipes
func (c *Client) ListImageRecipesRequest(input *ListImageRecipesInput) ListImageRecipesRequest {
	op := &aws.Operation{
		Name:       opListImageRecipes,
		HTTPMethod: "POST",
		HTTPPath:   "/ListImageRecipes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListImageRecipesInput{}
	}

	req := c.newRequest(op, input, &ListImageRecipesOutput{})
	return ListImageRecipesRequest{Request: req, Input: input, Copy: c.ListImageRecipesRequest}
}

// ListImageRecipesRequest is the request type for the
// ListImageRecipes API operation.
type ListImageRecipesRequest struct {
	*aws.Request
	Input *ListImageRecipesInput
	Copy  func(*ListImageRecipesInput) ListImageRecipesRequest
}

// Send marshals and sends the ListImageRecipes API request.
func (r ListImageRecipesRequest) Send(ctx context.Context) (*ListImageRecipesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListImageRecipesResponse{
		ListImageRecipesOutput: r.Request.Data.(*ListImageRecipesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListImageRecipesRequestPaginator returns a paginator for ListImageRecipes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListImageRecipesRequest(input)
//   p := imagebuilder.NewListImageRecipesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListImageRecipesPaginator(req ListImageRecipesRequest) ListImageRecipesPaginator {
	return ListImageRecipesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListImageRecipesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListImageRecipesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListImageRecipesPaginator struct {
	aws.Pager
}

func (p *ListImageRecipesPaginator) CurrentPage() *ListImageRecipesOutput {
	return p.Pager.CurrentPage().(*ListImageRecipesOutput)
}

// ListImageRecipesResponse is the response type for the
// ListImageRecipes API operation.
type ListImageRecipesResponse struct {
	*ListImageRecipesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListImageRecipes request.
func (r *ListImageRecipesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
