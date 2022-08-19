package vpc

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stangirard/yatas/internal/logger"
	"github.com/stangirard/yatas/internal/results"
	"github.com/stangirard/yatas/internal/yatas"
)

func checkIfOnlyOneVPC(checkConfig yatas.CheckConfig, vpcs []types.Vpc, testName string) {
	logger.Info(fmt.Sprint("Running ", testName))
	var check results.Check
	check.InitCheck("VPC Only One", "Check if VPC has only one VPC", testName)
	for _, vpc := range vpcs {
		if len(vpcs) > 1 {
			Message := "VPC Id:" + *vpc.VpcId
			result := results.Result{Status: "FAIL", Message: Message, ResourceID: *vpc.VpcId}
			check.AddResult(result)
		} else {
			Message := "VPC Id:" + *vpc.VpcId
			result := results.Result{Status: "OK", Message: Message, ResourceID: *vpc.VpcId}
			check.AddResult(result)
		}
	}

	checkConfig.Queue <- check
}