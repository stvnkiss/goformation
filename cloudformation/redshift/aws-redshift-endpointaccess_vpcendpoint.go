// Code generated by "go generate". Please don't change this file directly.

package redshift

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// EndpointAccess_VpcEndpoint AWS CloudFormation Resource (AWS::Redshift::EndpointAccess.VpcEndpoint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html
type EndpointAccess_VpcEndpoint struct {

	// NetworkInterfaces AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-networkinterfaces
	NetworkInterfaces []EndpointAccess_NetworkInterface `json:"NetworkInterfaces,omitempty"`

	// VpcEndpointId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcendpointid
	VpcEndpointId *string `json:"VpcEndpointId,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcid
	VpcId *string `json:"VpcId,omitempty"`

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
func (r *EndpointAccess_VpcEndpoint) AWSCloudFormationType() string {
	return "AWS::Redshift::EndpointAccess.VpcEndpoint"
}