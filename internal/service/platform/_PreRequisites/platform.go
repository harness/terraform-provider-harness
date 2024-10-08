package _PreRequisites

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecretText() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretTextCreate,
		Read:   resourceSecretTextRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_manager_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSecretTextCreate(d *schema.ResourceData, m interface{}) error {
	resourceName := "harness_platform_secret_text.resource.tf"
	identifier := d.Get("identifier").(string)

	// Mock logic for creating a secret text
	fmt.Printf("Creating secret with identifier: %s\n", identifier)

	d.SetId(identifier) // Set the ID to the identifier
	return resourceSecretTextRead(d, m)
}

func resourceSecretTextRead(d *schema.ResourceData, m interface{}) error {
	identifier := d.Id()

	// Mock logic for reading a secret text
	fmt.Printf("Reading secret with identifier: %s\n", identifier)

	// Simulate reading the resource by setting some data
	d.Set("identifier", identifier)
	d.Set("name", "Test Secret")
	d.Set("secret_manager_identifier", "harnessSecretManager")
	d.Set("value_type", "Inline")
	d.Set("value", "super-secret-value") // Sensitive data

	return nil
}

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	identifier := d.Get("identifier").(string)
	fmt.Printf("Creating project with identifier: %s\n", identifier)
	d.SetId(identifier)
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	identifier := d.Id()
	fmt.Printf("Reading project with identifier: %s\n", identifier)

	// Simulate reading the resource by setting some data
	d.Set("identifier", identifier)
	d.Set("name", "Test Project")
	d.Set("org_id", "default")
	d.Set("color", "#0063F7")

	return nil
}

func resourceAzureKeyVaultConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureKeyVaultConnectorCreate,
		Read:   resourceAzureKeyVaultConnectorRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vault_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"azure_environment_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAzureKeyVaultConnectorCreate(d *schema.ResourceData, m interface{}) error {
	identifier := d.Get("identifier").(string)
	fmt.Printf("Creating Azure Key Vault connector with identifier: %s\n", identifier)
	d.SetId(identifier)
	return resourceAzureKeyVaultConnectorRead(d, m)
}

func resourceAzureKeyVaultConnectorRead(d *schema.ResourceData, m interface{}) error {
	identifier := d.Id()
	fmt.Printf("Reading Azure Key Vault connector with identifier: %s\n", identifier)

	// Simulate reading the resource by setting some data
	d.Set("identifier", identifier)
	d.Set("name", "Test Azure Connector")
	d.Set("client_id", "test-client-id")

	return nil
}

func resourceAwsSecretManagerConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsSecretManagerConnectorCreate,
		Read:   resourceAwsSecretManagerConnectorRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_key_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAwsSecretManagerConnectorCreate(d *schema.ResourceData, m interface{}) error {
	identifier := d.Get("identifier").(string)
	fmt.Printf("Creating AWS Secret Manager connector with identifier: %s\n", identifier)
	d.SetId(identifier)
	return resourceAwsSecretManagerConnectorRead(d, m)
}

func resourceAwsSecretManagerConnectorRead(d *schema.ResourceData, m interface{}) error {
	identifier := d.Id()
	fmt.Printf("Reading AWS Secret Manager connector with identifier: %s\n", identifier)

	// Simulate reading the resource by setting some data
	d.Set("identifier", identifier)
	d.Set("name", "Test AWS Connector")
	d.Set("region", "us-east-1")

	return nil
}
