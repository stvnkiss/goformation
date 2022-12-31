// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Project_ServiceCatalogProvisioningDetails AWS CloudFormation Resource (AWS::SageMaker::Project.ServiceCatalogProvisioningDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisioningdetails.html
type Project_ServiceCatalogProvisioningDetails struct {

	// PathId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisioningdetails.html#cfn-sagemaker-project-servicecatalogprovisioningdetails-pathid
	PathId *string `json:"PathId,omitempty"`

	// ProductId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisioningdetails.html#cfn-sagemaker-project-servicecatalogprovisioningdetails-productid
	ProductId string `json:"ProductId"`

	// ProvisioningArtifactId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisioningdetails.html#cfn-sagemaker-project-servicecatalogprovisioningdetails-provisioningartifactid
	ProvisioningArtifactId *string `json:"ProvisioningArtifactId,omitempty"`

	// ProvisioningParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisioningdetails.html#cfn-sagemaker-project-servicecatalogprovisioningdetails-provisioningparameters
	ProvisioningParameters []Project_ProvisioningParameter `json:"ProvisioningParameters,omitempty"`

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
func (r *Project_ServiceCatalogProvisioningDetails) AWSCloudFormationType() string {
	return "AWS::SageMaker::Project.ServiceCatalogProvisioningDetails"
}