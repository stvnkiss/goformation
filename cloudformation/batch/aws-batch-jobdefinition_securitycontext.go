// Code generated by "go generate". Please don't change this file directly.

package batch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_SecurityContext AWS CloudFormation Resource (AWS::Batch::JobDefinition.SecurityContext)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html
type JobDefinition_SecurityContext struct {

	// Privileged AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html#cfn-batch-jobdefinition-ekscontainer-securitycontext-privileged
	Privileged *bool `json:"Privileged,omitempty"`

	// ReadOnlyRootFilesystem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html#cfn-batch-jobdefinition-ekscontainer-securitycontext-readonlyrootfilesystem
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty"`

	// RunAsGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html#cfn-batch-jobdefinition-ekscontainer-securitycontext-runasgroup
	RunAsGroup *int `json:"RunAsGroup,omitempty"`

	// RunAsNonRoot AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html#cfn-batch-jobdefinition-ekscontainer-securitycontext-runasnonroot
	RunAsNonRoot *bool `json:"RunAsNonRoot,omitempty"`

	// RunAsUser AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainer-securitycontext.html#cfn-batch-jobdefinition-ekscontainer-securitycontext-runasuser
	RunAsUser *int `json:"RunAsUser,omitempty"`

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
func (r *JobDefinition_SecurityContext) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.SecurityContext"
}
