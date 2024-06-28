package project

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:     schema.TypeList,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_resource_whitelist": {
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
									"source_repos": {
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

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Read API key from provider configuration or environment variables
	apiKey := "" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://qa.harness.io/gateway/gitops/api/v1/agents/%s/projects?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_id").(string),
		d.Get("account_id").(string),
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)

	// Prepare project data from Terraform resource data
	projectData := d.Get("project").(map[string]interface{})
	// Type assertion to slice of interface{}
	if sliceData, ok := d.Get("project").([]interface{}); ok {
		// Iterate over the slice elements
		for _, item := range sliceData {
			fmt.Println(item)
		}
	} else {
		fmt.Println("Data is not a slice of interface{}")
	}
	projectJSON, err := json.Marshal(projectData)
	if err != nil {
		return nil
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(projectJSON))
	if err != nil {
		return nil
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil
	}

	// Optionally handle response data if needed

	// Set resource ID if creation was successful
	d.SetId("project-identifier") // Replace with a unique identifier if available

	return diag.Diagnostics{}
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_idr").(string),
		d.Get("query_name").(string),
		d.Get("account_id").(string),
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)

	// Create HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	// Set header
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil
	}

	// Process response body if needed

	// Set resource ID if retrieval was successful
	d.SetId(fmt.Sprintf("%s/%s", d.Get("agent_identifier").(string), d.Get("query_name").(string)))

	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_id").(string),
		d.Get("project.metadata.name").(string),
		d.Get("account_id").(string),
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)

	// Prepare project data from Terraform resource data
	projectData := d.Get("project").(map[string]interface{})
	projectJSON, err := json.Marshal(projectData)
	if err != nil {
		return nil
	}

	// Create HTTP POST request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(projectJSON))
	if err != nil {
		return nil
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil
	}

	// Optionally handle response data if needed

	// Set resource ID if creation was successful
	d.SetId("project-identifier") // Replace with a unique identifier if available

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Read API key from provider configuration or environment variables
	apiKey := "YOUR_API_KEY_HERE" // Replace with your actual API key

	// Construct the URL with placeholders replaced by actual values from Terraform resource data
	url := fmt.Sprintf("https://app.harness.io/gitops/api/v1/agents/%s/projects/%s?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
		d.Get("agent_idr").(string),
		d.Get("project.metadata.name").(string),
		d.Get("account_id").(string),
		d.Get("org_id").(string),
		d.Get("project_id").(string),
	)

	// Create HTTP DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil
	}

	// Set header
	req.Header.Set("x-api-key", apiKey)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil
	}

	// Resource deleted successfully, so clear resource ID
	d.SetId("")

	return nil
}

type ProjectMetadata struct {
	Generation int    `json:"generation"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type ProjectSpec struct {
	ClusterResourceWhitelist []struct {
		Group string `json:"group"`
		Kind  string `json:"kind"`
	} `json:"cluster_resource_whitelist"`

	Destinations []struct {
		Namespace string `json:"namespace"`
		Server    string `json:"server"`
	} `json:"destinations"`

	SourceRepos []string `json:"source_repos"`
}

type Project struct {
	Metadata ProjectMetadata `json:"metadata"`
	Spec     ProjectSpec     `json:"spec"`
}
