// Code generated by "go generate". Please don't change this file directly.

package panorama

import (
	"bytes"
	"encoding/json"

	"github.com/awslabs/goformation/v6/cloudformation/policies"
	"github.com/awslabs/goformation/v6/cloudformation/tags"
)

// ApplicationInstance AWS CloudFormation Resource (AWS::Panorama::ApplicationInstance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html
type ApplicationInstance struct {

	// ApplicationInstanceIdToReplace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-applicationinstanceidtoreplace
	ApplicationInstanceIdToReplace *string `json:"ApplicationInstanceIdToReplace,omitempty"`

	// DefaultRuntimeContextDevice AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-defaultruntimecontextdevice
	DefaultRuntimeContextDevice string `json:"DefaultRuntimeContextDevice"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-description
	Description *string `json:"Description,omitempty"`

	// DeviceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-deviceid
	DeviceId *string `json:"DeviceId,omitempty"`

	// ManifestOverridesPayload AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestoverridespayload
	ManifestOverridesPayload *ApplicationInstance_ManifestOverridesPayload `json:"ManifestOverridesPayload,omitempty"`

	// ManifestPayload AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-manifestpayload
	ManifestPayload *ApplicationInstance_ManifestPayload `json:"ManifestPayload"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-name
	Name *string `json:"Name,omitempty"`

	// RuntimeRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-runtimerolearn
	RuntimeRoleArn *string `json:"RuntimeRoleArn,omitempty"`

	// StatusFilter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-statusfilter
	StatusFilter *string `json:"StatusFilter,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html#cfn-panorama-applicationinstance-tags
	Tags *[]tags.Tag `json:"Tags,omitempty"`

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
func (r *ApplicationInstance) AWSCloudFormationType() string {
	return "AWS::Panorama::ApplicationInstance"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r ApplicationInstance) MarshalJSON() ([]byte, error) {
	type Properties ApplicationInstance
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *ApplicationInstance) UnmarshalJSON(b []byte) error {
	type Properties ApplicationInstance
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           interface{}
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = ApplicationInstance(*res.Properties)
	}
	if res.DependsOn != nil {
		switch obj := res.DependsOn.(type) {
		case string:
			r.AWSCloudFormationDependsOn = []string{obj}
		case []interface{}:
			s := make([]string, 0, len(obj))
			for _, v := range obj {
				if value, ok := v.(string); ok {
					s = append(s, value)
				}
			}
			r.AWSCloudFormationDependsOn = s
		}
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
