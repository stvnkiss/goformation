// Code generated by "go generate". Please don't change this file directly.

package appflow

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Flow_RedshiftDestinationProperties AWS CloudFormation Resource (AWS::AppFlow::Flow.RedshiftDestinationProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html
type Flow_RedshiftDestinationProperties struct {

	// BucketPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-bucketprefix
	BucketPrefix *string `json:"BucketPrefix,omitempty"`

	// ErrorHandlingConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-errorhandlingconfig
	ErrorHandlingConfig *Flow_ErrorHandlingConfig `json:"ErrorHandlingConfig,omitempty"`

	// IntermediateBucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-intermediatebucketname
	IntermediateBucketName string `json:"IntermediateBucketName"`

	// Object AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-redshiftdestinationproperties.html#cfn-appflow-flow-redshiftdestinationproperties-object
	Object string `json:"Object"`

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
func (r *Flow_RedshiftDestinationProperties) AWSCloudFormationType() string {
	return "AWS::AppFlow::Flow.RedshiftDestinationProperties"
}
