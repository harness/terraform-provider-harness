# IDP integrations end-to-end example

`main.tf` creates IDP integrations, manages a kind mapping, and reads every IDP integration data source. It
exercises the full lifecycle: create, read, update, import and destroy.

| What | Created by default | Needs |
|------|--------------------|-------|
| `Custom` integration (`api` mode) at account scope | yes | nothing |
| Kind mapping for its `custom` kind | yes | nothing |
| `Custom` integration in `webhook` mode | with `webhook_secret_ref` | an HMAC secret |
| `HarnessK8s` integration in an org or project | with `k8s_agent_ref` | a service discovery agent and environment |
| `ServiceNow` integration, plus its catalog data source | with `servicenow_credentials_ref` | a ServiceNow credentials secret |

The credentials can be fake; syncs then fail, but the integrations are created and show in the UI. The references
must exist, though. `webhook_secret_ref` can name any secret, but `servicenow_credentials_ref` must be a ServiceNow
connector (`harness_platform_connector_service_now`), backed for example by a `harness_platform_secret_text` with a
dummy password.

Every identifier starts with `prefix` (default `tf_e2e`). If several people use one account, give each run its own prefix.

## 1. Build the provider

From the repository root:

```sh
mkdir -p /tmp/tfdev
go build -o /tmp/tfdev/terraform-provider-harness .
cat > /tmp/tfdev/terraformrc <<'EOF'
provider_installation {
  dev_overrides {
    "registry.terraform.io/harness/harness" = "/tmp/tfdev"
  }
  direct {}
}
EOF
export TF_CLI_CONFIG_FILE=/tmp/tfdev/terraformrc
```

`TF_CLI_CONFIG_FILE` uses the local build for this shell only, and leaves `~/.terraformrc` as it is. With dev overrides
there's no `terraform init`; Terraform prints a warning about the override, which is expected.

## 2. Point it at an environment

```sh
export HARNESS_ENDPOINT=https://app.harness.io/gateway        # or e.g. https://<devspace>.pr2.harness.io/gateway
export HARNESS_ACCOUNT_ID=<account id>
export HARNESS_PLATFORM_API_KEY=<pat or sat token>
```

The endpoint must end in `/gateway`. The provider calls `$HARNESS_ENDPOINT/integration/api/v1/...`.

To check the endpoint and key before running Terraform:

```sh
curl -s -H "x-api-key: $HARNESS_PLATFORM_API_KEY" \
  "$HARNESS_ENDPOINT/integration/api/v1/accounts/$HARNESS_ACCOUNT_ID/integration-types" | head -c 200
```

You should get JSON starting with `{"items":[`. An HTML page means the endpoint is wrong.

## 3. Run it

Run these from a copy of this directory, so the state file stays out of the repository:

```sh
cp -r examples/idp_integrations /tmp/idp_integrations && cd /tmp/idp_integrations

terraform apply                         # Custom integration and its mapping
terraform plan -detailed-exitcode       # exit code 0: no drift
```

Add optional integrations with variables, for example:

```sh
terraform apply \
  -var project_id=testproject \
  -var k8s_agent_ref=account.testinfra \
  -var k8s_environment=Test
```

Pass the same `-var` flags to every later `plan`, `apply`, `import` and `destroy`.

### Update

Change the integration's `name` or `tags`, or a mapping property in `main.tf`. Then run `terraform apply`, and
`terraform plan -detailed-exitcode` to check there's no drift. Changing the set of kinds, or the `idp_kind`/`idp_type`
of a kind, is rejected at plan time.

### Import

In a second copy of this directory with no state, import the resources, then plan. The plan should show no changes:

```sh
terraform import harness_platform_idp_integration.orders tf_e2e_orders
terraform import harness_platform_idp_integration_kind_mapping.orders tf_e2e_orders/custom
terraform import 'harness_platform_idp_integration.k8s[0]' default/testproject/tf_e2e_k8s   # if created
terraform plan -detailed-exitcode
```

Import IDs are `[<org>/[<project>/]]<integration>` for integrations, and `.../<integration>/<kind>` for mappings.

Delete that second copy afterwards rather than running `destroy` from it.

### Destroy

```sh
terraform destroy
```

Destroying a kind mapping leaves the mapping on the server as it is, because every kind always has one. Destroying the
integration removes it.

## Outputs

| Output | Value |
|--------|-------|
| `custom_kinds` | Kinds of the `Custom` type, from `harness_platform_idp_integration_type` |
| `orders_revision` | Revision after the mapping write; each write increments it |
| `orders_mapping_identifier` | The `ingestion.identifier` JQ read back through the mapping data source |
| `org_integrations` | Integrations in `org_id`/`project_id`, from `harness_platform_idp_integrations` |
| `webhook_urls` | Ingestion URLs, in webhook mode |
| `servicenow_streams` | Number of streams in the ServiceNow catalog, when ServiceNow is enabled |

## Troubleshooting

- **`422 VALIDATION_FAILED ... spec.webhook_auth: is required`**: webhook mode needs `webhook_auth`. Set
  `webhook_secret_ref`, or leave it empty to use `api` mode.
- **`422 ... spec.harness_k8s: is required`**: HarnessK8s needs `k8s_agent_ref` and `k8s_environment`.
- **`412 REVISION_MISMATCH`**: something else changed the integration since Terraform read it. Run `terraform apply`
  again.
- **HTML or 404 instead of JSON**: `HARNESS_ENDPOINT` doesn't end in `/gateway`.
