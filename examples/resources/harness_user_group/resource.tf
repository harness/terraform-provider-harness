resource "harness_user_group" "example" {
  name        = "example-group"
  description = "This group demonstrates account level and resource level permissions."

  permissions {

    account_permissions = ["ADMINISTER_OTHER_ACCOUNT_FUNCTIONS", "MANAGE_API_KEYS"]

    app_permissions {

      all {
        actions = ["CREATE", "READ", "UPDATE", "DELETE"]
      }

      deployment {
        actions = ["READ", "ROLLBACK_WORKFLOW", "EXECUTE_PIPELINE", "EXECUTE_WORKFLOW"]
        filters = ["NON_PRODUCTION_ENVIRONMENTS"]
      }

      deployment {
        actions = ["READ"]
        filters = ["PRODUCTION_ENVIRONMENTS"]
      }

      environment {
        actions = ["CREATE", "READ", "UPDATE", "DELETE"]
        filters = ["NON_PRODUCTION_ENVIRONMENTS"]
      }

      environment {
        actions = ["READ"]
        filters = ["PRODUCTION_ENVIRONMENTS"]
      }

      pipeline {
        actions = ["CREATE", "READ", "UPDATE", "DELETE"]
        filters = ["NON_PRODUCTION_PIPELINES"]
      }

      pipeline {
        actions = ["READ"]
        filters = ["PRODUCTION_PIPELINES"]
      }

      provisioner {
        actions = ["UPDATE", "DELETE"]
      }

      provisioner {
        actions = ["CREATE", "READ"]
      }

      service {
        actions = ["UPDATE", "DELETE"]
      }

      service {
        actions = ["UPDATE", "DELETE"]
      }

      template {
        actions = ["CREATE", "READ", "UPDATE", "DELETE"]
      }

      workflow {
        actions = ["UPDATE", "DELETE"]
        filters = ["NON_PRODUCTION_WORKFLOWS", ]
      }

      workflow {
        actions = ["CREATE", "READ"]
        filters = ["PRODUCTION_WORKFLOWS", "WORKFLOW_TEMPLATES"]
      }

    }
  }
}
