package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func cfCheckMetric(region, metricName, namespace, dimensionName, dimensionValue, unit *string, period *int64) ([]*cloudwatch.Datapoint, error) {
	startTime := time.Now()
	endTime := startTime.Add(-time.Second * time.Duration(*period))

	// Here we call to the svc creation for cloudwatch
	svc := cloudwatch.New(session.New(&aws.Config{Region: aws.String(*region)}))

	// Parameters required for getting the info
	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(startTime),     // Required
		MetricName: aws.String(*metricName), // Required
		Namespace:  aws.String(*namespace),  // Required
		Period:     aws.Int64(*period),      // Required
		StartTime:  aws.Time(endTime),       // Required
		Statistics: []*string{ // Required
			aws.String("Average"),
			aws.String("Sum"),
			aws.String("Minimum"),
			aws.String("Maximum"),
			aws.String("SampleCount"),
		},
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String(*dimensionName),  // Required
				Value: aws.String(*dimensionValue), // Required
			},
			// More values...
		},
		Unit: aws.String(*unit),
	}

	resp, err := svc.GetMetricStatistics(params)
	return resp.Datapoints, err
}
