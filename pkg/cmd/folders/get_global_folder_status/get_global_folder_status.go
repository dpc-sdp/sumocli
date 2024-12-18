package get_global_folder_status

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdGetGlobalFolderStatus(client *cip.APIClient) *cobra.Command {
	var jobId string

	cmd := &cobra.Command{
		Use:   "get-global-folder-status",
		Short: "Get the status of an asynchronous global folder job for the given job identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getGlobalFolderStatus(jobId, client)
		},
	}
	cmd.Flags().StringVar(&jobId, "jobId", "", "Specify the job id for the global folder (returned from running sumocli folders global-folder")
	cmd.MarkFlagRequired("jobId")
	return cmd
}

func getGlobalFolderStatus(jobId string, client *cip.APIClient) {
	data, response, err := client.GetGlobalFolderAsyncStatus(jobId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
