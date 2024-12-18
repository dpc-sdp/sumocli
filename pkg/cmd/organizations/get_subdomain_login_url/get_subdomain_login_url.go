package get_subdomain_login_url

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdOrganizationsGetSubdomainLoginUrl(client *cip.APIClient) *cobra.Command {
	var (
		organizationId     string
		parentDeploymentId string
	)
	cmd := &cobra.Command{
		Use:   "get-subdomain-login-url",
		Short: "Get the login URL for the subdomain configured organization based on the organization identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			getSubdomainLoginUrl(client, organizationId, parentDeploymentId)
		},
	}
	cmd.Flags().StringVar(&organizationId, "organizationId", "", "Specify the identifier of the organization for which the details are required.")
	cmd.Flags().StringVar(&parentDeploymentId, "parentDeploymentId", "", "Specify the identifier of the deployment on which the calling organization resides.")
	cmd.MarkFlagRequired("organizationId")
	cmd.MarkFlagRequired("parentDeploymentId")
	return cmd
}

func getSubdomainLoginUrl(client *cip.APIClient, organizationId string, parentDeploymentId string) {
	data, response, err := client.GetSubdomainLoginUrl(organizationId, parentDeploymentId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
