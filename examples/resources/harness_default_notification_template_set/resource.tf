resource "harness_default_notification_template_set" "default_set" {
  name        = "Default Template Set"
  identifier  = "default_template_set"
  description = "Default notification templates for all events"
  org_id      = "my_org"
  project_id  = "my_project"

  # Email template configuration
  email_templates {
    type = "PIPELINE_START"
    subject = "[${var.env}] Pipeline Started: ${pipeline.name}"
    body = <<-EOT
    <!DOCTYPE html>
    <html>
    <body>
      <h2>Pipeline Started</h2>
      <p><strong>Pipeline:</strong> ${pipeline.name}</p>
      <p><strong>Triggered by:</strong> ${pipeline.triggered_by}</p>
      <p><strong>Start Time:</strong> ${pipeline.start_time}</p>
      <p>View details: <a href="${pipeline.execution_url}">${pipeline.execution_url}</a></p>
    </body>
    </html>
    EOT
  }

  # Slack template configuration
  slack_templates {
    type = "PIPELINE_SUCCESS"
    message = "✅ *${pipeline.name}* succeeded!\n*Environment*: ${env.name}\n*Duration*: ${pipeline.duration}\n<${pipeline.execution_url}|View Details>"
  }

  # Default settings
  is_default = true

  tags = {
    environment = "production"
    team        = "devops"
  }
}

# Output the created template set details
output "created_template_set" {
  value = harness_default_notification_template_set.default_set
}
