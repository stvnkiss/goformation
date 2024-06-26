// Code generated by "go generate". Please don't change this file directly.

package s3

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// StorageLensGroup_MatchObjectSize AWS CloudFormation Resource (AWS::S3::StorageLensGroup.MatchObjectSize)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html
type StorageLensGroup_MatchObjectSize struct {

	// BytesGreaterThan AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-bytesgreaterthan
	BytesGreaterThan *int `json:"BytesGreaterThan,omitempty"`

	// BytesLessThan AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-byteslessthan
	BytesLessThan *int `json:"BytesLessThan,omitempty"`

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
func (r *StorageLensGroup_MatchObjectSize) AWSCloudFormationType() string {
	return "AWS::S3::StorageLensGroup.MatchObjectSize"
}
