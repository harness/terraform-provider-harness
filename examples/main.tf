terraform {
  required_providers {
    harness = {
      source  = "customers/pexa/harness"
      version = "0.1.0"
    }
  }
}

provider "harness" {
  platform_api_key = "pat.Ke-E1FX2SO2ZAL2TXqpLjg.674263f828e64616d95096a7.LSbySayRyWZVb7UyOcex" 
  account_id = "Ke-E1FX2SO2ZAL2TXqpLjg" 
  endpoint = "https://app.harness.io/gateway"
}

resource "harness_platform_connector_aws" "aws_connector" {


  org_id      = "default" 
  project_id  = "petclinic" 
  identifier = "HARNESS_AWS_CONNECTOR"
  name = "HARNESS_AWS_CONNECTOR"

   manual {
    access_key = "123" 
    secret_key_ref = "AWS_SECRET_KEY"
    execute_on_delegate = false	
}
}

