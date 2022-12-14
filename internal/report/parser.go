package report

import (
	"io/ioutil"

	"github.com/padok-team/yatas/plugins/commons"
	"gopkg.in/yaml.v3"
)

func parseReportYaml(filename string) ([]commons.Tests, error) {
	var report []commons.Tests
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return report, err
	}
	err = yaml.Unmarshal(data, &report)
	return report, err
}

var fullName = map[string]string{
	"S3":  "S3 Bucket",
	"VOL": "Volume",
	"BAK": "Backup",
	"RDS": "RDS",
	"VPC": "VPC",
	"COG": "Cognito",
	"CLD": "CloudTrail",
	"ECR": "ECR",
	"LMD": "Lambda",
	"DYN": "DynamoDB",
	"EC2": "EC2",
	"IAM": "IAM",
	"CFT": "Cloudfront",
	"APG": "APIGateway",
	"ASG": "AutoScaling",
	"ELB": "LoadBalancer",
	"GDT": "GuardDuty",
	"SHU": "SecurityHub",
	"ACM": "AWS Certificate Manager",
}
