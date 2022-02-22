// Code generated by "go generate". Please don't change this file directly.

package guardduty

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Detector_CFNDataSourceConfigurations AWS CloudFormation Resource (AWS::GuardDuty::Detector.CFNDataSourceConfigurations)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfndatasourceconfigurations.html
type Detector_CFNDataSourceConfigurations struct {

	// Kubernetes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfndatasourceconfigurations.html#cfn-guardduty-detector-cfndatasourceconfigurations-kubernetes
	Kubernetes *Detector_CFNKubernetesConfiguration `json:"Kubernetes,omitempty"`

	// S3Logs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfndatasourceconfigurations.html#cfn-guardduty-detector-cfndatasourceconfigurations-s3logs
	S3Logs *Detector_CFNS3LogsConfiguration `json:"S3Logs,omitempty"`

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
func (r *Detector_CFNDataSourceConfigurations) AWSCloudFormationType() string {
	return "AWS::GuardDuty::Detector.CFNDataSourceConfigurations"
}
