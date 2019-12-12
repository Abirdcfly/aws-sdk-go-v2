// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInsightRuleReportInput struct {
	_ struct{} `type:"structure"`

	// The end time of the data to use in the report. When used in a raw HTTP Query
	// API, it is formatted as yyyy-MM-dd'T'HH:mm:ss. For example, 2019-07-01T23:59:59.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// The maximum number of contributors to include in the report. The range is
	// 1 to 100. If you omit this, the default of 10 is used.
	MaxContributorCount *int64 `type:"integer"`

	// Specifies which metrics to use for aggregation of contributor values for
	// the report. You can specify one or more of the following metrics:
	//
	//    * UniqueContributors -- the number of unique contributors for each data
	//    point.
	//
	//    * MaxContributorValue -- the value of the top contributor for each data
	//    point. The identity of the contributor may change for each data point
	//    in the graph. If this rule aggregates by COUNT, the top contributor for
	//    each data point is the contributor with the most occurrences in that period.
	//    If the rule aggregates by SUM, the top contributor is the contributor
	//    with the highest sum in the log field specified by the rule's Value, during
	//    that period.
	//
	//    * SampleCount -- the number of data points matched by the rule.
	//
	//    * Sum -- the sum of the values from all contributors during the time period
	//    represented by that data point.
	//
	//    * Minimum -- the minimum value from a single observation during the time
	//    period represented by that data point.
	//
	//    * Maximum -- the maximum value from a single observation during the time
	//    period represented by that data point.
	//
	//    * Average -- the average value from all contributors during the time period
	//    represented by that data point.
	Metrics []string `type:"list"`

	// Determines what statistic to use to rank the contributors. Valid values are
	// SUM and MAXIMUM.
	OrderBy *string `min:"1" type:"string"`

	// The period, in seconds, to use for the statistics in the InsightRuleMetricDatapoint
	// results.
	//
	// Period is a required field
	Period *int64 `min:"1" type:"integer" required:"true"`

	// The name of the rule that you want to see data from.
	//
	// RuleName is a required field
	RuleName *string `min:"1" type:"string" required:"true"`

	// The start time of the data to use in the report. When used in a raw HTTP
	// Query API, it is formatted as yyyy-MM-dd'T'HH:mm:ss. For example, 2019-07-01T23:59:59.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s GetInsightRuleReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInsightRuleReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInsightRuleReportInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if s.OrderBy != nil && len(*s.OrderBy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrderBy", 1))
	}

	if s.Period == nil {
		invalidParams.Add(aws.NewErrParamRequired("Period"))
	}
	if s.Period != nil && *s.Period < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 1))
	}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInsightRuleReportOutput struct {
	_ struct{} `type:"structure"`

	// The sum of the values from all individual contributors that match the rule.
	AggregateValue *float64 `type:"double"`

	// Specifies whether this rule aggregates contributor data by COUNT or SUM.
	AggregationStatistic *string `type:"string"`

	// An approximate count of the unique contributors found by this rule in this
	// time period.
	ApproximateUniqueCount *int64 `type:"long"`

	// An array of the unique contributors found by this rule in this time period.
	// If the rule contains multiple keys, each combination of values for the keys
	// counts as a unique contributor.
	Contributors []InsightRuleContributor `type:"list"`

	// An array of the strings used as the keys for this rule. The keys are the
	// dimensions used to classify contributors. If the rule contains more than
	// one key, then each unique combination of values for the keys is counted as
	// a unique contributor.
	KeyLabels []string `type:"list"`

	// A time series of metric data points that matches the time period in the rule
	// request.
	MetricDatapoints []InsightRuleMetricDatapoint `type:"list"`
}

// String returns the string representation
func (s GetInsightRuleReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInsightRuleReport = "GetInsightRuleReport"

// GetInsightRuleReportRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// This operation returns the time series data collected by a Contributor Insights
// rule. The data includes the identity and number of contributors to the log
// group.
//
// You can also optionally return one or more statistics about each data point
// in the time series. These statistics can include the following:
//
//    * UniqueContributors -- the number of unique contributors for each data
//    point.
//
//    * MaxContributorValue -- the value of the top contributor for each data
//    point. The identity of the contributor may change for each data point
//    in the graph. If this rule aggregates by COUNT, the top contributor for
//    each data point is the contributor with the most occurrences in that period.
//    If the rule aggregates by SUM, the top contributor is the contributor
//    with the highest sum in the log field specified by the rule's Value, during
//    that period.
//
//    * SampleCount -- the number of data points matched by the rule.
//
//    * Sum -- the sum of the values from all contributors during the time period
//    represented by that data point.
//
//    * Minimum -- the minimum value from a single observation during the time
//    period represented by that data point.
//
//    * Maximum -- the maximum value from a single observation during the time
//    period represented by that data point.
//
//    * Average -- the average value from all contributors during the time period
//    represented by that data point.
//
//    // Example sending a request using GetInsightRuleReportRequest.
//    req := client.GetInsightRuleReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/GetInsightRuleReport
func (c *Client) GetInsightRuleReportRequest(input *GetInsightRuleReportInput) GetInsightRuleReportRequest {
	op := &aws.Operation{
		Name:       opGetInsightRuleReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInsightRuleReportInput{}
	}

	req := c.newRequest(op, input, &GetInsightRuleReportOutput{})
	return GetInsightRuleReportRequest{Request: req, Input: input, Copy: c.GetInsightRuleReportRequest}
}

// GetInsightRuleReportRequest is the request type for the
// GetInsightRuleReport API operation.
type GetInsightRuleReportRequest struct {
	*aws.Request
	Input *GetInsightRuleReportInput
	Copy  func(*GetInsightRuleReportInput) GetInsightRuleReportRequest
}

// Send marshals and sends the GetInsightRuleReport API request.
func (r GetInsightRuleReportRequest) Send(ctx context.Context) (*GetInsightRuleReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInsightRuleReportResponse{
		GetInsightRuleReportOutput: r.Request.Data.(*GetInsightRuleReportOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInsightRuleReportResponse is the response type for the
// GetInsightRuleReport API operation.
type GetInsightRuleReportResponse struct {
	*GetInsightRuleReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInsightRuleReport request.
func (r *GetInsightRuleReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
