---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_secret_sshkey Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating an ssh key type secret.
---

# harness_platform_secret_sshkey (Resource)

Resource for creating an ssh key type secret.

## Example Usage

```terraform
resource "harness_platform_secret_sshkey" "key_tab_file_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_key_tab_file_path_spec {
      key_path = "key_path"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "KeyTabFilePath"
  }
}

resource "harness_platform_secret_sshkey" " tgt_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  kerberos {
    tgt_password_spec {
      password = "account.${secret.id}"
    }
    principal             = "principal"
    realm                 = "realm"
    tgt_generation_method = "Password"
  }
}

resource "harness_platform_secret_sshkey" "sshkey_reference" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_reference_credential {
      user_name            = "user_name"
      key                  = "account.${key.id}"
      encrypted_passphrase = "account.${secret.id}"
    }
    credential_type = "KeyReference"
  }
}

resource "harness_platform_secret_sshkey" " sshkey_path" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    sshkey_path_credential {
      user_name            = "user_name"
      key_path             = "key_path"
      encrypted_passphrase = "encrypted_passphrase"
    }
    credential_type = "KeyPath"
  }
}

resource "harness_platform_secret_sshkey" "ssh_password" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]
  port        = 22
  ssh {
    ssh_password_credential {
      user_name = "user_name"
      password  = "account.${secret.id}"
    }
    credential_type = "Password"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `description` (String) Description of the resource.
- `kerberos` (Block List, Max: 1) Kerberos authentication scheme (see [below for nested schema](#nestedblock--kerberos))
- `org_id` (String) Unique identifier of the organization.
- `port` (Number) SSH port
- `project_id` (String) Unique identifier of the project.
- `ssh` (Block List, Max: 1) Kerberos authentication scheme (see [below for nested schema](#nestedblock--ssh))
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--kerberos"></a>
### Nested Schema for `kerberos`

Required:

- `principal` (String) Username to use for authentication.
- `realm` (String) Reference to a secret containing the password to use for authentication.

Optional:

- `tgt_generation_method` (String) Method to generate tgt
- `tgt_key_tab_file_path_spec` (Block List, Max: 1) Authenticate to App Dynamics using username and password. (see [below for nested schema](#nestedblock--kerberos--tgt_key_tab_file_path_spec))
- `tgt_password_spec` (Block List, Max: 1) Authenticate to App Dynamics using username and password. (see [below for nested schema](#nestedblock--kerberos--tgt_password_spec))

<a id="nestedblock--kerberos--tgt_key_tab_file_path_spec"></a>
### Nested Schema for `kerberos.tgt_key_tab_file_path_spec`

Optional:

- `key_path` (String) key path


<a id="nestedblock--kerberos--tgt_password_spec"></a>
### Nested Schema for `kerberos.tgt_password_spec`

Optional:

- `password` (String) password. To reference a password at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a password at the account scope, prefix 'account` to the expression: account.{identifier}



<a id="nestedblock--ssh"></a>
### Nested Schema for `ssh`

Required:

- `credential_type` (String) This specifies SSH credential type as Password, KeyPath or KeyReference

Optional:

- `ssh_password_credential` (Block List, Max: 1) SSH credential of type keyReference (see [below for nested schema](#nestedblock--ssh--ssh_password_credential))
- `sshkey_path_credential` (Block List, Max: 1) SSH credential of type keyPath (see [below for nested schema](#nestedblock--ssh--sshkey_path_credential))
- `sshkey_reference_credential` (Block List, Max: 1) SSH credential of type keyReference (see [below for nested schema](#nestedblock--ssh--sshkey_reference_credential))

<a id="nestedblock--ssh--ssh_password_credential"></a>
### Nested Schema for `ssh.ssh_password_credential`

Required:

- `password` (String) SSH Password. To reference a password at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a password at the account scope, prefix 'account` to the expression: account.{identifier}
- `user_name` (String) SSH Username.


<a id="nestedblock--ssh--sshkey_path_credential"></a>
### Nested Schema for `ssh.sshkey_path_credential`

Required:

- `key_path` (String) Path of the key file.
- `user_name` (String) SSH Username.

Optional:

- `encrypted_passphrase` (String) Encrypted Passphrase . To reference a encryptedPassphrase at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a encryptedPassPhrase at the account scope, prefix 'account` to the expression: account.{identifier}


<a id="nestedblock--ssh--sshkey_reference_credential"></a>
### Nested Schema for `ssh.sshkey_reference_credential`

Required:

- `key` (String) SSH key. To reference a key at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a key at the account scope, prefix 'account` to the expression: account.{identifier}
- `user_name` (String) SSH Username.

Optional:

- `encrypted_passphrase` (String) Encrypted Passphrase. To reference a encryptedPassphrase at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a encryptedPassPhrase at the account scope, prefix 'account` to the expression: account.{identifier}

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Import account level secret sshkey
terraform import harness_platform_secret_sshkey.example <secret_sshkey_id>

# Import org level secret sshkey
terraform import harness_platform_secret_sshkey.example <ord_id>/<secret_sshkey_id>

# Import project level secret sshkey
terraform import harness_platform_secret_sshkey.example <org_id>/<project_id>/<secret_sshkey_id>
```
