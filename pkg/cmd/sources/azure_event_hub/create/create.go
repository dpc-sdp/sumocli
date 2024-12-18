package create

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdAzureEventHubSourceCreate(client *cip.APIClient) *cobra.Command {
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
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates an Azure Event Hub source",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			createEventHubSource(authorizationRuleName, category, collectorId, consumerGroup, description, eventHubKey,
				eventHubName, fieldNames, fieldValues, name, namespace, client)
		},
	}
	cmd.Flags().StringVar(&authorizationRuleName, "authorizationRuleName", "", "Specify the name of the Event Hub Authorization Rule")
	cmd.Flags().StringVar(&category, "category", "", "Specify the source category for the source")
	cmd.Flags().StringVar(&collectorId, "collectorId", "", "Specify the collector id to associate the source to")
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
	cmd.MarkFlagRequired("authorizationRuleName")
	cmd.MarkFlagRequired("category")
	cmd.MarkFlagRequired("collectorId")
	cmd.MarkFlagRequired("eventHubKey")
	cmd.MarkFlagRequired("eventHubName")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("namespace")
	return cmd
}

func createEventHubSource(authorizationRuleName string, category string, collectorId string, consumerGroup string, description string,
	eventHubKey string, eventHubName string, fieldNames []string, fieldValues []string, name string, namespace string,
	client *cip.APIClient) {
	fields := cmdutils.GenerateFieldsMap(fieldNames, fieldValues)
	body := types.CreateEventHubSourceRequest{
		ApiVersion: "v1",
		Source: types.EventHubSource{
			SchemaRef: types.EventHubSourceSchema{
				Type: "Azure Event Hubs",
			},
			Config: types.EventHubSourceConfigurationDefinition{
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
	data, response, err := client.CreateEventHubSource(body, collectorId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
