package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCreateProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"agent_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"generation": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"spec": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"clusterResourceWhitelist": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"kind": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"destinations": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"server": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"sourceRepos": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
s
func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) error {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_identifier").(string),
		d.Get("account_identifier").(string),
		d.Get("org_identifier").(string),
		d.Get("project_identifier").(string),
	)

	// Prepare project data from Terraform resource data
	projectData := d.Get("project").(map[string]interface{})
	projectJSON, err := json.Marshal(projectData)
	if err != nil {
		return fmt.Errorf("error marshalling project data: %s", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(projectJSON))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %s", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Optionally handle response data if needed

	// Set resource ID if creation was successful
	d.SetId("project-identifier") // Replace with a unique identifier if available

	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_identifier").(string),
		d.Get("query_name").(string),
		d.Get("account_identifier").(string),
		d.Get("org_identifier").(string),
		d.Get("project_identifier").(string),
	)

	// Create HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err)
	}

	// Set header
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %s", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Process response body if needed

	// Set resource ID if retrieval was successful
	d.SetId(fmt.Sprintf("%s/%s", d.Get("agent_identifier").(string), d.Get("query_name").(string)))

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_identifier").(string),
		d.Get("project.metadata.name").(string),
		d.Get("account_identifier").(string),
		d.Get("org_identifier").(string),
		d.Get("project_identifier").(string),
	)

	// Prepare project data from Terraform resource data
	projectData := d.Get("project").(map[string]interface{})
	projectJSON, err := json.Marshal(projectData)
	if err != nil {
		return fmt.Errorf("error marshalling project data: %s", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(projectJSON))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %s", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Optionally handle response data if needed

	// Set resource ID if creation was successful
	d.SetId("project-identifier") // Replace with a unique identifier if available

	return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_identifier").(string),
		d.Get("project.metadata.name").(string),
		d.Get("account_identifier").(string),
		d.Get("org_identifier").(string),
		d.Get("project_identifier").(string),
	)

	// Create HTTP DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %s", err)
	}

	// Set header
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %s", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Resource deleted successfully, so clear resource ID
	d.SetId("")

	return nil
}
