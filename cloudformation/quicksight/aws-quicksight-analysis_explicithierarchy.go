// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_ExplicitHierarchy AWS CloudFormation Resource (AWS::QuickSight::Analysis.ExplicitHierarchy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-explicithierarchy.html
type Analysis_ExplicitHierarchy struct {

	// Columns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-explicithierarchy.html#cfn-quicksight-analysis-explicithierarchy-columns
	Columns []Analysis_ColumnIdentifier `json:"Columns"`

	// DrillDownFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-explicithierarchy.html#cfn-quicksight-analysis-explicithierarchy-drilldownfilters
	DrillDownFilters []Analysis_DrillDownFilter `json:"DrillDownFilters,omitempty"`

	// HierarchyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-explicithierarchy.html#cfn-quicksight-analysis-explicithierarchy-hierarchyid
	HierarchyId string `json:"HierarchyId"`

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
func (r *Analysis_ExplicitHierarchy) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ExplicitHierarchy"
}