// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_SheetControlInfoIconLabelOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.SheetControlInfoIconLabelOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html
type Dashboard_SheetControlInfoIconLabelOptions struct {

	// InfoIconText AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-dashboard-sheetcontrolinfoiconlabeloptions-infoicontext
	InfoIconText *string `json:"InfoIconText,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.html#cfn-quicksight-dashboard-sheetcontrolinfoiconlabeloptions-visibility
	Visibility *string `json:"Visibility,omitempty"`

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
func (r *Dashboard_SheetControlInfoIconLabelOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.SheetControlInfoIconLabelOptions"
}
