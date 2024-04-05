// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Agent_PromptOverrideConfiguration AWS CloudFormation Resource (AWS::Bedrock::Agent.PromptOverrideConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html
type Agent_PromptOverrideConfiguration struct {

	// OverrideLambda AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html#cfn-bedrock-agent-promptoverrideconfiguration-overridelambda
	OverrideLambda *string `json:"OverrideLambda,omitempty"`

	// PromptConfigurations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html#cfn-bedrock-agent-promptoverrideconfiguration-promptconfigurations
	PromptConfigurations []Agent_PromptConfiguration `json:"PromptConfigurations"`

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
func (r *Agent_PromptOverrideConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::Agent.PromptOverrideConfiguration"
}
