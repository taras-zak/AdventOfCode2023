package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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

	initialRanges := PartRange{
		XRange: [2]int{1, 4000},
		MRange: [2]int{1, 4000},
		ARange: [2]int{1, 4000},
		SRange: [2]int{1, 4000},
	}

	fmt.Println("Your answer is:", countAcceptedCombinations(workflows, "in", initialRanges))
}

func countAcceptedCombinations(workflows map[string][]WorkflowRule, curr string, currRange PartRange) int {
	if curr == "R" {
		return 0
	}
	if curr == "A" {
		return currRange.count()
	}
	var count int
	for _, rule := range workflows[curr] {
		curr = rule.Action
		if rule.Op == 0 {
			count += countAcceptedCombinations(workflows, curr, currRange)
			continue
		}

		var interval [2]int
		switch rule.Rating {
		case "x":
			interval = currRange.XRange
		case "m":
			interval = currRange.MRange
		case "a":
			interval = currRange.ARange
		case "s":
			interval = currRange.SRange
		}

		var match [2]int
		var notMatch [2]int
		switch rule.Op {
		case '<':
			match = [2]int{interval[0], rule.Value - 1}
			notMatch = [2]int{rule.Value, interval[1]}
		case '>':
			match = [2]int{rule.Value + 1, interval[1]}
			notMatch = [2]int{interval[0], rule.Value}
		}
		if match[0] <= match[1] {
			count += countAcceptedCombinations(workflows, curr, currRange.updateInterval(rule.Rating, match))
		}
		if notMatch[0] > notMatch[1] {
			fmt.Println("notMatch", notMatch)
			break
		}
		currRange = currRange.updateInterval(rule.Rating, notMatch)
	}
	return count
}

// PartRange represents a range of possible values for each rating (x, m, a, s)
type PartRange struct {
	XRange [2]int
	MRange [2]int
	ARange [2]int
	SRange [2]int
}

func (r PartRange) updateInterval(attribute string, interval [2]int) PartRange {
	switch attribute {
	case "x":
		r.XRange = interval
	case "m":
		r.MRange = interval
	case "a":
		r.ARange = interval
	case "s":
		r.SRange = interval
	}
	return r
}

func (r PartRange) count() int {
	res := r.XRange[1] - r.XRange[0] + 1
	res *= r.MRange[1] - r.MRange[0] + 1
	res *= r.ARange[1] - r.ARange[0] + 1
	res *= r.SRange[1] - r.SRange[0] + 1
	return res
}

// Workflow represents a workflow with a name and rules
type Workflow struct {
	Name  string
	Rules []WorkflowRule
}

// WorkflowRule represents a single rule in the workflow
type WorkflowRule struct {
	Rating string
	Value  int
	Op     rune
	Action string
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
		r, err := parseRule(rule)
		if err != nil {
			panic(err)
		}
		workflowRules = append(workflowRules, r)
	}

	return workflowRules
}

// parseRule parses a rule into a WorkflowRule struct
func parseRule(ruleString string) (WorkflowRule, error) {
	var rule WorkflowRule

	// Split the rule string into rating, condition, and action
	parts := strings.SplitN(ruleString, ":", 2)
	if len(parts) != 2 {
		return WorkflowRule{Action: ruleString}, nil
	}

	ruleSegment := parts[0]
	action := parts[1]

	// Extract the rating, operator, and value from the rule segment
	rule.Rating, rule.Op, rule.Value, _ = parseRatingCondition(ruleSegment)
	rule.Action = action

	return rule, nil
}

// parseRatingCondition extracts the rating, operator, and value from a rule segment
func parseRatingCondition(ruleSegment string) (string, rune, int, error) {
	var rating string
	var op rune
	var value int

	// Find the first non-letter index
	index := strings.IndexFunc(ruleSegment, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	if index == -1 {
		return "", 0, 0, fmt.Errorf("invalid rule segment: %s", ruleSegment)
	}

	// Extract the rating
	rating = ruleSegment[:index]

	// Parse the remaining part
	_, err := fmt.Sscanf(ruleSegment[index:], "%c%d", &op, &value)
	if err != nil {
		return "", 0, 0, err
	}

	return rating, op, value, nil
}
