// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_ForecastConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.ForecastConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastconfiguration.html
type Template_ForecastConfiguration struct {

	// ForecastProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastconfiguration.html#cfn-quicksight-template-forecastconfiguration-forecastproperties
	ForecastProperties *Template_TimeBasedForecastProperties `json:"ForecastProperties,omitempty"`

	// Scenario AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-forecastconfiguration.html#cfn-quicksight-template-forecastconfiguration-scenario
	Scenario *Template_ForecastScenario `json:"Scenario,omitempty"`

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
func (r *Template_ForecastConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ForecastConfiguration"
}
