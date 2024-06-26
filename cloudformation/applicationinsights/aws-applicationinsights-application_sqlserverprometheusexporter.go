// Code generated by "go generate". Please don't change this file directly.

package applicationinsights

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Application_SQLServerPrometheusExporter AWS CloudFormation Resource (AWS::ApplicationInsights::Application.SQLServerPrometheusExporter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html
type Application_SQLServerPrometheusExporter struct {

	// PrometheusPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html#cfn-applicationinsights-application-sqlserverprometheusexporter-prometheusport
	PrometheusPort string `json:"PrometheusPort"`

	// SQLSecretName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-sqlserverprometheusexporter.html#cfn-applicationinsights-application-sqlserverprometheusexporter-sqlsecretname
	SQLSecretName string `json:"SQLSecretName"`

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
func (r *Application_SQLServerPrometheusExporter) AWSCloudFormationType() string {
	return "AWS::ApplicationInsights::Application.SQLServerPrometheusExporter"
}
