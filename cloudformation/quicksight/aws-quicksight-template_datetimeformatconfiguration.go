// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_DateTimeFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.DateTimeFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeformatconfiguration.html
type Template_DateTimeFormatConfiguration struct {

	// DateTimeFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeformatconfiguration.html#cfn-quicksight-template-datetimeformatconfiguration-datetimeformat
	DateTimeFormat *string `json:"DateTimeFormat,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeformatconfiguration.html#cfn-quicksight-template-datetimeformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Template_NullValueFormatConfiguration `json:"NullValueFormatConfiguration,omitempty"`

	// NumericFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeformatconfiguration.html#cfn-quicksight-template-datetimeformatconfiguration-numericformatconfiguration
	NumericFormatConfiguration *Template_NumericFormatConfiguration `json:"NumericFormatConfiguration,omitempty"`

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
func (r *Template_DateTimeFormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DateTimeFormatConfiguration"
}
