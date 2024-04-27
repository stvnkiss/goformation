// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_DefaultRelativeDateTimeControlOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DefaultRelativeDateTimeControlOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultrelativedatetimecontroloptions.html
type Dashboard_DefaultRelativeDateTimeControlOptions struct {

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultrelativedatetimecontroloptions.html#cfn-quicksight-dashboard-defaultrelativedatetimecontroloptions-displayoptions
	DisplayOptions *Dashboard_RelativeDateTimeControlDisplayOptions `json:"DisplayOptions,omitempty"`

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
func (r *Dashboard_DefaultRelativeDateTimeControlOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DefaultRelativeDateTimeControlOptions"
}
