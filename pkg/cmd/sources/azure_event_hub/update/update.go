package update

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"strconv"
)

func NewCmdAzureEventHubSourceUpdate(client *cip.APIClient) *cobra.Command {
	var (
		authorizationRuleName string
		category              string
		collectorId           string
		consumerGroup         string
		description           string
		eventHubKey           string
		eventHubName          string
		fieldNames            []string
		fieldValues           []string
		name                  string
		namespace             string
		sourceId              string
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates an Azure Event Hub source",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			updateEventHubSource(authorizationRuleName, category, collectorId, consumerGroup, description, eventHubKey,
				eventHubName, fieldNames, fieldValues, name, namespace, sourceId, client)
		},
	}
	cmd.Flags().StringVar(&authorizationRuleName, "authorizationRuleName", "", "Specify the name of the Event Hub Authorization Rule")
	cmd.Flags().StringVar(&category, "category", "", "Specify the source category for the source")
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id that the source is associated to")
	cmd.Flags().StringVar(&consumerGroup, "consumerGroup", "$Default", "Specify a custom event hub consumer group if required")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the source")
	cmd.Flags().StringVar(&eventHubKey, "eventHubKey", "", "Specify either the primary or secondary Event Hub key")
	cmd.Flags().StringVar(&eventHubName, "eventHubName", "", "Specify the name of the Event Hub")
	cmd.Flags().StringSliceVar(&fieldNames, "fieldNames", []string{}, "Specify the names of fields to add to the source "+
		"{names need to be comma separated e.g. field1,field2")
	cmd.Flags().StringSliceVar(&fieldValues, "fieldValues", []string{}, "Specify the values of fields to add to the source "+
		"(values need to be comma separated e.g. value1,value2")
	cmd.Flags().StringVar(&name, "name", "", "Specify the name for the source")
	cmd.Flags().StringVar(&namespace, "namespace", "", "Specify the name of the Event Hub Namespace")
	cmd.Flags().StringVar(&sourceId, "sourceId", "", "Specify the identifier of the source to update")
	cmd.MarkFlagRequired("authorizationRuleName")
	cmd.MarkFlagRequired("category")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("eventHubKey")
	cmd.MarkFlagRequired("eventHubName")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("sourceId")
	return cmd
}

func updateEventHubSource(authorizationRuleName string, category string, collectorId string, consumerGroup string, description string,
	eventHubKey string, eventHubName string, fieldNames []string, fieldValues []string, name string, namespace string, sourceId string,
	client *cip.APIClient) {
	sourceIdInt, err := strconv.Atoi(sourceId)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert string to int")
	}
	fields := cmdutils.GenerateFieldsMap(fieldNames, fieldValues)
	body := types.UpdateEventHubSourceRequest{
		Source: types.EventHubSourceUpdateModel{
			Id: sourceIdInt,
			SchemaRef: types.EventHubSourceSchema{
				Type: "Azure Event Hubs",
			},
			Config: types.EventHubSourceUpdateDefinition{
				Name:                    name,
				Description:             description,
				Namespace:               namespace,
				HubName:                 eventHubName,
				AccessPolicyName:        authorizationRuleName,
				AccessPolicyKey:         eventHubKey,
				ConsumerGroup:           consumerGroup,
				Fields:                  fields,
				Category:                category,
				ReceiveWithLatestOffset: true,
			},
			SourceType: "Universal",
		},
	}
	data, response, err := client.UpdateEventHubSource(body, collectorId, sourceId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
