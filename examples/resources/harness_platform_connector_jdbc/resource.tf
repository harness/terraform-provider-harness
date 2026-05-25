resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "ServiceAccount"
    service_account {
      token_ref = "account.secret_id"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "KeyPair"
    key_pair {
      username                   = "snowflake_user"
      private_key_file_ref       = "account.private_key_secret"
      private_key_passphrase_ref = "account.passphrase_secret"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "KeyPair"
    key_pair {
      username_ref               = "account.user_ref"
      private_key_file_ref       = "account.private_key_secret"
      private_key_passphrase_ref = "account.passphrase_secret"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "KeyPair"
    key_pair {
      username_ref               = "account.user_ref"
      private_key_file_ref       = "account.private_key_secret"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "KeyPair"
    key_pair {
      username                   = "admin"
      private_key_file_ref       = "account.private_key_secret"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "UsernamePassword"
    username_password {
      username     = "admin"
      password_ref = "account.secret_id"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "UsernamePassword"
    username_password {
      username_ref = "account.user_ref"
      password_ref = "account.secret_id"
    }
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type    = "UsernamePassword"
    username     = "admin"
    password_ref = "account.secret_id"
  }
}

resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type    = "UsernamePassword"
    username_ref = "account.user_ref"
    password_ref = "account.secret_id"
  }
}

resource "harness_platform_connector_jdbc" "oidc_example" {
  identifier         = "jdbc_oidc_gcp"
  name               = "JDBC OIDC GCP"
  description        = "JDBC connector using GCP OIDC authentication"
  tags               = ["foo:bar"]
  url                = "jdbc:postgresql://cloudsql-proxy:5432/mydb"
  delegate_selectors = ["harness-delegate"]
  credentials {
    auth_type = "Oidc"
    oidc {
      provider_type = "Gcp"
      gcp_oidc {
        project_number        = "145904791365"
        workload_pool_id      = "harness-identity-pool"
        provider_id           = "harness-oidc-provider"
        service_account_email = "db-sa@project.iam.gserviceaccount.com"
      }
    }
  }
}