package get

import (
	"fmt"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdGet(client *cip.APIClient) *cobra.Command {
	var (
		contentId string
		path      string
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets a content item corresponding to the provided path or via the identifier of the content.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getContent(contentId, path, client)
		},
	}
	cmd.Flags().StringVar(&contentId, "contentId", "", "Specify the id of the content")
	cmd.Flags().StringVar(&path, "path", "", "Specify the path of the content you want to retrieve (e.g. /Library/Users/user@demo.com/SampleFolder)")
	return cmd
}

func getContent(contentId string, path string, client *cip.APIClient) {
	if contentId != "" {
		data, response, err := client.GetPathById(contentId)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(data, response, err, "")
		}
	} else if path != "" {
		data, response, err := client.GetItemByPath(path)
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			cmdutils.Output(data, response, err, "")
		}
	} else if contentId != "" && path != "" {
		fmt.Println("Please specify only contentId or path.")
	}
}
