// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains details about an activity that failed during an execution.
type ActivityFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ActivityFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an activity.
type ActivityListItem struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the activity.
	//
	// ActivityArn is a required field
	ActivityArn *string `locationName:"activityArn" min:"1" type:"string" required:"true"`

	// The date the activity is created.
	//
	// CreationDate is a required field
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" required:"true"`

	// The name of the activity.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ActivityListItem) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an activity schedule failure that occurred during
// an execution.
type ActivityScheduleFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ActivityScheduleFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an activity scheduled during an execution.
type ActivityScheduledEventDetails struct {
	_ struct{} `type:"structure"`

	// The maximum allowed duration between two heartbeats for the activity task.
	HeartbeatInSeconds *int64 `locationName:"heartbeatInSeconds" type:"long"`

	// The JSON data input to the activity task.
	Input *string `locationName:"input" type:"string" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the scheduled activity.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The maximum allowed duration of the activity task.
	TimeoutInSeconds *int64 `locationName:"timeoutInSeconds" type:"long"`
}

// String returns the string representation
func (s ActivityScheduledEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the start of an activity during an execution.
type ActivityStartedEventDetails struct {
	_ struct{} `type:"structure"`

	// The name of the worker that the task is assigned to. These names are provided
	// by the workers when calling GetActivityTask.
	WorkerName *string `locationName:"workerName" type:"string"`
}

// String returns the string representation
func (s ActivityStartedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an activity that successfully terminated during an
// execution.
type ActivitySucceededEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data output by the activity task.
	Output *string `locationName:"output" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ActivitySucceededEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an activity timeout that occurred during an execution.
type ActivityTimedOutEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the timeout.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ActivityTimedOutEventDetails) String() string {
	return awsutil.Prettify(s)
}

type CloudWatchLogsLogGroup struct {
	_ struct{} `type:"structure"`

	// The ARN of the the CloudWatch log group to which you want your logs emitted
	// to. The ARN must end with :*
	LogGroupArn *string `locationName:"logGroupArn" min:"1" type:"string"`
}

// String returns the string representation
func (s CloudWatchLogsLogGroup) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CloudWatchLogsLogGroup) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CloudWatchLogsLogGroup"}
	if s.LogGroupArn != nil && len(*s.LogGroupArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about an abort of an execution.
type ExecutionAbortedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ExecutionAbortedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an execution failure event.
type ExecutionFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ExecutionFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an execution.
type ExecutionListItem struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// ExecutionArn is a required field
	ExecutionArn *string `locationName:"executionArn" min:"1" type:"string" required:"true"`

	// The name of the execution.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The date the execution started.
	//
	// StartDate is a required field
	StartDate *time.Time `locationName:"startDate" type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the executed state machine.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`

	// The current status of the execution.
	//
	// Status is a required field
	Status ExecutionStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// If the execution already ended, the date the execution stopped.
	StopDate *time.Time `locationName:"stopDate" type:"timestamp"`
}

// String returns the string representation
func (s ExecutionListItem) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the start of the execution.
type ExecutionStartedEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data input to the execution.
	Input *string `locationName:"input" type:"string" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the IAM role used for executing AWS Lambda
	// tasks.
	RoleArn *string `locationName:"roleArn" min:"1" type:"string"`
}

// String returns the string representation
func (s ExecutionStartedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the successful termination of the execution.
type ExecutionSucceededEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data output by the execution.
	Output *string `locationName:"output" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ExecutionSucceededEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the execution timeout that occurred during the execution.
type ExecutionTimedOutEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the timeout.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s ExecutionTimedOutEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the events of an execution.
type HistoryEvent struct {
	_ struct{} `type:"structure"`

	// Contains details about an activity that failed during an execution.
	ActivityFailedEventDetails *ActivityFailedEventDetails `locationName:"activityFailedEventDetails" type:"structure"`

	// Contains details about an activity schedule event that failed during an execution.
	ActivityScheduleFailedEventDetails *ActivityScheduleFailedEventDetails `locationName:"activityScheduleFailedEventDetails" type:"structure"`

	// Contains details about an activity scheduled during an execution.
	ActivityScheduledEventDetails *ActivityScheduledEventDetails `locationName:"activityScheduledEventDetails" type:"structure"`

	// Contains details about the start of an activity during an execution.
	ActivityStartedEventDetails *ActivityStartedEventDetails `locationName:"activityStartedEventDetails" type:"structure"`

	// Contains details about an activity that successfully terminated during an
	// execution.
	ActivitySucceededEventDetails *ActivitySucceededEventDetails `locationName:"activitySucceededEventDetails" type:"structure"`

	// Contains details about an activity timeout that occurred during an execution.
	ActivityTimedOutEventDetails *ActivityTimedOutEventDetails `locationName:"activityTimedOutEventDetails" type:"structure"`

	// Contains details about an abort of an execution.
	ExecutionAbortedEventDetails *ExecutionAbortedEventDetails `locationName:"executionAbortedEventDetails" type:"structure"`

	// Contains details about an execution failure event.
	ExecutionFailedEventDetails *ExecutionFailedEventDetails `locationName:"executionFailedEventDetails" type:"structure"`

	// Contains details about the start of the execution.
	ExecutionStartedEventDetails *ExecutionStartedEventDetails `locationName:"executionStartedEventDetails" type:"structure"`

	// Contains details about the successful termination of the execution.
	ExecutionSucceededEventDetails *ExecutionSucceededEventDetails `locationName:"executionSucceededEventDetails" type:"structure"`

	// Contains details about the execution timeout that occurred during the execution.
	ExecutionTimedOutEventDetails *ExecutionTimedOutEventDetails `locationName:"executionTimedOutEventDetails" type:"structure"`

	// The id of the event. Events are numbered sequentially, starting at one.
	//
	// Id is a required field
	Id *int64 `locationName:"id" type:"long" required:"true"`

	// Contains details about a lambda function that failed during an execution.
	LambdaFunctionFailedEventDetails *LambdaFunctionFailedEventDetails `locationName:"lambdaFunctionFailedEventDetails" type:"structure"`

	// Contains details about a failed lambda function schedule event that occurred
	// during an execution.
	LambdaFunctionScheduleFailedEventDetails *LambdaFunctionScheduleFailedEventDetails `locationName:"lambdaFunctionScheduleFailedEventDetails" type:"structure"`

	// Contains details about a lambda function scheduled during an execution.
	LambdaFunctionScheduledEventDetails *LambdaFunctionScheduledEventDetails `locationName:"lambdaFunctionScheduledEventDetails" type:"structure"`

	// Contains details about a lambda function that failed to start during an execution.
	LambdaFunctionStartFailedEventDetails *LambdaFunctionStartFailedEventDetails `locationName:"lambdaFunctionStartFailedEventDetails" type:"structure"`

	// Contains details about a lambda function that terminated successfully during
	// an execution.
	LambdaFunctionSucceededEventDetails *LambdaFunctionSucceededEventDetails `locationName:"lambdaFunctionSucceededEventDetails" type:"structure"`

	// Contains details about a lambda function timeout that occurred during an
	// execution.
	LambdaFunctionTimedOutEventDetails *LambdaFunctionTimedOutEventDetails `locationName:"lambdaFunctionTimedOutEventDetails" type:"structure"`

	// Contains details about an iteration of a Map state that was aborted.
	MapIterationAbortedEventDetails *MapIterationEventDetails `locationName:"mapIterationAbortedEventDetails" type:"structure"`

	// Contains details about an iteration of a Map state that failed.
	MapIterationFailedEventDetails *MapIterationEventDetails `locationName:"mapIterationFailedEventDetails" type:"structure"`

	// Contains details about an iteration of a Map state that was started.
	MapIterationStartedEventDetails *MapIterationEventDetails `locationName:"mapIterationStartedEventDetails" type:"structure"`

	// Contains details about an iteration of a Map state that succeeded.
	MapIterationSucceededEventDetails *MapIterationEventDetails `locationName:"mapIterationSucceededEventDetails" type:"structure"`

	// Contains details about Map state that was started.
	MapStateStartedEventDetails *MapStateStartedEventDetails `locationName:"mapStateStartedEventDetails" type:"structure"`

	// The id of the previous event.
	PreviousEventId *int64 `locationName:"previousEventId" type:"long"`

	// Contains details about a state entered during an execution.
	StateEnteredEventDetails *StateEnteredEventDetails `locationName:"stateEnteredEventDetails" type:"structure"`

	// Contains details about an exit from a state during an execution.
	StateExitedEventDetails *StateExitedEventDetails `locationName:"stateExitedEventDetails" type:"structure"`

	// Contains details about the failure of a task.
	TaskFailedEventDetails *TaskFailedEventDetails `locationName:"taskFailedEventDetails" type:"structure"`

	// Contains details about a task that was scheduled.
	TaskScheduledEventDetails *TaskScheduledEventDetails `locationName:"taskScheduledEventDetails" type:"structure"`

	// Contains details about a task that failed to start.
	TaskStartFailedEventDetails *TaskStartFailedEventDetails `locationName:"taskStartFailedEventDetails" type:"structure"`

	// Contains details about a task that was started.
	TaskStartedEventDetails *TaskStartedEventDetails `locationName:"taskStartedEventDetails" type:"structure"`

	// Contains details about a task that where the submit failed.
	TaskSubmitFailedEventDetails *TaskSubmitFailedEventDetails `locationName:"taskSubmitFailedEventDetails" type:"structure"`

	// Contains details about a submitted task.
	TaskSubmittedEventDetails *TaskSubmittedEventDetails `locationName:"taskSubmittedEventDetails" type:"structure"`

	// Contains details about a task that succeeded.
	TaskSucceededEventDetails *TaskSucceededEventDetails `locationName:"taskSucceededEventDetails" type:"structure"`

	// Contains details about a task that timed out.
	TaskTimedOutEventDetails *TaskTimedOutEventDetails `locationName:"taskTimedOutEventDetails" type:"structure"`

	// The date and time the event occurred.
	//
	// Timestamp is a required field
	Timestamp *time.Time `locationName:"timestamp" type:"timestamp" required:"true"`

	// The type of the event.
	//
	// Type is a required field
	Type HistoryEventType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s HistoryEvent) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a lambda function that failed during an execution.
type LambdaFunctionFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s LambdaFunctionFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a failed lambda function schedule event that occurred
// during an execution.
type LambdaFunctionScheduleFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s LambdaFunctionScheduleFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a lambda function scheduled during an execution.
type LambdaFunctionScheduledEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data input to the lambda function.
	Input *string `locationName:"input" type:"string" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the scheduled lambda function.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The maximum allowed duration of the lambda function.
	TimeoutInSeconds *int64 `locationName:"timeoutInSeconds" type:"long"`
}

// String returns the string representation
func (s LambdaFunctionScheduledEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a lambda function that failed to start during an execution.
type LambdaFunctionStartFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s LambdaFunctionStartFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a lambda function that successfully terminated during
// an execution.
type LambdaFunctionSucceededEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data output by the lambda function.
	Output *string `locationName:"output" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s LambdaFunctionSucceededEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a lambda function timeout that occurred during an
// execution.
type LambdaFunctionTimedOutEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the timeout.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s LambdaFunctionTimedOutEventDetails) String() string {
	return awsutil.Prettify(s)
}

type LogDestination struct {
	_ struct{} `type:"structure"`

	// An object describing a CloudWatch log group. For more information, see AWS::Logs::LogGroup
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html)
	// in the AWS CloudFormation User Guide.
	CloudWatchLogsLogGroup *CloudWatchLogsLogGroup `locationName:"cloudWatchLogsLogGroup" type:"structure"`
}

