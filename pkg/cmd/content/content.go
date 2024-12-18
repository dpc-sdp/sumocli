package content

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdContentGet "github.com/dpc-sdp/sumocli/pkg/cmd/content/get"
	cmdContentCopyStatus "github.com/dpc-sdp/sumocli/pkg/cmd/content/get_copy_status"
	cmdContentDeletionStatus "github.com/dpc-sdp/sumocli/pkg/cmd/content/get_deletion_status"
	cmdContentExportResult "github.com/dpc-sdp/sumocli/pkg/cmd/content/get_export_result"
	cmdContentExportStatus "github.com/dpc-sdp/sumocli/pkg/cmd/content/get_export_status"
	cmdContentImportStatus "github.com/dpc-sdp/sumocli/pkg/cmd/content/get_import_status"
	cmdContentMove "github.com/dpc-sdp/sumocli/pkg/cmd/content/move"
	cmdContentStartCopy "github.com/dpc-sdp/sumocli/pkg/cmd/content/start_copy"
	cmdContentStartDeletion "github.com/dpc-sdp/sumocli/pkg/cmd/content/start_deletion"
	cmdContentStartExport "github.com/dpc-sdp/sumocli/pkg/cmd/content/start_export"
	cmdContentStartImport "github.com/dpc-sdp/sumocli/pkg/cmd/content/start_import"
	"github.com/spf13/cobra"
)

func NewCmdContent(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content <command>",
		Short: "Manage content",
		Long:  "Commands that allow you to manage content in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdContentCopyStatus.NewCmdGetCopyStatus(client))
	cmd.AddCommand(cmdContentDeletionStatus.NewCmdGetDeletionStatus(client))
	cmd.AddCommand(cmdContentExportResult.NewCmdGetExportResult(client))
	cmd.AddCommand(cmdContentExportStatus.NewCmdExportStatus(client))
	cmd.AddCommand(cmdContentGet.NewCmdGet(client))
	cmd.AddCommand(cmdContentImportStatus.NewCmdGetImportStatus(client))
	cmd.AddCommand(cmdContentMove.NewCmdMove(client))
	cmd.AddCommand(cmdContentStartDeletion.NewCmdStartDeletion(client))
	cmd.AddCommand(cmdContentStartCopy.NewCmdStartCopy(client))
	cmd.AddCommand(cmdContentStartExport.NewCmdStartExport(client))
	cmd.AddCommand(cmdContentStartImport.NewCmdStartImport(client))
	return cmd
}
