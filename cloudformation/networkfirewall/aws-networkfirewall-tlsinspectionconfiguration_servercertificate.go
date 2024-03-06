// Code generated by "go generate". Please don't change this file directly.

package networkfirewall

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TLSInspectionConfiguration_ServerCertificate AWS CloudFormation Resource (AWS::NetworkFirewall::TLSInspectionConfiguration.ServerCertificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificate.html
type TLSInspectionConfiguration_ServerCertificate struct {

	// ResourceArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificate.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificate-resourcearn
	ResourceArn *string `json:"ResourceArn,omitempty"`

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
func (r *TLSInspectionConfiguration_ServerCertificate) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::TLSInspectionConfiguration.ServerCertificate"
}