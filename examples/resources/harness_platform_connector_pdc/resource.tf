# Only one host
resource "harness_platform_connector_pdc" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  host {
    hostname = "host1"
  }
}

# One host with attributes
resource "harness_platform_connector_pdc" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  host {
    hostname = "host1"
    attributes = {
      type        = "node"
      region      = "east"
      ip          = "54.87.11.191"
      anotherattr = "some value"
    }
  }
}

# Many hosts
resource "harness_platform_connector_pdc" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  host {
    hostname = "host1"
    attributes = {
      type        = "node"
      region      = "east"
      ip          = "54.87.11.191"
      anotherattr = "some value"
    }
  }
  host {
    hostname = "host2"
  }
  host {
    hostname = "host3"
  }
}
