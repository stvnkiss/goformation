// Code generated by "go generate". Please don't change this file directly.

package datazone

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataSource_RedshiftStorage AWS CloudFormation Resource (AWS::DataZone::DataSource.RedshiftStorage)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html
type DataSource_RedshiftStorage struct {

	// RedshiftClusterSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html#cfn-datazone-datasource-redshiftstorage-redshiftclustersource
	RedshiftClusterSource *DataSource_RedshiftClusterStorage `json:"RedshiftClusterSource,omitempty"`

	// RedshiftServerlessSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftstorage.html#cfn-datazone-datasource-redshiftstorage-redshiftserverlesssource
	RedshiftServerlessSource *DataSource_RedshiftServerlessStorage `json:"RedshiftServerlessSource,omitempty"`

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
func (r *DataSource_RedshiftStorage) AWSCloudFormationType() string {
	return "AWS::DataZone::DataSource.RedshiftStorage"
}