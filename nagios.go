// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/olorin/nagiosplugin"
)

// Abstraction of the check for warning or critical
func warningCritical(statistic *string, value *float64, check *nagiosplugin.Check, warningRange, criticalRange *nagiosplugin.Range) {
	check.AddResult(nagiosplugin.OK, "Values are normal")
	check.AddPerfDatum(*statistic, "", *value)

	if warningRange.Check(*value) {
		check.AddResultf(nagiosplugin.WARNING, "The %s is %.2f.", *statistic, *value)
	}
	if criticalRange.Check(*value) {
		check.AddResultf(nagiosplugin.CRITICAL, "The %s is %.2f.", *statistic, *value)
	}
}

func nagiosCheck(warning, critical, statistic *string, datapoints []*cloudwatch.Datapoint, awsErr error) {
	// Initialize the check - this will return an UNKNOWN result
	// until more results are added.
	check := nagiosplugin.NewCheck()
	// If we exit early or panic() we'll still output a result.
	defer check.Finish()

	// If the amazon error is not nil then fail with Critical
	if awsErr != nil {
		check.Criticalf("Cloudwatch check failed - %s", awsErr.Error())
	}

	// Validate ranges
	warnRange, err := nagiosplugin.ParseRange(*warning)
	if err != nil {
		check.AddResult(nagiosplugin.UNKNOWN, "Error parsing warning range")
	}
	critRange, err := nagiosplugin.ParseRange(*critical)
	if err != nil {
		check.AddResult(nagiosplugin.UNKNOWN, "Error parsing critical range")
	}

	// Check which metric we need to validate
	switch *statistic {
	case "Average":
		warningCritical(statistic, datapoints[0].Average, check, warnRange, critRange)
	case "Sum":
		warningCritical(statistic, datapoints[0].Sum, check, warnRange, critRange)
	case "SampleCount":
		warningCritical(statistic, datapoints[0].SampleCount, check, warnRange, critRange)
	case "Maximum":
		warningCritical(statistic, datapoints[0].Maximum, check, warnRange, critRange)
	case "Minimum":
		warningCritical(statistic, datapoints[0].Minimum, check, warnRange, critRange)
	default:
		check.Criticalf("Unknown statistic: %s", *statistic)
	}
}
