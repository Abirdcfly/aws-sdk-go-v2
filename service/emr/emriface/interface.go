// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package emriface provides an interface to enable mocking the Amazon Elastic MapReduce service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package emriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
)

// ClientAPI provides an interface to enable mocking the
// emr.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EMR.
//    func myFunc(svc emriface.ClientAPI) bool {
//        // Make svc.AddInstanceFleet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := emr.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        emriface.ClientPI
//    }
//    func (m *mockClientClient) AddInstanceFleet(input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AddInstanceFleetRequest(*emr.AddInstanceFleetInput) emr.AddInstanceFleetRequest

	AddInstanceGroupsRequest(*emr.AddInstanceGroupsInput) emr.AddInstanceGroupsRequest

	AddJobFlowStepsRequest(*emr.AddJobFlowStepsInput) emr.AddJobFlowStepsRequest

	AddTagsRequest(*emr.AddTagsInput) emr.AddTagsRequest

	CancelStepsRequest(*emr.CancelStepsInput) emr.CancelStepsRequest

	CreateSecurityConfigurationRequest(*emr.CreateSecurityConfigurationInput) emr.CreateSecurityConfigurationRequest

	DeleteSecurityConfigurationRequest(*emr.DeleteSecurityConfigurationInput) emr.DeleteSecurityConfigurationRequest

	DescribeClusterRequest(*emr.DescribeClusterInput) emr.DescribeClusterRequest

	DescribeJobFlowsRequest(*emr.DescribeJobFlowsInput) emr.DescribeJobFlowsRequest

	DescribeSecurityConfigurationRequest(*emr.DescribeSecurityConfigurationInput) emr.DescribeSecurityConfigurationRequest

	DescribeStepRequest(*emr.DescribeStepInput) emr.DescribeStepRequest

	GetBlockPublicAccessConfigurationRequest(*emr.GetBlockPublicAccessConfigurationInput) emr.GetBlockPublicAccessConfigurationRequest

	ListBootstrapActionsRequest(*emr.ListBootstrapActionsInput) emr.ListBootstrapActionsRequest

	ListClustersRequest(*emr.ListClustersInput) emr.ListClustersRequest

	ListInstanceFleetsRequest(*emr.ListInstanceFleetsInput) emr.ListInstanceFleetsRequest

	ListInstanceGroupsRequest(*emr.ListInstanceGroupsInput) emr.ListInstanceGroupsRequest

	ListInstancesRequest(*emr.ListInstancesInput) emr.ListInstancesRequest

	ListSecurityConfigurationsRequest(*emr.ListSecurityConfigurationsInput) emr.ListSecurityConfigurationsRequest

	ListStepsRequest(*emr.ListStepsInput) emr.ListStepsRequest

	ModifyClusterRequest(*emr.ModifyClusterInput) emr.ModifyClusterRequest

	ModifyInstanceFleetRequest(*emr.ModifyInstanceFleetInput) emr.ModifyInstanceFleetRequest

	ModifyInstanceGroupsRequest(*emr.ModifyInstanceGroupsInput) emr.ModifyInstanceGroupsRequest

	PutAutoScalingPolicyRequest(*emr.PutAutoScalingPolicyInput) emr.PutAutoScalingPolicyRequest

	PutBlockPublicAccessConfigurationRequest(*emr.PutBlockPublicAccessConfigurationInput) emr.PutBlockPublicAccessConfigurationRequest

	RemoveAutoScalingPolicyRequest(*emr.RemoveAutoScalingPolicyInput) emr.RemoveAutoScalingPolicyRequest

	RemoveTagsRequest(*emr.RemoveTagsInput) emr.RemoveTagsRequest

	RunJobFlowRequest(*emr.RunJobFlowInput) emr.RunJobFlowRequest

	SetTerminationProtectionRequest(*emr.SetTerminationProtectionInput) emr.SetTerminationProtectionRequest

	SetVisibleToAllUsersRequest(*emr.SetVisibleToAllUsersInput) emr.SetVisibleToAllUsersRequest

	TerminateJobFlowsRequest(*emr.TerminateJobFlowsInput) emr.TerminateJobFlowsRequest

	WaitUntilClusterRunning(context.Context, *emr.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilClusterTerminated(context.Context, *emr.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilStepComplete(context.Context, *emr.DescribeStepInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*emr.Client)(nil)
