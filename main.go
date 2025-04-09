package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/ollama/ollama/cmd"

	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	NewRelicAppName := os.Getenv("NEW_RELIC_APP_NAME")
	if NewRelicAppName == "" {
		log.Fatal("Environment variable NEW_RELIC_APP_NAME is not set")
	}

	NewRelicLicense := os.Getenv("NEW_RELIC_LICENSE")
	if NewRelicLicense == "" {
		log.Fatal("Environment variable NEW_RELIC_LICENSE is not set")
	}
	// Initialize a New Relic app
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(NewRelicAppName),
		newrelic.ConfigLicense(NewRelicLicense),
	)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Fatal(err)
	}

	cobra.CheckErr(cmd.NewCLI(app).ExecuteContext(context.Background()))
}
