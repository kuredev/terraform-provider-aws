---
subcategory: "Organizations"
layout: "aws"
page_title: "AWS: aws_organizations_organization"
description: |-
  Get information about the organization that the user's account belongs to
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_organizations_organization

Get information about the organization that the user's account belongs to

## Example Usage

### List all account IDs for the organization

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformOutput, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_organizations_organization import DataAwsOrganizationsOrganization
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsOrganizationsOrganization(self, "example")
        TerraformOutput(self, "account_ids",
            value=Fn.lookup_nested(example.accounts, ["*", "id"])
        )
```

### SNS topic that can be interacted by the organization only

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_iam_policy_document import DataAwsIamPolicyDocument
from imports.aws.data_aws_organizations_organization import DataAwsOrganizationsOrganization
from imports.aws.sns_topic import SnsTopic
from imports.aws.sns_topic_policy import SnsTopicPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        sns_topic = SnsTopic(self, "sns_topic",
            name="my-sns-topic"
        )
        example = DataAwsOrganizationsOrganization(self, "example")
        sns_topic_policy = DataAwsIamPolicyDocument(self, "sns_topic_policy",
            statement=[DataAwsIamPolicyDocumentStatement(
                actions=["SNS:Subscribe", "SNS:Publish"],
                condition=[DataAwsIamPolicyDocumentStatementCondition(
                    test="StringEquals",
                    values=[Token.as_string(example.id)],
                    variable="aws:PrincipalOrgID"
                )
                ],
                effect="Allow",
                principals=[DataAwsIamPolicyDocumentStatementPrincipals(
                    identifiers=["*"],
                    type="AWS"
                )
                ],
                resources=[sns_topic.arn]
            )
            ]
        )
        aws_sns_topic_policy_sns_topic_policy = SnsTopicPolicy(self, "sns_topic_policy_3",
            arn=sns_topic.arn,
            policy=Token.as_string(sns_topic_policy.json)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sns_topic_policy_sns_topic_policy.override_logical_id("sns_topic_policy")
```

## Argument Reference

This data source does not support any arguments.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the organization.
* `feature_set` - FeatureSet of the organization.
* `id` - ID of the organization.
* `master_account_arn` - ARN of the account that is designated as the master account for the organization.
* `master_account_email` - The email address that is associated with the AWS account that is designated as the master account for the organization.
* `master_account_id` - Unique identifier (ID) of the master account of an organization.
* `master_account_name` - Name of the master account of an organization.

### Master Account or Delegated Administrator Attribute Reference

If the account is the master account or a delegated administrator for the organization, the following attributes are also exported:

* `accounts` - List of organization accounts including the master account. For a list excluding the master account, see the `non_master_accounts` attribute. All elements have these attributes:
    * `arn` - ARN of the account
    * `email` - Email of the account
    * `id` - Identifier of the account
    * `name` - Name of the account
    * `status` - Status of the account
* `aws_service_access_principals` - A list of AWS service principal names that have integration enabled with your organization. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
* `enabled_policy_types` - A list of Organizations policy types that are enabled in the Organization Root. Organization must have `feature_set` set to `ALL`. For additional information about valid policy types (e.g., `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
* `non_master_accounts` - List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
    * `arn` - ARN of the account
    * `email` - Email of the account
    * `id` - Identifier of the account
    * `name` - Name of the account
    * `status` - Status of the account
* `roots` - List of organization roots. All elements have these attributes:
    * `arn` - ARN of the root
    * `id` - Identifier of the root
    * `name` - Name of the root
    * `policy_types` - List of policy types enabled for this root. All elements have these attributes:
        * `name` - The name of the policy type
        * `status` - The status of the policy type as it relates to the associated root

<!-- cache-key: cdktf-0.20.8 input-dc97097a2b75109df7dc118299f41db93148ad09f52c2b8edb841c0e2bda1e0d -->