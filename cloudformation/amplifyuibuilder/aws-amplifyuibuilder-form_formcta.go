// Code generated by "go generate". Please don't change this file directly.

package amplifyuibuilder

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Form_FormCTA AWS CloudFormation Resource (AWS::AmplifyUIBuilder::Form.FormCTA)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html
type Form_FormCTA struct {

	// Cancel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-cancel
	Cancel *Form_FormButton `json:"Cancel,omitempty"`

	// Clear AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-clear
	Clear *Form_FormButton `json:"Clear,omitempty"`

	// Position AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-position
	Position *string `json:"Position,omitempty"`

	// Submit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formcta.html#cfn-amplifyuibuilder-form-formcta-submit
	Submit *Form_FormButton `json:"Submit,omitempty"`

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
func (r *Form_FormCTA) AWSCloudFormationType() string {
	return "AWS::AmplifyUIBuilder::Form.FormCTA"
}
