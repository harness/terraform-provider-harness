# Create Tas connector using username as plain text and password ref 

resource "harness_platform_connector_tas" "tas" {
  identifier  = "example_tas_cloud_provider"
  name        = "Example tas cloud provider"
  description = "description of tas connector"
  tags        = ["foo:bar"]
  credentials {
    type = "ManualConfig"
    tas_manual_details {
      endpoint_url = "https://tas.example.com"
      username = "admin"
      password_ref = "account.secret_id"
    }
  }
}


# Create Tas connector using username ref and password ref, and execute on delegate

resource "harness_platform_connector_tas" "tas" {
  identifier  = "example_tas_cloud_provider"
  name        = "Example tas cloud provider"
  description = "description of tas connector"
  tags        = ["foo:bar"]
  credentials {
    type = "ManualConfig"
    tas_manual_details {
      endpoint_url = "https://tas.example.com"
      username_ref = "account.username_id"
      password_ref = "account.secret_id"
    }
  }
  delegate_selectors = ["harness-delegate"]
  execute_on_delegate = true
}
