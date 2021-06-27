package unlock

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdUnlockUser() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "unlock",
		Short: "Unlocks a Sumo Logic user account",
		Run: func(cmd *cobra.Command, args []string) {
			unlockUser(id)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user account to unlock")
	cmd.MarkFlagRequired("id")
	return cmd
}

func unlockUser(id string) {
	log := logging.GetConsoleLogger()
	requestUrl := "v1/users/" + id + "/unlock"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	if response.StatusCode != 204 {
		var responseError api.ResponseError
		err = json.Unmarshal(responseBody, &responseError)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal response body")
		}
	} else {
		fmt.Println("User account was unlocked.")
	}
}
