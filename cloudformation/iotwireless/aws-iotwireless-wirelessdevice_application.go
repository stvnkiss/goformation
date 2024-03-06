// Code generated by "go generate". Please don't change this file directly.

package iotwireless

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// WirelessDevice_Application AWS CloudFormation Resource (AWS::IoTWireless::WirelessDevice.Application)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html
type WirelessDevice_Application struct {

	// DestinationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-destinationname
	DestinationName *string `json:"DestinationName,omitempty"`

	// FPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-fport
	FPort *int `json:"FPort,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-type
	Type *string `json:"Type,omitempty"`

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
func (r *WirelessDevice_Application) AWSCloudFormationType() string {
	return "AWS::IoTWireless::WirelessDevice.Application"
}