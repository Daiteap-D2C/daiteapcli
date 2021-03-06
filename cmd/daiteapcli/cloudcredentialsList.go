package daiteapcli

import (
	"fmt"
	"encoding/json"

	"github.com/Daiteap-D2C/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var cloudcredentialsListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
    Use:   "list",
    Aliases: []string{},
    Short:  "Command to list cloudcredentials from current tenant",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		method := "GET"
		endpoint := "/getCloudCredentials"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
    },
}

func init() {
    cloudcredentialsCmd.AddCommand(cloudcredentialsListCmd)
}