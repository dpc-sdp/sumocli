package sources

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdAwsCloudTrailSource "github.com/dpc-sdp/sumocli/pkg/cmd/sources/aws_cloudtrail"
	cmdAWSS3ArchiveSource "github.com/dpc-sdp/sumocli/pkg/cmd/sources/aws_s3_archive"
	cmdAzureEventHubSource "github.com/dpc-sdp/sumocli/pkg/cmd/sources/azure_event_hub"
	cmdSourcesDelete "github.com/dpc-sdp/sumocli/pkg/cmd/sources/delete"
	cmdHttpSources "github.com/dpc-sdp/sumocli/pkg/cmd/sources/http"
	cmdSourcesList "github.com/dpc-sdp/sumocli/pkg/cmd/sources/list"
	cmdLocalFileSources "github.com/dpc-sdp/sumocli/pkg/cmd/sources/local-file"
	"github.com/spf13/cobra"
)

func NewCmdSources(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sources",
		Short: "Manages sources assigned to collectors",
	}
	cmd.AddCommand(cmdAwsCloudTrailSource.NewCmdAWSCloudTrailSource())
	cmd.AddCommand(cmdAWSS3ArchiveSource.NewCmdAWSS3ArchiveSource())
	cmd.AddCommand(cmdAzureEventHubSource.NewCmdAzureEventHubSource(client))
	cmd.AddCommand(cmdSourcesDelete.NewCmdDeleteSource())
	cmd.AddCommand(cmdHttpSources.NewCmdHttpSources())
	cmd.AddCommand(cmdSourcesList.NewCmdSourceList())
	cmd.AddCommand(cmdLocalFileSources.NewCmdLocalFileSources())
	return cmd
}
