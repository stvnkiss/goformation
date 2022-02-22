// Code generated by "go generate". Please don't change this file directly.

package appmesh

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// VirtualGateway_VirtualGatewayListenerTlsAcmCertificate AWS CloudFormation Resource (AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsAcmCertificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsacmcertificate.html
type VirtualGateway_VirtualGatewayListenerTlsAcmCertificate struct {

	// CertificateArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsacmcertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsacmcertificate-certificatearn
	CertificateArn string `json:"CertificateArn"`

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
func (r *VirtualGateway_VirtualGatewayListenerTlsAcmCertificate) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualGateway.VirtualGatewayListenerTlsAcmCertificate"
}
