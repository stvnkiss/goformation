// Code generated by "go generate". Please don't change this file directly.

package gamelift

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ContainerGroupDefinition_PortConfiguration AWS CloudFormation Resource (AWS::GameLift::ContainerGroupDefinition.PortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-portconfiguration.html
type ContainerGroupDefinition_PortConfiguration struct {

	// ContainerPortRanges AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-portconfiguration.html#cfn-gamelift-containergroupdefinition-portconfiguration-containerportranges
	ContainerPortRanges []ContainerGroupDefinition_ContainerPortRange `json:"ContainerPortRanges"`

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
func (r *ContainerGroupDefinition_PortConfiguration) AWSCloudFormationType() string {
	return "AWS::GameLift::ContainerGroupDefinition.PortConfiguration"
}