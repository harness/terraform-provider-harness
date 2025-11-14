
# ============================================================================
# ACCOUNT LEVEL TESTS (3 scenarios)
# ============================================================================

# 1. Account-level NTLM
resource "harness_platform_secret_text" "account_ntlm_password" {
  identifier                = "account_ntlm_password_v3"
  name                      = "account_ntlm_password_v3"
  description               = "Password for account-level NTLM"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "account_ntlm_pass"
}

resource "harness_platform_secret_winrm" "account_ntlm" {
  identifier  = "account_ntlm_v3"
  name        = "Account NTLM v3"
  description = "Account-level WinRM with NTLM"
  tags        = ["scope:account", "auth:ntlm"]

  port = 5986

  ntlm {
    domain          = "example.com"
    username        = "admin"
    password_ref    = "account.${harness_platform_secret_text.account_ntlm_password.id}"
    use_ssl         = true
    skip_cert_check = false
    use_no_profile  = true
  }
}

# 2. Account-level Kerberos with KeyTab
resource "harness_platform_secret_winrm" "account_kerberos_keytab" {
  identifier  = "account_kerberos_keytab_v3"
  name        = "Account Kerberos KeyTab v3"
  description = "Account-level WinRM with Kerberos KeyTab"
  tags        = ["scope:account", "auth:kerberos-keytab"]

  port = 5986

  kerberos {
    principal             = "service@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "KeyTabFilePath"
    use_ssl               = true
    skip_cert_check       = true
    use_no_profile        = true

    tgt_key_tab_file_path_spec {
      key_path = "/etc/krb5.keytab"
    }
  }
}

# 3. Account-level Kerberos with Password
resource "harness_platform_secret_text" "account_kerberos_password_1" {
  identifier                = "account_kerb_pass_20251111"
  name                      = "account_kerb_pass_20251111"
  description               = "Password for account-level Kerberos"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "account_kerberos_pass"
}

resource "harness_platform_secret_winrm" "account_kerberos_password_1" {
  identifier  = "account_kerb_winrm_20251111"
  name        = "Account Kerberos WinRM 20251111"
  description = "Account-level WinRM with Kerberos Password"
  tags        = ["scope:account", "auth:kerberos-password"]

  port = 5986

  kerberos {
    principal             = "user@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "Password"
    use_ssl               = true
    skip_cert_check       = false
    use_no_profile        = true

    tgt_password_spec {
      password_ref = "account.${harness_platform_secret_text.account_kerberos_password_1.id}"
    }
  }
}

# ============================================================================
# ORGANIZATION LEVEL TESTS (3 scenarios)
# ============================================================================

# 4. Org-level NTLM
resource "harness_platform_secret_text" "org_ntlm_password" {
  identifier                = "org_ntlm_password_v3"
  name                      = "org_ntlm_password_v3"
  description               = "Password for org-level NTLM"
  org_id                    = "default"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "org_ntlm_pass"
}

resource "harness_platform_secret_winrm" "org_ntlm" {
  identifier  = "org_ntlm_v3"
  name        = "Org NTLM v3"
  description = "Org-level WinRM with NTLM"
  org_id      = "default"
  tags        = ["scope:org", "auth:ntlm"]

  port = 5985

  ntlm {
    domain          = "org.example.com"
    username        = "orgadmin"
    password_ref    = "org.${harness_platform_secret_text.org_ntlm_password.id}"
    use_ssl         = false
    skip_cert_check = false
    use_no_profile  = true
  }
}

# 5. Org-level Kerberos with KeyTab
resource "harness_platform_secret_winrm" "org_kerberos_keytab" {
  identifier  = "org_kerberos_keytab_v3"
  name        = "Org Kerberos KeyTab v3"
  description = "Org-level WinRM with Kerberos KeyTab"
  org_id      = "default"
  tags        = ["scope:org", "auth:kerberos-keytab"]

  port = 5986

  kerberos {
    principal             = "orgservice@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "KeyTabFilePath"
    use_ssl               = true
    skip_cert_check       = true
    use_no_profile        = true

    tgt_key_tab_file_path_spec {
      key_path = "/etc/org.keytab"
    }
  }
}

