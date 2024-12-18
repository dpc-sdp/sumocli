package get_deployments

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGetDeployments(client *cip.APIClient) *cobra.Command {
	var parentDeploymentId string
	cmd := &cobra.Command{
		Use:   "get-deployments",
		Short: "Get deployment details where organizations can be created.",
		Run: func(cmd *cobra.Command, args []string) {
			getDeployment(client, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getDeployment(client *cip.APIClient, parentDeploymentId string) {
	data, response, err := client.GetDeployments(parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
