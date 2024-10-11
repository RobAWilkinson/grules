package rules

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

const ruleDef = `
rule CheckValues "Check the default values" salience 10 {
	when 
		MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
	then
		MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
		Retract("CheckValues");
}
rule CheckBooleanAttribute "Check the boolean attribute" salience 20 {
	when
		MF.BooleanAttribute == true
	then
		MF.WhatToSay = MF.GetWhatToSay("Boolean attribute is true");
		Retract("CheckBooleanAttribute");
}
`

func BuildKnowledgeBase() (*ast.KnowledgeBase, error) {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	
	bs := pkg.NewBytesResource([]byte(ruleDef))
	err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
	if err != nil {
		return nil, err
	}
	
	return knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")
}