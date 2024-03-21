resource "harness_custom_secrets_manager" "example" {
	name                = "Custom Secrets Manager"
	description         = ""
	identifier          = "identifier"
  
	type                = "CustomSecretManager"
	on_delegate         = true
	timeout             = 20
	tags                = []
  
	template_ref        = "account.templateIdentifier" # doubt: this is an input correct?
	version_label       = "1"
  
	template_inputs = {
	  environment_variables = [
		{
		  name  = "var1_name" 
		  type  = "String"     
		  value = "var1_value" 
		},
		{
		  name  = "var2_name"  
		  type  = "String"  
		  value = "var2_value" 
		}
		# Add more variables as needed
	  ]
	}
  
  }
  
  ## useAsDefault: true, false
  
  resource "harness_custom_secrets_manager" "example" {
	name                = "Custom Secrets Manager"
	description         = ""
	identifier          = "identifier"
  
	type                = "CustomSecretManager"
	on_delegate         = false
	timeout             = 20
	tags                = []
  
	template_ref        = "account.templateIdentifier" # doubt: this is an input correct?
	version_label       = "1"
  
	template_inputs = {
	  environment_variables = [
		{
		  name  = "var1_name" 
		  type  = "String"     
		  value = "var1_value" 
		}
		# Add more variables as needed
	  ]
	}
	
	target_host = "target_host"
	ssh_secret_ref = "ssh_connection_secret_ref"
	working_directory = "working_directory_path"
  
  }