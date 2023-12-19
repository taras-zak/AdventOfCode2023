package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	workflows := make(map[string][]WorkflowRule)
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			break
		}
		wf := ParseWorkflow(line)
		workflows[wf.Name] = wf.Rules
	}
	var parts []PartRating
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			break
		}
		part := ParsePart(line)
		parts = append(parts, part)
	}

	// Process part ratings
	totalRatingSum := 0
	for _, partRating := range parts {
		workflowName := "in"
		// Process through workflows
		for {
			rules, ok := workflows[workflowName]
			if !ok {
				break
			}

			workflowName = ApplyRules(partRating, rules)
			if workflowName == "" {
				break
			}
		}

		// Check if the final workflow is accepted
		if workflowName == "A" {
			ratingSum := partRating.X + partRating.M + partRating.A + partRating.S
			totalRatingSum += ratingSum
		}
	}

	fmt.Println("Your answer is:", totalRatingSum)
}

// ApplyRules applies the rules of a workflow to a part rating
func ApplyRules(partRating PartRating, rules []WorkflowRule) string {
	for _, rule := range rules {
		// Evaluate conditions
		if evaluateCondition(rule.Condition, partRating) {
			return rule.Action
		}
	}

	// Default action if no rule matches
	return ""
}

// evaluateCondition evaluates the condition of a rule for a given part rating
func evaluateCondition(condition string, partRating PartRating) bool {
	// Check if the condition is empty (rule without condition)
	if condition == "" {
		return true
	}

	// Assuming the conditions are simple comparisons (e.g., x>10)
	parts := strings.Split(condition, ">")
	if len(parts) == 2 {
		attribute := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		return compareValues(attribute, value, partRating)
	}

	// Handle "<" comparisons
	parts = strings.Split(condition, "<")
	if len(parts) == 2 {
		attribute := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		return !compareValues(attribute, value, partRating)
	}

	return false
}

// compareValues compares the attribute value with the given value based on the operator
func compareValues(attribute, value string, partRating PartRating) bool {
	switch attribute {
	case "x":
		return partRating.X > atoi(value)
	case "m":
		return partRating.M > atoi(value)
	case "a":
		return partRating.A > atoi(value)
	case "s":
		return partRating.S > atoi(value)
	}

	return false
}

// atoi converts a string to an integer
func atoi(s string) int {
	var val int
	_, _ = fmt.Sscanf(s, "%d", &val)
	return val
}

// Workflow represents a workflow with a name and rules
type Workflow struct {
	Name  string
	Rules []WorkflowRule
}

// WorkflowRule represents a single rule in the workflow
type WorkflowRule struct {
	Condition string
	Action    string
}

// ParseWorkflow parses the workflow string and returns a Workflow struct
func ParseWorkflow(workflow string) *Workflow {
	openBracketIndex := strings.Index(workflow, "{")
	closeBracketIndex := strings.Index(workflow, "}")

	if openBracketIndex == -1 || closeBracketIndex == -1 {
		// Invalid workflow format
		return nil
	}

	name := strings.TrimSpace(workflow[:openBracketIndex])
	rulesString := workflow[openBracketIndex+1 : closeBracketIndex]
	rules := parseRules(rulesString)

	return &Workflow{Name: name, Rules: rules}
}

// parseRules parses the rules from the rules string
func parseRules(rulesString string) []WorkflowRule {
	rules := strings.Split(rulesString, ",")
	var workflowRules []WorkflowRule

	for _, rule := range rules {
		parts := strings.SplitN(rule, ":", 2)
		if len(parts) == 2 {
			condition := strings.TrimSpace(parts[0])
			action := strings.TrimSpace(parts[1])
			workflowRules = append(workflowRules, WorkflowRule{Condition: condition, Action: action})
		} else if len(parts) == 1 {
			// Rule without condition, use an empty condition
			action := strings.TrimSpace(parts[0])
			workflowRules = append(workflowRules, WorkflowRule{Condition: "", Action: action})
		}
	}

	return workflowRules
}

// PartRating represents the rating of a part with x, m, a, and s values
type PartRating struct {
	X int
	M int
	A int
	S int
}

// ParsePart parses a part string and returns a PartRating struct
func ParsePart(partString string) PartRating {
	partString = strings.Trim(partString, "{}") // Remove curly braces
	parts := strings.Split(partString, ",")

	var partRating PartRating
	for _, part := range parts {
		attributeValue := strings.SplitN(part, "=", 2)
		if len(attributeValue) != 2 {
			panic(fmt.Errorf("invalid attribute-value format: %s", part))
		}

		attribute := strings.TrimSpace(attributeValue[0])
		value := strings.TrimSpace(attributeValue[1])

		switch attribute {
		case "x":
			partRating.X = atoi(value)
		case "m":
			partRating.M = atoi(value)
		case "a":
			partRating.A = atoi(value)
		case "s":
			partRating.S = atoi(value)
		default:
			panic(fmt.Errorf("unknown attribute: %s", attribute))
		}
	}
	return partRating
}
