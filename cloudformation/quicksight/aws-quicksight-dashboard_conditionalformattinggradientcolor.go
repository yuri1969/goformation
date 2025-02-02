// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_ConditionalFormattingGradientColor AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ConditionalFormattingGradientColor)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html
type Dashboard_ConditionalFormattingGradientColor struct {

	// Color AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-color
	Color *Dashboard_GradientColor `json:"Color"`

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-conditionalformattinggradientcolor.html#cfn-quicksight-dashboard-conditionalformattinggradientcolor-expression
	Expression string `json:"Expression"`

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
func (r *Dashboard_ConditionalFormattingGradientColor) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ConditionalFormattingGradientColor"
}
