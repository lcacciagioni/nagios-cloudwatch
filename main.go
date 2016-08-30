package main

import "flag"

func main() {
	critRangePtr := flag.String("c", "", "Critical if threshold is outside RANGE.")
	warnRangePtr := flag.String("w", "", "Warning if threshold is outside RANGE.")
	namespacePtr := flag.String("n", "", "Namespace for cloudwatch metric.")
	dimensionNamePtr := flag.String("dn", "", "Dimension name of cloudwatch metric.")
	dimensionValuePtr := flag.String("dv", "", "Dimension value of cloudwatch metric.")
	metricPtr := flag.String("m", "", "Metric name.")
	statisticPtr := flag.String("s", "Average", "Statistic used to evaluate metric. Options: Average,Sum,SampleCount,Maximum,Minimum.")
	periodPtr := flag.Int64("p", 60, "The period in seconds over which the statistic is applied.")
	regionPtr := flag.String("r", "us-west-1", "The AWS region to read metrics from.")
	unitPtr := flag.String("u", "None", "Unit in which it will be represented.")
	flag.Parse()

	// Gather the values
	datapoints, err := cfCheckMetric(regionPtr, metricPtr, namespacePtr, dimensionNamePtr, dimensionValuePtr, unitPtr, periodPtr)

	// Validate result
	nagiosCheck(warnRangePtr, critRangePtr, statisticPtr, datapoints, err)
}
