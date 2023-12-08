# Create Aws connector using Manual Auth

resource "harness_platform_connector_aws" "aws" {
  identifier  = "example_aws_connector"
  name        = "Example aws connector"
  description = "description of aws connector"
  tags        = ["foo:bar"]

  manual {
    access_key_ref     = "account.access_id"
    secret_key_ref     = "account.secret_id"
    delegate_selectors = ["harness-delegate"]
    region             = "aws_region"  
  }
}

# Create Aws connector using Manual Auth and Equal Jitter BackOff Strategy

resource "harness_platform_connector_aws" "aws" {
  identifier  = "example_aws_connector"
  name        = "Example aws connector"
  description = "description of aws connector"
  tags        = ["foo:bar"]

  manual {
    access_key_ref     = "account.access_id"
    secret_key_ref     = "account.secret_id"
    delegate_selectors = ["harness-delegate"]
    region             = "aws_region"
  }
  equal_jitter_backoff_strategy {
    base_delay       = 10
    max_backoff_time = 65
    retry_count      = 3
  }
}

# Create Aws connector using Manual Auth and Full Jitter BackOff Strategy

resource "harness_platform_connector_aws" "aws" {
  identifier  = "example_aws_connector"
  name        = "Example aws connector"
  description = "description of aws connector"
  tags        = ["foo:bar"]

  manual {
    access_key_ref     = "account.access_id"
    secret_key_ref     = "account.secret_id"
    delegate_selectors = ["harness-delegate"]
    region             = "aws_region"
  }
  full_jitter_backoff_strategy {
    base_delay       = 10
    max_backoff_time = 65
    retry_count      = 3
  }
}

# Create Aws connector using Manual Auth and Fixed Delay BackOff Strategy

resource "harness_platform_connector_aws" "aws" {
  identifier  = "example_aws_connector"
  name        = "Example aws connector"
  description = "description of aws connector"
  tags        = ["foo:bar"]

  manual {
    access_key_ref     = "account.access_id"
    secret_key_ref     = "account.secret_id"
    delegate_selectors = ["harness-delegate"]
    region             = "aws_region"
  }
  fixed_delay_backoff_strategy {
    fixed_backoff = 10
    retry_count   = 3
  }
}
