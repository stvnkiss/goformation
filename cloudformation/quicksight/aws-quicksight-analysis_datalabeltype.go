// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_DataLabelType AWS CloudFormation Resource (AWS::QuickSight::Analysis.DataLabelType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html
type Analysis_DataLabelType struct {

	// DataPathLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html#cfn-quicksight-analysis-datalabeltype-datapathlabeltype
	DataPathLabelType *Analysis_DataPathLabelType `json:"DataPathLabelType,omitempty"`

	// FieldLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html#cfn-quicksight-analysis-datalabeltype-fieldlabeltype
	FieldLabelType *Analysis_FieldLabelType `json:"FieldLabelType,omitempty"`

	// MaximumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html#cfn-quicksight-analysis-datalabeltype-maximumlabeltype
	MaximumLabelType *Analysis_MaximumLabelType `json:"MaximumLabelType,omitempty"`

	// MinimumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html#cfn-quicksight-analysis-datalabeltype-minimumlabeltype
	MinimumLabelType *Analysis_MinimumLabelType `json:"MinimumLabelType,omitempty"`

	// RangeEndsLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datalabeltype.html#cfn-quicksight-analysis-datalabeltype-rangeendslabeltype
	RangeEndsLabelType *Analysis_RangeEndsLabelType `json:"RangeEndsLabelType,omitempty"`

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
func (r *Analysis_DataLabelType) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DataLabelType"
}
