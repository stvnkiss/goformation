// Code generated by "go generate". Please don't change this file directly.

package internetmonitor

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Monitor_LocalHealthEventsConfig AWS CloudFormation Resource (AWS::InternetMonitor::Monitor.LocalHealthEventsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html
type Monitor_LocalHealthEventsConfig struct {

	// HealthScoreThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-healthscorethreshold
	HealthScoreThreshold *float64 `json:"HealthScoreThreshold,omitempty"`

	// MinTrafficImpact AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-mintrafficimpact
	MinTrafficImpact *float64 `json:"MinTrafficImpact,omitempty"`

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-status
	Status *string `json:"Status,omitempty"`

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
func (r *Monitor_LocalHealthEventsConfig) AWSCloudFormationType() string {
	return "AWS::InternetMonitor::Monitor.LocalHealthEventsConfig"
}
