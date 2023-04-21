// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_DateDimensionField AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DateDimensionField)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html
type Dashboard_DateDimensionField struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-column
	Column *Dashboard_ColumnIdentifier `json:"Column"`

	// DateGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-dategranularity
	DateGranularity *string `json:"DateGranularity,omitempty"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-fieldid
	FieldId string `json:"FieldId"`

	// FormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-formatconfiguration
	FormatConfiguration *Dashboard_DateTimeFormatConfiguration `json:"FormatConfiguration,omitempty"`

	// HierarchyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-hierarchyid
	HierarchyId *string `json:"HierarchyId,omitempty"`

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
func (r *Dashboard_DateDimensionField) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DateDimensionField"
}