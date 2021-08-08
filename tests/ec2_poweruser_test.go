package test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEc2PowerUser(t *testing.T) {

	expusername := "ec2powerusertest"
	expectedpolicyname := "ec2poweruserpolicytest"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../moduleone",
	})
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	expectedjsonenc := terraform.Output(t, terraformOptions, "policydataoutput")
	expectedjsondec, err := url.QueryUnescape(expectedjsonenc)
	if err != nil {
		fmt.Println("Error", err)
		assert.Equal(t, 1, 0)
	}
	fmt.Println("expected : ", expectedjsondec)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create a IAM service client.
	svc := iam.New(sess)

	result, err := svc.GetUserPolicy(&iam.GetUserPolicyInput{
		PolicyName: &expectedpolicyname,
		UserName:   &expusername,
	})

	if err != nil {
		fmt.Println("Error", err)
		assert.Equal(t, 1, 0)
	}

	decodedValue, err := url.QueryUnescape(*result.PolicyDocument)
	if err != nil {
		fmt.Println("Error", err)
		assert.Equal(t, 1, 0)
	}
	fmt.Println("actual : ", decodedValue)
	assert.Equal(t, expectedjsondec, decodedValue)

}
