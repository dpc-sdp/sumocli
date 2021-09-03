package get_admin_recommended_folder

import (
	"github.com/SumoLogic-Incubator/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Incubator/sumologic-go-sdk/service/cip/types"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewCmdGetAdminRecommendedFolder(client *cip.APIClient) *cobra.Command {
	var isAdminMode bool

	cmd := &cobra.Command{
		Use:   "get-admin-recommended-folder",
		Short: "Schedule an asynchronous job to get the top-level Admin Recommended content items.",
		Run: func(cmd *cobra.Command, args []string) {
			getAdminRecommendedFolder(isAdminMode, client)
		},
	}
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	return cmd
}

func getAdminRecommendedFolder(isAdminMode bool, client *cip.APIClient) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetAdminRecommendedFolderAsync(&types.FolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}