// Code generated by "go generate". Please don't change this file directly.

package verifiedpermissions

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// PolicyStore_SchemaDefinition AWS CloudFormation Resource (AWS::VerifiedPermissions::PolicyStore.SchemaDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html
type PolicyStore_SchemaDefinition struct {

	// CedarJson AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html#cfn-verifiedpermissions-policystore-schemadefinition-cedarjson
	CedarJson *string `json:"CedarJson,omitempty"`

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
func (r *PolicyStore_SchemaDefinition) AWSCloudFormationType() string {
	return "AWS::VerifiedPermissions::PolicyStore.SchemaDefinition"
}
