// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_Eac3AtmosSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.Eac3AtmosSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html
type Channel_Eac3AtmosSettings struct {

	// Bitrate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-bitrate
	Bitrate *float64 `json:"Bitrate,omitempty"`

	// CodingMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-codingmode
	CodingMode *string `json:"CodingMode,omitempty"`

	// Dialnorm AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-dialnorm
	Dialnorm *int `json:"Dialnorm,omitempty"`

	// DrcLine AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-drcline
	DrcLine *string `json:"DrcLine,omitempty"`

	// DrcRf AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-drcrf
	DrcRf *string `json:"DrcRf,omitempty"`

	// HeightTrim AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-heighttrim
	HeightTrim *float64 `json:"HeightTrim,omitempty"`

	// SurroundTrim AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-eac3atmossettings.html#cfn-medialive-channel-eac3atmossettings-surroundtrim
	SurroundTrim *float64 `json:"SurroundTrim,omitempty"`

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
func (r *Channel_Eac3AtmosSettings) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.Eac3AtmosSettings"
}
