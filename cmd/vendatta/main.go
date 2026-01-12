package main

import (
	"context"
	"github.com/IniZio/vendatta-config/pkg/metrics"
)

func main() {
	app := &cli.App{
		Name:  "vendatta",
		Usage: `Vendatta CLI for productivity optimization and analytics`,
		Commands: []cli.Command{
			{
				Name:  "usage",
				Aliases: []string{"u"},
				Usage:   "Show usage statistics and metrics",
				Action:  cli.UsageSummaryCommand(),
			},
			{
				Name:  "usage summary",
				Aliases: []string{"us"},
				Usage:   "Generate daily usage summary report",
				Action:  cli.UsageSummaryCommand(),
			},
			{
				Name:  "usage metrics",
				Aliases: []string{"um"},
				Usage:   "Calculate and display productivity metrics",
				Action:  cli.UsageMetricsCommand(),
			},
			{
				Name:  "usage patterns",
				Aliases: []string{"up"},
				Usage:   "Analyze usage patterns over time period",
				Action:  cli.UsagePatternsCommand(),
			},
			{
				Name:  "usage benchmark",
				Aliases: []string{"ub"},
				Usage:   "Compare current usage against baseline period",
				Action:  cli.UsageBenchmarkCommand(),
			},
		},
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
