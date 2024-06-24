// Code generated by "go generate". Please don't change this file directly.

package apptest

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TestCase_ResourceAction AWS CloudFormation Resource (AWS::AppTest::TestCase.ResourceAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html
type TestCase_ResourceAction struct {

	// CloudFormationAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-cloudformationaction
	CloudFormationAction *TestCase_CloudFormationAction `json:"CloudFormationAction,omitempty"`

	// M2ManagedApplicationAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-m2managedapplicationaction
	M2ManagedApplicationAction *TestCase_M2ManagedApplicationAction `json:"M2ManagedApplicationAction,omitempty"`

	// M2NonManagedApplicationAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-m2nonmanagedapplicationaction
	M2NonManagedApplicationAction *TestCase_M2NonManagedApplicationAction `json:"M2NonManagedApplicationAction,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *TestCase_ResourceAction) AWSCloudFormationType() string {
	return "AWS::AppTest::TestCase.ResourceAction"
}