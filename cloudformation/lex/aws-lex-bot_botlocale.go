package lex

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Bot_BotLocale AWS CloudFormation Resource (AWS::Lex::Bot.BotLocale)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html
type Bot_BotLocale struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-description
	Description string `json:"Description,omitempty"`

	// Intents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-intents
	Intents []Bot_Intent `json:"Intents,omitempty"`

	// LocaleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-localeid
	LocaleId string `json:"LocaleId,omitempty"`

	// NluConfidenceThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-nluconfidencethreshold
	NluConfidenceThreshold float64 `json:"NluConfidenceThreshold"`

	// SlotTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-slottypes
	SlotTypes []Bot_SlotType `json:"SlotTypes,omitempty"`

	// VoiceSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-botlocale.html#cfn-lex-bot-botlocale-voicesettings
	VoiceSettings *Bot_VoiceSettings `json:"VoiceSettings,omitempty"`

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
func (r *Bot_BotLocale) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.BotLocale"
}