# nagios-cloudwatch

Simple implementation in golang of a cloudwatch monitor for nagios / icinga and any nagios compatible monitor.

## Install

```bash
$ go get -u github.com/lcacciagioni/nagios-cloudwatch
```

## Usage

```bash
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ nagios-cloudwatch -c "400" -w "300" -dn DBInstanceIdentifier -dv bob-production-master -r sa-east-1 -p 120 -m WriteIOPS -n "AWS/RDS" -u "Count/Second"
OK: Values are normal | Average=250.375;;;;
```

## Options
```
Usage of ./nagios-cloudwatch:
  -c string
        Critical if threshold is outside RANGE
  -dn string
        Dimension name of cloudwatch metric.
  -dv string
        Dimension value of cloudwatch metric.
  -m string
        Metric name.
  -n string
        Namespace for cloudwatch metric.
  -p int
        The period in seconds over which the statistic is applied. (default 60)
  -r string
        The AWS region to read metrics from. (default "us-west-1")
  -s string
        Statistic used to evaluate metric. Options: Average,Sum,SampleCount,Maximum,Minimum. (default "Average")
  -u string
        Unit in which it will be represented. (default "None")
  -w string
        Warning if threshold is outside RANGE.
```

> Information:
> * Nagios Ranges: https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT
> * AWS Namespaces, Dimensions and Metrics: http://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CW_Support_For_AWS.html

## License

```
This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3 of the License, or
(at your option) any later version.
```

## TODO

- [ ] Testing
- [ ] Nagios units in PerfData (https://nagios-plugins.org/doc/guidelines.html#AEN200)
- [ ] Binary Release

### Last Words

Hope this software can help you... And remember that any help is always welcome. Report BUGS and features request [here](https://github.com/lcacciagioni/nagios-cloudwatch/issues).
