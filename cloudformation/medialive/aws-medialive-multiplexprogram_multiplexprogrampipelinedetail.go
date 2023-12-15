// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Multiplexprogram_MultiplexProgramPipelineDetail AWS CloudFormation Resource (AWS::MediaLive::Multiplexprogram.MultiplexProgramPipelineDetail)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html
type Multiplexprogram_MultiplexProgramPipelineDetail struct {

	// ActiveChannelPipeline AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html#cfn-medialive-multiplexprogram-multiplexprogrampipelinedetail-activechannelpipeline
	ActiveChannelPipeline *string `json:"ActiveChannelPipeline,omitempty"`

	// PipelineId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html#cfn-medialive-multiplexprogram-multiplexprogrampipelinedetail-pipelineid
	PipelineId *string `json:"PipelineId,omitempty"`

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
func (r *Multiplexprogram_MultiplexProgramPipelineDetail) AWSCloudFormationType() string {
	return "AWS::MediaLive::Multiplexprogram.MultiplexProgramPipelineDetail"
}