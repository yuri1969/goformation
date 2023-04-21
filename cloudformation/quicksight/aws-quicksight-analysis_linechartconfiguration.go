// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_LineChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.LineChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html
type Analysis_LineChartConfiguration struct {

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Analysis_ContributionAnalysisDefault `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-datalabels
	DataLabels *Analysis_DataLabelOptions `json:"DataLabels,omitempty"`

	// DefaultSeriesSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-defaultseriessettings
	DefaultSeriesSettings *Analysis_LineChartDefaultSeriesSettings `json:"DefaultSeriesSettings,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-fieldwells
	FieldWells *Analysis_LineChartFieldWells `json:"FieldWells,omitempty"`

	// ForecastConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-forecastconfigurations
	ForecastConfigurations []Analysis_ForecastConfiguration `json:"ForecastConfigurations,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-legend
	Legend *Analysis_LegendOptions `json:"Legend,omitempty"`

	// PrimaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-primaryyaxisdisplayoptions
	PrimaryYAxisDisplayOptions *Analysis_LineSeriesAxisDisplayOptions `json:"PrimaryYAxisDisplayOptions,omitempty"`

	// PrimaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-primaryyaxislabeloptions
	PrimaryYAxisLabelOptions *Analysis_ChartAxisLabelOptions `json:"PrimaryYAxisLabelOptions,omitempty"`

	// ReferenceLines AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-referencelines
	ReferenceLines []Analysis_ReferenceLine `json:"ReferenceLines,omitempty"`

	// SecondaryYAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-secondaryyaxisdisplayoptions
	SecondaryYAxisDisplayOptions *Analysis_LineSeriesAxisDisplayOptions `json:"SecondaryYAxisDisplayOptions,omitempty"`

	// SecondaryYAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-secondaryyaxislabeloptions
	SecondaryYAxisLabelOptions *Analysis_ChartAxisLabelOptions `json:"SecondaryYAxisLabelOptions,omitempty"`

	// Series AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-series
	Series []Analysis_SeriesItem `json:"Series,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Analysis_SmallMultiplesOptions `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-sortconfiguration
	SortConfiguration *Analysis_LineChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions `json:"Tooltip,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-type
	Type *string `json:"Type,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-visualpalette
	VisualPalette *Analysis_VisualPalette `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Analysis_AxisDisplayOptions `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartconfiguration.html#cfn-quicksight-analysis-linechartconfiguration-xaxislabeloptions
	XAxisLabelOptions *Analysis_ChartAxisLabelOptions `json:"XAxisLabelOptions,omitempty"`

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
func (r *Analysis_LineChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.LineChartConfiguration"
}