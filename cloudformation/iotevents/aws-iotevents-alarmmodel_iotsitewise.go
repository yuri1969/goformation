// Code generated by "go generate". Please don't change this file directly.

package iotevents

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// AlarmModel_IotSiteWise AWS CloudFormation Resource (AWS::IoTEvents::AlarmModel.IotSiteWise)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html
type AlarmModel_IotSiteWise struct {

	// AssetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-assetid
	AssetId *string `json:"AssetId,omitempty"`

	// EntryId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-entryid
	EntryId *string `json:"EntryId,omitempty"`

	// PropertyAlias AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyalias
	PropertyAlias *string `json:"PropertyAlias,omitempty"`

	// PropertyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyid
	PropertyId *string `json:"PropertyId,omitempty"`

	// PropertyValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyvalue
	PropertyValue *AlarmModel_AssetPropertyValue `json:"PropertyValue"`

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
func (r *AlarmModel_IotSiteWise) AWSCloudFormationType() string {
	return "AWS::IoTEvents::AlarmModel.IotSiteWise"
}