// Code generated by "go generate". Please don't change this file directly.

package athena

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// CapacityReservation_CapacityAssignmentConfiguration AWS CloudFormation Resource (AWS::Athena::CapacityReservation.CapacityAssignmentConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignmentconfiguration.html
type CapacityReservation_CapacityAssignmentConfiguration struct {

	// CapacityAssignments AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignmentconfiguration.html#cfn-athena-capacityreservation-capacityassignmentconfiguration-capacityassignments
	CapacityAssignments []CapacityReservation_CapacityAssignment `json:"CapacityAssignments"`

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
func (r *CapacityReservation_CapacityAssignmentConfiguration) AWSCloudFormationType() string {
	return "AWS::Athena::CapacityReservation.CapacityAssignmentConfiguration"
}
