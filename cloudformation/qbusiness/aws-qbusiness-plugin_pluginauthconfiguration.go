// Code generated by "go generate". Please don't change this file directly.

package qbusiness

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Plugin_PluginAuthConfiguration AWS CloudFormation Resource (AWS::QBusiness::Plugin.PluginAuthConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html
type Plugin_PluginAuthConfiguration struct {

	// BasicAuthConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-basicauthconfiguration
	BasicAuthConfiguration *Plugin_BasicAuthConfiguration `json:"BasicAuthConfiguration,omitempty"`

	// NoAuthConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-noauthconfiguration
	NoAuthConfiguration interface{} `json:"NoAuthConfiguration,omitempty"`

	// OAuth2ClientCredentialConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-oauth2clientcredentialconfiguration
	OAuth2ClientCredentialConfiguration *Plugin_OAuth2ClientCredentialConfiguration `json:"OAuth2ClientCredentialConfiguration,omitempty"`

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
func (r *Plugin_PluginAuthConfiguration) AWSCloudFormationType() string {
	return "AWS::QBusiness::Plugin.PluginAuthConfiguration"
}