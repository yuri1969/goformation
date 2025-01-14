// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// RuleGroup_Allow AWS CloudFormation Resource (AWS::WAFv2::RuleGroup.Allow)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-allow.html
type RuleGroup_Allow struct {

	// CustomRequestHandling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-allow.html#cfn-wafv2-rulegroup-allow-customrequesthandling
	CustomRequestHandling *RuleGroup_CustomRequestHandling `json:"CustomRequestHandling,omitempty"`

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
func (r *RuleGroup_Allow) AWSCloudFormationType() string {
	return "AWS::WAFv2::RuleGroup.Allow"
}
