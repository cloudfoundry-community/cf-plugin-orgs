package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/cloudfoundry/cli/plugin"
	"os"
)

type cfPluginOrgs struct{}

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, "error:", err)
		os.Exit(1)
	}
}

func (c *cfPluginOrgs) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "test-orgs" || args[0] == "torgs" {
		err := listOrganizations(cliConnection)
		fatalIf(err)
	} else {
		fmt.Print("Please provide an application name")
	}
}

//Small wrapper around Curl Get Command
func curlGet(cliConnection plugin.CliConnection, path string) ([]string, error) {
	return cliConnection.CliCommandWithoutTerminalOutput("curl", path, "-X", "GET")
}

//List all organization
func listOrganizations(cliConnection plugin.CliConnection) error {
	fmt.Println("Listing All organisation... ")
	output, err := curlGet(cliConnection, "/v2/organizations")
	if err != nil {
		return err
	}
	v, err := jason.NewObjectFromBytes([]byte(output[0]))
	if err != nil {
		return err
	}
	organizations, err := v.GetObjectArray("resources")
	if err != nil {
		return err
	}
	for _, org := range organizations {
		entity, err := org.GetObject("entity")
		if err != nil {
			return err
		}
		name, _ := entity.GetString("name")
		status, _ := entity.GetString("status")
		fmt.Printf("%s %s ", name, status)
		fmt.Println("")
	}
	return err
}

func (c *cfPluginOrgs) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "testPlugin",
		Commands: []plugin.Command{
			{
				Name:     "test-orgs",
				Alias:    "torgs",
				HelpText: "list ALL organisation",
			},
		},
	}
}

func main() {
	plugin.Start(new(cfPluginOrgs))
}
