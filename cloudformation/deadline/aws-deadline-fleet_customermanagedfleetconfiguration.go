// Code generated by "go generate". Please don't change this file directly.

package deadline

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Fleet_CustomerManagedFleetConfiguration AWS CloudFormation Resource (AWS::Deadline::Fleet.CustomerManagedFleetConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html
type Fleet_CustomerManagedFleetConfiguration struct {

	// Mode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-mode
	Mode string `json:"Mode"`

	// StorageProfileId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-storageprofileid
	StorageProfileId *string `json:"StorageProfileId,omitempty"`

	// WorkerCapabilities AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-workercapabilities
	WorkerCapabilities *Fleet_CustomerManagedWorkerCapabilities `json:"WorkerCapabilities"`

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
func (r *Fleet_CustomerManagedFleetConfiguration) AWSCloudFormationType() string {
	return "AWS::Deadline::Fleet.CustomerManagedFleetConfiguration"
}