# 6. Org-level Kerberos with Password
resource "harness_platform_secret_text" "org_kerberos_password" {
  identifier                = "org_kerb_pass_v3"
  name                      = "org_kerb_pass_v3"
  description               = "Password for org-level Kerberos"
  org_id                    = "default"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "org_kerberos_pass"
}

resource "harness_platform_secret_winrm" "org_kerberos_password" {
  identifier  = "org_kerb_winrm_v3"
  name        = "Org Kerberos WinRM v3"
  description = "Org-level WinRM with Kerberos Password"
  org_id      = "default"
  tags        = ["scope:org", "auth:kerberos-password"]

  port = 5986

  kerberos {
    principal             = "orguser@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "Password"
    use_ssl               = true
    skip_cert_check       = false
    use_no_profile        = true

    tgt_password_spec {
      password_ref = "org.${harness_platform_secret_text.org_kerberos_password.id}"
    }
  }
}

# ============================================================================
# PROJECT LEVEL TESTS (3 scenarios)
# ============================================================================

# 7. Project-level NTLM
resource "harness_platform_secret_text" "project_ntlm_password" {
  identifier                = "proj_ntlm_pass_v3"
  name                      = "proj_ntlm_pass_v3"
  description               = "Password for project-level NTLM"
  org_id                    = "default"
  project_id                = "winrm_support_terraform"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "project_ntlm_pass"
}

resource "harness_platform_secret_winrm" "project_ntlm" {
  identifier  = "proj_ntlm_winrm_v3"
  name        = "Project NTLM WinRM v3"
  description = "Project-level WinRM with NTLM"
  org_id      = "default"
  project_id  = "winrm_support_terraform"
  tags        = ["scope:project", "auth:ntlm"]

  port = 5986

  ntlm {
    domain          = "project.example.com"
    username        = "projectadmin"
    password_ref    = harness_platform_secret_text.project_ntlm_password.id
    use_ssl         = true
    skip_cert_check = false
    use_no_profile  = false
  }
}

# 8. Project-level Kerberos with KeyTab
resource "harness_platform_secret_winrm" "project_kerberos_keytab" {
  identifier  = "proj_kerb_keytab_v3"
  name        = "Project Kerberos KeyTab v3"
  description = "Project-level WinRM with Kerberos KeyTab"
  org_id      = "default"
  project_id  = "winrm_support_terraform"
  tags        = ["scope:project", "auth:kerberos-keytab"]

  port = 5986

  kerberos {
    principal             = "projectservice@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "KeyTabFilePath"
    use_ssl               = false
    skip_cert_check       = false
    use_no_profile        = false

    tgt_key_tab_file_path_spec {
      key_path = "/etc/project.keytab"
    }
  }
}

# 9. Project-level Kerberos with Password
resource "harness_platform_secret_text" "project_kerberos_password" {
  identifier                = "proj_kerb_pass_v3"
  name                      = "proj_kerb_pass_v3"
  description               = "Password for project-level Kerberos"
  org_id                    = "default"
  project_id                = "winrm_support_terraform"
  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = "project_kerberos_pass"
}

resource "harness_platform_secret_winrm" "project_kerberos_password" {
  identifier  = "proj_kerb_winrm_v3"
  name        = "Project Kerberos WinRM v3"
  description = "Project-level WinRM with Kerberos Password"
  org_id      = "default"
  project_id  = "winrm_support_terraform"
  tags        = ["scope:project", "auth:kerberos-password"]

  port = 5986

  kerberos {
    principal             = "projectuser@EXAMPLE.COM"
    realm                 = "EXAMPLE.COM"
    tgt_generation_method = "Password"
    use_ssl               = false
    skip_cert_check       = true
    use_no_profile        = true

    tgt_password_spec {
      password_ref = harness_platform_secret_text.project_kerberos_password.id
    }
  }
}