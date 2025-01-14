// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition_AuthorizationConfig AWS CloudFormation Resource (AWS::DevBatch::JobDefinition.AuthorizationConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-authorizationconfig.html
type JobDefinition_AuthorizationConfig struct {

	// AccessPointId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-authorizationconfig.html#cfn-devbatch-jobdefinition-authorizationconfig-accesspointid
	AccessPointId *string `json:"AccessPointId,omitempty"`

	// Iam AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devbatch-jobdefinition-authorizationconfig.html#cfn-devbatch-jobdefinition-authorizationconfig-iam
	Iam *string `json:"Iam,omitempty"`

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
func (r *JobDefinition_AuthorizationConfig) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition.AuthorizationConfig"
}
