// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBotChannelMsTeamsInvalidResourceGroupNameRule checks the pattern is valid
type AzurermBotChannelMsTeamsInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBotChannelMsTeamsInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermBotChannelMsTeamsInvalidResourceGroupNameRule() *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule {
	return &AzurermBotChannelMsTeamsInvalidResourceGroupNameRule{
		resourceType:  "azurerm_bot_channel_ms_teams",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
	}
}

// Name returns the rule name
func (r *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule) Name() string {
	return "azurerm_bot_channel_ms_teams_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBotChannelMsTeamsInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
