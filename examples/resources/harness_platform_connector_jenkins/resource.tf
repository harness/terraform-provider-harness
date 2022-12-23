# Auth mechanism username password
resource "harness_platform_connector_jenkins" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  jenkins_url        = "https://jenkinss.com/"
  delegate_selectors = ["harness-delegate"]
  auth {
    type = "UsernamePassword"
    jenkins_user_name_password {
      username     = "username"
      password_ref = "account.${harness_platform_secret_text.test.id}"
    }
  }
}

# Auth mechanism anonymous
resource "harness_platform_connector_jenkins" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  jenkins_url        = "https://jenkinss.com/"
  delegate_selectors = ["harness-delegate"]
  auth {
    type = "Anonymous"
  }
}

# Auth mechanism bearer token
resource "harness_platform_connector_jenkins" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  jenkins_url        = "https://jenkinss.com/"
  delegate_selectors = ["harness-delegate"]
  auth {
    type = "Bearer Token(HTTP Header)"
    jenkins_bearer_token {
      token_ref = "account.${harness_platform_secret_text.test.id}"
    }
  }
}
