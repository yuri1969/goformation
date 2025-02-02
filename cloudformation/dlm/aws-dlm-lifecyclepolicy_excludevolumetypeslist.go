// Code generated by "go generate". Please don't change this file directly.

package dlm

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// LifecyclePolicy_ExcludeVolumeTypesList AWS CloudFormation Resource (AWS::DLM::LifecyclePolicy.ExcludeVolumeTypesList)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-excludevolumetypeslist.html
type LifecyclePolicy_ExcludeVolumeTypesList struct {

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
func (r *LifecyclePolicy_ExcludeVolumeTypesList) AWSCloudFormationType() string {
	return "AWS::DLM::LifecyclePolicy.ExcludeVolumeTypesList"
}
