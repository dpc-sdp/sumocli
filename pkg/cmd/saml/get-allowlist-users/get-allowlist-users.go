package get_allowlist_users

import (
	"encoding/json"
	"fmt"
	"github.com/dpc-sdp/sumocli/api"
	"github.com/dpc-sdp/sumocli/pkg/cmd/factory"
	"github.com/dpc-sdp/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdSamlGetAllowListUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-allowlist-users",
		Short: "Get a list of allowlisted users.",
		Run: func(cmd *cobra.Command, args []string) {
			getAllowListUsers()
		},
	}
	return cmd
}

func getAllowListUsers() {
	var allowListResponse []api.GetSamlAllowListUsers
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/saml/allowlistedUsers"
	client, request := factory.NewHttpRequest("GET", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &allowListResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	allowListResponseJson, err := json.MarshalIndent(allowListResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(allowListResponseJson))
	}
}
