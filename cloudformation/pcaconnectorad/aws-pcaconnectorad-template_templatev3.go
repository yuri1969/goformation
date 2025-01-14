// Code generated by "go generate". Please don't change this file directly.

package pcaconnectorad

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_TemplateV3 AWS CloudFormation Resource (AWS::PCAConnectorAD::Template.TemplateV3)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html
type Template_TemplateV3 struct {

	// CertificateValidity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-certificatevalidity
	CertificateValidity *Template_CertificateValidity `json:"CertificateValidity"`

	// EnrollmentFlags AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-enrollmentflags
	EnrollmentFlags *Template_EnrollmentFlagsV3 `json:"EnrollmentFlags"`

	// Extensions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-extensions
	Extensions *Template_ExtensionsV3 `json:"Extensions"`

	// GeneralFlags AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-generalflags
	GeneralFlags *Template_GeneralFlagsV3 `json:"GeneralFlags"`

	// HashAlgorithm AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-hashalgorithm
	HashAlgorithm string `json:"HashAlgorithm"`

	// PrivateKeyAttributes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyattributes
	PrivateKeyAttributes *Template_PrivateKeyAttributesV3 `json:"PrivateKeyAttributes"`

	// PrivateKeyFlags AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyflags
	PrivateKeyFlags *Template_PrivateKeyFlagsV3 `json:"PrivateKeyFlags"`

	// SubjectNameFlags AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-subjectnameflags
	SubjectNameFlags *Template_SubjectNameFlagsV3 `json:"SubjectNameFlags"`

	// SupersededTemplates AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-supersededtemplates
	SupersededTemplates []string `json:"SupersededTemplates,omitempty"`

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
func (r *Template_TemplateV3) AWSCloudFormationType() string {
	return "AWS::PCAConnectorAD::Template.TemplateV3"
}
