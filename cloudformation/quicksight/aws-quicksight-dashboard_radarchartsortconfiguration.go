// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_RadarChartSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.RadarChartSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartsortconfiguration.html
type Dashboard_RadarChartSortConfiguration struct {

	// CategoryItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartsortconfiguration.html#cfn-quicksight-dashboard-radarchartsortconfiguration-categoryitemslimit
	CategoryItemsLimit *Dashboard_ItemsLimitConfiguration `json:"CategoryItemsLimit,omitempty"`

	// CategorySort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartsortconfiguration.html#cfn-quicksight-dashboard-radarchartsortconfiguration-categorysort
	CategorySort []Dashboard_FieldSortOptions `json:"CategorySort,omitempty"`

	// ColorItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartsortconfiguration.html#cfn-quicksight-dashboard-radarchartsortconfiguration-coloritemslimit
	ColorItemsLimit *Dashboard_ItemsLimitConfiguration `json:"ColorItemsLimit,omitempty"`

	// ColorSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartsortconfiguration.html#cfn-quicksight-dashboard-radarchartsortconfiguration-colorsort
	ColorSort []Dashboard_FieldSortOptions `json:"ColorSort,omitempty"`

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
func (r *Dashboard_RadarChartSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.RadarChartSortConfiguration"
}