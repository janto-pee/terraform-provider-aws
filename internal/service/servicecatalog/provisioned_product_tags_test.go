// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalog_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccServiceCatalogProvisionedProduct_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1updated"),
					"tagKey2":   config.StringVariable("key2"),
					"tagValue2": config.StringVariable("value2"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1updated"),
					"tagKey2":   config.StringVariable("key2"),
					"tagValue2": config.StringVariable("value2"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":     config.StringVariable(rName),
			// 		"tagKey1":   config.StringVariable("key2"),
			// 		"tagValue1": config.StringVariable("value2"),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
			// 	),
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":     config.StringVariable(rName),
			// 		"tagKey1":   config.StringVariable("key2"),
			// 		"tagValue1": config.StringVariable("value2"),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 	),
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_null(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull/"),
				ConfigVariables: config.Variables{
					"rName":   config.StringVariable(rName),
					"tagKey1": config.StringVariable("key1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull/"),
				ConfigVariables: config.Variables{
					"rName":   config.StringVariable(rName),
					"tagKey1": config.StringVariable("key1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_AddOnUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnCreate(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable(""),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable(""),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
				ConfigVariables: config.Variables{
					"rName": config.StringVariable(rName),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnUpdate_Add(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
					"tagKey2":   config.StringVariable("key2"),
					"tagValue2": config.StringVariable(""),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
					"tagKey2":   config.StringVariable("key2"),
					"tagValue2": config.StringVariable(""),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_EmptyTag_OnUpdate_Replace(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy:             testAccCheckProvisionedProductDestroy(ctx),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable(""),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable(""),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default2/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1updated"),
					"providerTagKey2":   config.StringVariable("key2"),
					"providerTagValue2": config.StringVariable("value2"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default2/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1updated"),
					"providerTagKey2":   config.StringVariable("key2"),
					"providerTagValue2": config.StringVariable("value2"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":             config.StringVariable(rName),
			// 		"providerTagKey1":   config.StringVariable("key2"),
			// 		"providerTagValue1": config.StringVariable("value2"),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
			// 	),
			// },
			// {
			// 	ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName":             config.StringVariable(rName),
			// 		"providerTagKey1":   config.StringVariable("key2"),
			// 		"providerTagValue1": config.StringVariable("value2"),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
			// 	),
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("resourcekey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("resourcekey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1updated"),
					"tagKey1":           config.StringVariable("resourcekey1"),
					"tagValue1":         config.StringVariable("resourcevalue1updated"),
					"tagKey2":           config.StringVariable("resourcekey2"),
					"tagValue2":         config.StringVariable("resourcevalue2"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1updated"),
					"tagKey1":           config.StringVariable("resourcekey1"),
					"tagValue1":         config.StringVariable("resourcevalue1updated"),
					"tagKey2":           config.StringVariable("resourcekey2"),
					"tagValue2":         config.StringVariable("resourcevalue2"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckProvisionedProductExists(ctx, resourceName, &v),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
			// 	),
			// },
			// {
			// 	ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
			// 	ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags0/"),
			// 	ConfigVariables: config.Variables{
			// 		"rName": config.StringVariable(rName),
			// 	},
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateVerifyIgnore: []string{
			// 		"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
			// 	},
			// },
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_overlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2_default2/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"providerTagKey2":   config.StringVariable("overlapkey2"),
					"providerTagValue2": config.StringVariable("providervalue2"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
					"tagKey2":           config.StringVariable("overlapkey2"),
					"tagValue2":         config.StringVariable("resourcevalue2"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags2_default2/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"providerTagKey2":   config.StringVariable("overlapkey2"),
					"providerTagValue2": config.StringVariable("providervalue2"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue1"),
					"tagKey2":           config.StringVariable("overlapkey2"),
					"tagValue2":         config.StringVariable("resourcevalue2"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue2"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("overlapkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("overlapkey1"),
					"tagValue1":         config.StringVariable("resourcevalue2"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_updateToProviderOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags0_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/ProvisionedProduct/tags1/"),
				ConfigVariables: config.Variables{
					"rName":     config.StringVariable(rName),
					"tagKey1":   config.StringVariable("key1"),
					"tagValue1": config.StringVariable("value1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_emptyResourceTag(t *testing.T) {
	t.Skip("Resource ProvisionedProduct does not support empty tags")

	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
					"tagKey1":           config.StringVariable("key1"),
					"tagValue1":         config.StringVariable(""),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", ""),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tags1_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("value1"),
					"tagKey1":           config.StringVariable("key1"),
					"tagValue1":         config.StringVariable(""),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nullOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("key1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "providervalue1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("key1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("key1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}

func TestAccServiceCatalogProvisionedProduct_tags_DefaultTags_nullNonOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v servicecatalog.ProvisionedProductDetail
	resourceName := "aws_servicecatalog_provisioned_product.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.ServiceCatalogServiceID),
		CheckDestroy: testAccCheckProvisionedProductDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("resourcekey1"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProvisionedProductExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
				),
			},
			{
				ConfigDirectory: config.StaticDirectory("testdata/ProvisionedProduct/tagsNull_default1/"),
				ConfigVariables: config.Variables{
					"rName":             config.StringVariable(rName),
					"providerTagKey1":   config.StringVariable("providerkey1"),
					"providerTagValue1": config.StringVariable("providervalue1"),
					"tagKey1":           config.StringVariable("resourcekey1"),
				},
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"accept_language", "ignore_errors", "provisioning_artifact_name", "provisioning_parameters", "retain_physical_resources",
				},
			},
		},
	})
}