// String returns the string representation
func (s LogDestination) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LogDestination) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LogDestination"}
	if s.CloudWatchLogsLogGroup != nil {
		if err := s.CloudWatchLogsLogGroup.Validate(); err != nil {
			invalidParams.AddNested("CloudWatchLogsLogGroup", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type LoggingConfiguration struct {
	_ struct{} `type:"structure"`

	// An object that describes where your execution history events will be logged.
	// Limited to size 1. Required, if your log level is not set to OFF.
	Destinations []LogDestination `locationName:"destinations" type:"list"`

	// Determines whether execution history data is included in your log. When set
	// to FALSE, data is excluded.
	IncludeExecutionData *bool `locationName:"includeExecutionData" type:"boolean"`

	// Defines which category of execution history events are logged.
	Level LogLevel `locationName:"level" type:"string" enum:"true"`
}

// String returns the string representation
func (s LoggingConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LoggingConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LoggingConfiguration"}
	if s.Destinations != nil {
		for i, v := range s.Destinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Destinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about an iteration of a Map state.
type MapIterationEventDetails struct {
	_ struct{} `type:"structure"`

	// The index of the array belonging to the Map state iteration.
	Index *int64 `locationName:"index" type:"integer"`

	// The name of the iteration’s parent Map state.
	Name *string `locationName:"name" min:"1" type:"string"`
}

// String returns the string representation
func (s MapIterationEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Details about a Map state that was started.
type MapStateStartedEventDetails struct {
	_ struct{} `type:"structure"`

	// The size of the array for Map state iterations.
	Length *int64 `locationName:"length" type:"integer"`
}

// String returns the string representation
func (s MapStateStartedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a state entered during an execution.
type StateEnteredEventDetails struct {
	_ struct{} `type:"structure"`

	// The string that contains the JSON input data for the state.
	Input *string `locationName:"input" type:"string" sensitive:"true"`

	// The name of the state.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StateEnteredEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about an exit from a state during an execution.
type StateExitedEventDetails struct {
	_ struct{} `type:"structure"`

	// The name of the state.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The JSON output data of the state.
	Output *string `locationName:"output" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s StateExitedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the state machine.
type StateMachineListItem struct {
	_ struct{} `type:"structure"`

	// The date the state machine is created.
	//
	// CreationDate is a required field
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" required:"true"`

	// The name of the state machine.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that identifies the state machine.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`

	// Type is a required field
	Type StateMachineType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StateMachineListItem) String() string {
	return awsutil.Prettify(s)
}

// Tags are key-value pairs that can be associated with Step Functions state
// machines and activities.
//
// An array of key-value pairs. For more information, see Using Cost Allocation
// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide, and Controlling Access
// Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).
//
// Tags may only contain Unicode letters, digits, white space, or these symbols:
// _ . : / = + - @.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of a tag.
	Key *string `locationName:"key" min:"1" type:"string"`

	// The value of a tag.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about a task failure event.
type TaskFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a task scheduled during an execution.
type TaskScheduledEventDetails struct {
	_ struct{} `type:"structure"`

	// The JSON data passed to the resource referenced in a task state.
	//
	// Parameters is a required field
	Parameters *string `locationName:"parameters" type:"string" required:"true" sensitive:"true"`

	// The region of the scheduled task
	//
	// Region is a required field
	Region *string `locationName:"region" min:"1" type:"string" required:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`

	// The maximum allowed duration of the task.
	TimeoutInSeconds *int64 `locationName:"timeoutInSeconds" type:"long"`
}

// String returns the string representation
func (s TaskScheduledEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a task that failed to start during an execution.
type TaskStartFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskStartFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the start of a task during an execution.
type TaskStartedEventDetails struct {
	_ struct{} `type:"structure"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskStartedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a task that failed to submit during an execution.
type TaskSubmitFailedEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskSubmitFailedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a task submitted to a resource .
type TaskSubmittedEventDetails struct {
	_ struct{} `type:"structure"`

	// The response from a resource when a task has started.
	Output *string `locationName:"output" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskSubmittedEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about the successful completion of a task state.
type TaskSucceededEventDetails struct {
	_ struct{} `type:"structure"`

	// The full JSON response from a resource when a task has succeeded. This response
	// becomes the output of the related task.
	Output *string `locationName:"output" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskSucceededEventDetails) String() string {
	return awsutil.Prettify(s)
}

// Contains details about a resource timeout that occurred during an execution.
type TaskTimedOutEventDetails struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The service name of the resource in a task state.
	//
	// Resource is a required field
	Resource *string `locationName:"resource" min:"1" type:"string" required:"true"`

	// The action of the resource called by a task state.
	//
	// ResourceType is a required field
	ResourceType *string `locationName:"resourceType" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TaskTimedOutEventDetails) String() string {
	return awsutil.Prettify(s)
}
