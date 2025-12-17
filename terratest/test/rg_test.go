package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestResourceGroupModule(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../resource_group",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}
