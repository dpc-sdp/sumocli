package delete

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdFieldExtractionRulesDelete(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a field extraction rule with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteFieldExtractionRules(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the field extraction rule")
	cmd.MarkFlagRequired("id")
	return cmd
}

func deleteFieldExtractionRules(id string, client *cip.APIClient) {
	response, err := client.DeleteExtractionRule(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "Extraction rule was deleted successfully.")
	}
}
