// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestEchoFunction_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			//tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::testfunctions::echo("testvalue")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "testvalue"),
				),
			},
		},
	})
}

func TestEchoFunction_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			//tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::testfunctions::echo(null)
				}
				`,
				// The parameter does not enable AllowNullValue
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestEchoFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			//tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = "testvalue"
				}
				
				output "test" {
					value = provider::testfunctions::echo(terraform_data.test.output)
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "testvalue"),
				),
			},
		},
	})
}
