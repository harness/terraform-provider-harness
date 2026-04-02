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