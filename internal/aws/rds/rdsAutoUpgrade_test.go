package rds

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/stangirard/yatas/internal/results"
	"github.com/stangirard/yatas/internal/yatas"
)

func Test_checkIfAutoUpgradeEnabled(t *testing.T) {
	type args struct {
		checkConfig yatas.CheckConfig
		instances   []types.DBInstance
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_checkIfAutoUpgradeEnabled",
			args: args{
				checkConfig: yatas.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan results.Check, 1),
				},
				instances: []types.DBInstance{
					{
						DBInstanceIdentifier:    aws.String("test"),
						DBInstanceArn:           aws.String("arn:aws:rds:us-east-1:123456789012:db:test"),
						StorageEncrypted:        true,
						AutoMinorVersionUpgrade: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkIfAutoUpgradeEnabled(tt.args.checkConfig, tt.args.instances, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Results[0].Status != "OK" {
						t.Errorf("CheckIfAutoUpgrade() = %v, want %v", check.Results[0].Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func Test_checkIfAutoUpgradeEnabledFail(t *testing.T) {
	type args struct {
		checkConfig yatas.CheckConfig
		instances   []types.DBInstance
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_checkIfAutoUpgradeEnabled",
			args: args{
				checkConfig: yatas.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan results.Check, 1),
				},
				instances: []types.DBInstance{
					{
						DBInstanceIdentifier:    aws.String("test"),
						DBInstanceArn:           aws.String("arn:aws:rds:us-east-1:123456789012:db:test"),
						StorageEncrypted:        true,
						AutoMinorVersionUpgrade: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkIfAutoUpgradeEnabled(tt.args.checkConfig, tt.args.instances, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Results[0].Status != "FAIL" {
						t.Errorf("CheckIfAutoUpgrade() = %v, want %v", check.Results[0].Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
