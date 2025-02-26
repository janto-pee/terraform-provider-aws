---
subcategory: "Global Accelerator"
layout: "aws"
page_title: "AWS: aws_globalaccelerator_custom_routing_accelerator"
description: |-
  Provides a Global Accelerator custom routing accelerator data source.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_globalaccelerator_custom_routing_accelerator

Provides information about a Global Accelerator custom routing accelerator.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { VariableType, TerraformVariable, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsGlobalacceleratorCustomRoutingAccelerator } from "./.gen/providers/aws/data-aws-globalaccelerator-custom-routing-accelerator";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const acceleratorArn = new TerraformVariable(this, "accelerator_arn", {
      default: "",
      type: VariableType.STRING,
    });
    const acceleratorName = new TerraformVariable(this, "accelerator_name", {
      default: "",
      type: VariableType.STRING,
    });
    new DataAwsGlobalacceleratorCustomRoutingAccelerator(this, "example", {
      arn: acceleratorArn.stringValue,
      name: acceleratorName.stringValue,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `arn` - (Optional) Full ARN of the custom routing accelerator.
* `name` - (Optional) Unique name of the custom routing accelerator.

~> **NOTE:** When both `arn` and `name` are specified, `arn` takes precedence.

## Attribute Reference

See the [`aws_globalaccelerator_custom_routing_accelerator` resource](/docs/providers/aws/r/globalaccelerator_custom_routing_accelerator.html) for details on the
returned attributes - they are identical.

<!-- cache-key: cdktf-0.20.8 input-fb49b3a4aad72882bbe051a5e4ad191bf14607b4e514940f142786757b6b692c -->