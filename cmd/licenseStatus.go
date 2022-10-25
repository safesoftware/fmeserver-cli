package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type LicenseStatus struct {
	ExpiryDate       string `json:"expiryDate"`
	MaximumEngines   int    `json:"maximumEngines"`
	SerialNumber     string `json:"serialNumber"`
	IsLicenseExpired bool   `json:"isLicenseExpired"`
	IsLicensed       bool   `json:"isLicensed"`
	MaximumAuthors   int    `json:"maximumAuthors"`
}

// licenseStatusCmd represents the licenseStatus command
var licenseStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieves status of the installed FME Server license.",
	Long:  `Retrieves status of the installed FME Server license.`,
	Args:  NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// --json overrides --output
		if jsonOutput {
			outputType = "json"
		}
		// set up http
		client := &http.Client{}

		// call the status endpoint to see if it is finished
		request, err := buildFmeServerRequest("/fmerest/v3/licensing/license/status", "GET", nil)
		if err != nil {
			return err
		}
		response, err := client.Do(&request)
		if err != nil {
			return err
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		var result LicenseStatus
		if err := json.Unmarshal(responseData, &result); err != nil {
			return err
		} else {
			if outputType == "table" {
				// output all values returned by the JSON in a table
				t := createTableWithDefaultColumns(result)

				if noHeaders {
					t.ResetHeaders()
				}
				fmt.Println(t.Render())
			} else if outputType == "json" {
				prettyJSON, err := prettyPrintJSON(responseData)
				if err != nil {
					return err
				}
				fmt.Println(prettyJSON)
			} else if strings.HasPrefix(outputType, "custom-columns") {
				// parse the columns and json queries
				columnsString := ""
				if strings.HasPrefix(outputType, "custom-columns=") {
					columnsString = outputType[len("custom-columns="):]
				}
				if len(columnsString) == 0 {
					return errors.New("custom-columns format specified but no custom columns given")
				}

				// we have to marshal the Items array, then create an array of marshalled items
				// to pass to the creation of the table.
				marshalledItems := [][]byte{}
				mJson, err := json.Marshal(result)
				if err != nil {
					return err
				}
				marshalledItems = append(marshalledItems, mJson)

				columnsInput := strings.Split(columnsString, ",")
				t, err := createTableFromCustomColumns(marshalledItems, columnsInput)
				if err != nil {
					return err
				}
				if noHeaders {
					t.ResetHeaders()
				}
				fmt.Println(t.Render())
			} else {
				fmt.Println(string(responseData))
			}

		}
		return nil

	},
}

func init() {
	licenseCmd.AddCommand(licenseStatusCmd)
	licenseStatusCmd.Flags().StringVarP(&outputType, "output", "o", "table", "Specify the output type. Should be one of table, json, or custom-columns")
	licenseStatusCmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Don't print column headers")
}
