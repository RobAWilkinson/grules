package main

import (
	"fmt"
	"log"
	"time"

	"easi-grules/internal/facts"
	"easi-grules/internal/rules"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

func main() {
	myFact := createMyFact()
	dataCtx := createDataContext(myFact)
	knowledgeBase := buildKnowledgeBase()

	if err := executeRules(dataCtx, knowledgeBase); err != nil {
		log.Fatalf("Error executing rules: %v", err)
	}

	fmt.Println(myFact.WhatToSay)
}

func createMyFact() *facts.MyFact {
	return &facts.MyFact{
		IntAttribute:     1234,
		StringAttribute:  "Some string value",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}
}

func createDataContext(fact *facts.MyFact) *ast.DataContext {
	dataCtx := ast.NewDataContext()
	if err := dataCtx.Add("MF", fact); err != nil {
		log.Fatalf("Error adding fact to data context: %v", err)
	}
	return dataCtx.(*ast.DataContext)
}

func buildKnowledgeBase() *ast.KnowledgeBase {
	knowledgeBase, err := rules.BuildKnowledgeBase()
	if err != nil {
		log.Fatalf("Error building knowledge base: %v", err)
	}
	return knowledgeBase
}

func executeRules(dataCtx *ast.DataContext, knowledgeBase *ast.KnowledgeBase) error {
	engine := engine.NewGruleEngine()
	return engine.Execute(dataCtx, knowledgeBase)
}
