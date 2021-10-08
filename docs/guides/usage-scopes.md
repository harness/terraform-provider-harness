---
subcategory: ""
page_title: "Configuring usage scopes - Harness Provider"
description: |-
    An example of how to apply usage scopes to a resource.
---

# Configure usage scopes for a resource

There are a number of resources that can be scoped to a specific set of applications and environments. These include cloud providers, secrets, connectors, and more. Configuring it is the same across all of these resources.

In this example we are configuring a cloud provider to be used by any application in any environment.

```terraform
resource "harness_cloudprovider_kubernetes" "test" {
  name = "test"

  usage_scope {
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }

  usage_scope {
    environment_filter_type = "PRODUCTION_ENVIRONMENTS"
  }

}
```

In this more advanced scenario we show how you can scope the cloud provider to a specific application or to a specific environment.

```terraform
resource "harness_application" "example" {
  name = "myapp"
}

resource "harness_environment" "qa" {
  name   = "qa"
  app_id = harness_application.example.id
  type   = "NON_PROD"
}

resource "harness_cloudprovider_kubernetes" "k8s" {
  name = "k8s"

  // Example of scoping to all non-prod environments of a specific application
  usage_scope {
    application_id          = harness_application.example.id
    environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
  }

  // Example of scoping to a specific environment
  usage_scope {
    application_id = harness_application.example.id
    environment_id = harness_environment.qa.id
  }
}
```

