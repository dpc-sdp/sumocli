package field_extraction_rules

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	NewCmdFieldExtractionRulesCreate "github.com/dpc-sdp/sumocli/pkg/cmd/field_extraction_rules/create"
	NewCmdFieldExtractionRulesDelete "github.com/dpc-sdp/sumocli/pkg/cmd/field_extraction_rules/delete"
	NewCmdFieldExtractionRulesGet "github.com/dpc-sdp/sumocli/pkg/cmd/field_extraction_rules/get"
	NewCmdFieldExtractionRulesList "github.com/dpc-sdp/sumocli/pkg/cmd/field_extraction_rules/list"
	NewCmdFieldExtractionRulesUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/field_extraction_rules/update"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRules(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "field-extraction-rules",
		Short: "Manage field extraction rules",
		Long:  "Field Extraction Rules allow you to parse fields from your log messages at the time the messages are ingested eliminating the need to parse fields in your query.",
	}
	cmd.AddCommand(NewCmdFieldExtractionRulesCreate.NewCmdFieldExtractionRulesCreate(client))
	cmd.AddCommand(NewCmdFieldExtractionRulesDelete.NewCmdFieldExtractionRulesDelete(client))
	cmd.AddCommand(NewCmdFieldExtractionRulesGet.NewCmdFieldExtractionRulesGet(client))
	cmd.AddCommand(NewCmdFieldExtractionRulesList.NewCmdFieldExtractionRulesList(client))
	cmd.AddCommand(NewCmdFieldExtractionRulesUpdate.NewCmdFieldExtractionRulesUpdate(client))
	return cmd
}
