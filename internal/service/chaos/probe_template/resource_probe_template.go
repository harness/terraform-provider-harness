package probe_template

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProbeTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Harness Chaos Probe Templates.",

		CreateContext: resourceProbeTemplateCreate,
		ReadContext:   resourceProbeTemplateRead,
		UpdateContext: resourceProbeTemplateUpdate,
		DeleteContext: resourceProbeTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceProbeTemplateImport,
		},

		Schema: resourceProbeTemplateSchema(),
	}
}

func resourceProbeTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)
	identity := d.Get("identity").(string)

	// Build the request
	req := chaos.ChaosprobetemplateProbeTemplate{
		Identity:    identity,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		HubRef:      hubIdentity,
	}

	// Set probe type
	if v, ok := d.GetOk("type"); ok {
		probeType := chaos.ProbeProbeType(v.(string))
		req.Type_ = &probeType
	}

	// Set infrastructure type
	if v, ok := d.GetOk("infrastructure_type"); ok {
		infraType := chaos.ProbeInfrastructureType(v.(string))
		req.InfrastructureType = &infraType
	}

	// Set tags
	if v, ok := d.GetOk("tags"); ok {
		tags := v.([]interface{})
		tagList := make([]string, len(tags))
		for i, tag := range tags {
			tagList[i] = tag.(string)
		}
		req.Tags = tagList
	}

	// Build probe properties (simplified)
	if err := buildProbePropertiesSimplified(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build run properties
	if err := buildRunPropertiesSimplified(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build variables
	buildVariablesSimplified(d, &req)

	log.Printf("[DEBUG] Creating probe template: %s in hub: %s", identity, hubIdentity)

	// Create the probe template
	resp, httpResp, err := c.DefaultApi.CreateProbeTemplate(ctx, req, accountID, orgID, projectID)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	// Set the ID
	d.SetId(generateID(accountID, orgID, projectID, hubIdentity, identity))

	// Set computed fields
	d.Set("account_id", accountID)
	if resp.Revision != 0 {
		d.Set("revision", resp.Revision)
	}
	d.Set("is_default", resp.IsDefault)
	d.Set("hub_ref", resp.HubRef)

	log.Printf("[DEBUG] Created probe template: %s", identity)

	return resourceProbeTemplateRead(ctx, d, meta)
}

func resourceProbeTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	// Parse the ID
	accountID, orgID, projectID, hubIdentity, identity, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Reading probe template: %s from hub: %s", identity, hubIdentity)

	// Get the probe template
	resp, httpResp, err := c.DefaultApi.GetProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, identity, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Probe template not found, removing from state: %s", identity)
			d.SetId("")
			return nil
		}
		return helpers.HandleChaosReadApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		log.Printf("[WARN] Probe template data is nil, removing from state: %s", identity)
		d.SetId("")
		return nil
	}

	// Set the resource data
	return setProbeTemplateDataSimplified(d, resp.Data, accountID, orgID, projectID, hubIdentity)
}

func resourceProbeTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubIdentity := d.Get("hub_identity").(string)
	identity := d.Get("identity").(string)

	// Build the update request
	req := chaos.ChaosprobetemplateProbeTemplate{
		Identity:    identity,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		HubRef:      hubIdentity,
	}

	// Set probe type
	if v, ok := d.GetOk("type"); ok {
		probeType := chaos.ProbeProbeType(v.(string))
		req.Type_ = &probeType
	}

	// Set infrastructure type
	if v, ok := d.GetOk("infrastructure_type"); ok {
		infraType := chaos.ProbeInfrastructureType(v.(string))
		req.InfrastructureType = &infraType
	}

	// Set tags
	if v, ok := d.GetOk("tags"); ok {
		tags := v.([]interface{})
		tagList := make([]string, len(tags))
		for i, tag := range tags {
			tagList[i] = tag.(string)
		}
		req.Tags = tagList
	}

	// Build probe properties
	if err := buildProbePropertiesSimplified(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build run properties
	if err := buildRunPropertiesSimplified(d, &req); err != nil {
		return diag.FromErr(err)
	}

	// Build variables
	buildVariablesSimplified(d, &req)

	log.Printf("[DEBUG] Updating probe template: %s", identity)

	// Update the probe template
	_, httpResp, err := c.DefaultApi.UpdateProbeTemplate(ctx, req, accountID, orgID, projectID, identity, nil)
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Updated probe template: %s", identity)

	return resourceProbeTemplateRead(ctx, d, meta)
}

func resourceProbeTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	// Parse the ID
	accountID, orgID, projectID, hubIdentity, identity, err := parseID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Deleting probe template: %s from hub: %s", identity, hubIdentity)

	// Delete the probe template
	_, httpResp, err := c.DefaultApi.DeleteProbeTemplate(ctx, accountID, orgID, projectID, hubIdentity, identity, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			log.Printf("[WARN] Probe template not found during delete: %s", identity)
			return nil
		}
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	log.Printf("[DEBUG] Deleted probe template: %s", identity)

	return nil
}

func resourceProbeTemplateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")

	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	accountID := c.AccountId
	if accountID == "" {
		return nil, fmt.Errorf("account ID must be configured in the provider")
	}

	var orgID, projectID, hubIdentity, identity string

	switch len(parts) {
	case 4:
		orgID = parts[0]
		projectID = parts[1]
		hubIdentity = parts[2]
		identity = parts[3]
	case 3:
		orgID = parts[0]
		hubIdentity = parts[1]
		identity = parts[2]
	case 2:
		hubIdentity = parts[0]
		identity = parts[1]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected: org_id/project_id/hub_identity/identity or org_id/hub_identity/identity or hub_identity/identity, got: %s", d.Id())
	}

	d.SetId(generateID(accountID, orgID, projectID, hubIdentity, identity))

	if orgID != "" {
		d.Set("org_id", orgID)
	}
	if projectID != "" {
		d.Set("project_id", projectID)
	}
	d.Set("hub_identity", hubIdentity)
	d.Set("identity", identity)

	log.Printf("[DEBUG] Importing probe template: %s from hub: %s", identity, hubIdentity)

	diags := resourceProbeTemplateRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read probe template during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

// Helper functions

func generateID(accountID, orgID, projectID, hubIdentity, identity string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", accountID, orgID, projectID, hubIdentity, identity)
}

func parseID(id string) (accountID, orgID, projectID, hubIdentity, identity string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 5 {
		return "", "", "", "", "", fmt.Errorf("invalid ID format: %s", id)
	}
	return parts[0], parts[1], parts[2], parts[3], parts[4], nil
}

func setProbeTemplateDataSimplified(d *schema.ResourceData, template *chaos.GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosprobetemplateChaosProbeTemplate, accountID, orgID, projectID, hubIdentity string) diag.Diagnostics {
	d.Set("account_id", accountID)
	d.Set("identity", template.Identity)
	d.Set("name", template.Name)
	d.Set("hub_identity", hubIdentity)

	if orgID != "" {
		d.Set("org_id", orgID)
	}
	if projectID != "" {
		d.Set("project_id", projectID)
	}

	if template.Description != "" {
		d.Set("description", template.Description)
	}
	if len(template.Tags) > 0 {
		d.Set("tags", template.Tags)
	}
	if template.Type_ != nil {
		d.Set("type", string(*template.Type_))
	}
	if template.InfrastructureType != nil {
		d.Set("infrastructure_type", string(*template.InfrastructureType))
	}

	d.Set("revision", template.Revision)
	d.Set("is_default", template.IsDefault)
	d.Set("hub_ref", template.HubRef)

	// Parse probe properties
	if template.ProbeProperties != nil {
		if template.ProbeProperties.HttpProbe != nil {
			httpBlock := map[string]interface{}{
				"url": template.ProbeProperties.HttpProbe.Url,
			}

			// Parse method object
			if template.ProbeProperties.HttpProbe.Method != nil {
				methodBlock := map[string]interface{}{}

				// Parse GET method
				if template.ProbeProperties.HttpProbe.Method.Get != nil {
					getBlock := map[string]interface{}{}
					if template.ProbeProperties.HttpProbe.Method.Get.Criteria != "" {
						getBlock["criteria"] = template.ProbeProperties.HttpProbe.Method.Get.Criteria
					}
					if template.ProbeProperties.HttpProbe.Method.Get.ResponseBody != "" {
						getBlock["response_body"] = template.ProbeProperties.HttpProbe.Method.Get.ResponseBody
					}
					if template.ProbeProperties.HttpProbe.Method.Get.ResponseCode != "" {
						getBlock["response_code"] = template.ProbeProperties.HttpProbe.Method.Get.ResponseCode
					}
					if len(getBlock) > 0 {
						methodBlock["get"] = []map[string]interface{}{getBlock}
					}
				}

				// Parse POST method
				if template.ProbeProperties.HttpProbe.Method.Post != nil {
					postBlock := map[string]interface{}{}
					if template.ProbeProperties.HttpProbe.Method.Post.Criteria != "" {
						postBlock["criteria"] = template.ProbeProperties.HttpProbe.Method.Post.Criteria
					}
					if template.ProbeProperties.HttpProbe.Method.Post.ResponseBody != "" {
						postBlock["response_body"] = template.ProbeProperties.HttpProbe.Method.Post.ResponseBody
					}
					if template.ProbeProperties.HttpProbe.Method.Post.ResponseCode != "" {
						postBlock["response_code"] = template.ProbeProperties.HttpProbe.Method.Post.ResponseCode
					}
					if template.ProbeProperties.HttpProbe.Method.Post.Body != "" {
						postBlock["body"] = template.ProbeProperties.HttpProbe.Method.Post.Body
					}
					if template.ProbeProperties.HttpProbe.Method.Post.BodyPath != "" {
						postBlock["body_path"] = template.ProbeProperties.HttpProbe.Method.Post.BodyPath
					}
					if template.ProbeProperties.HttpProbe.Method.Post.ContentType != "" {
						postBlock["content_type"] = template.ProbeProperties.HttpProbe.Method.Post.ContentType
					}
					if len(postBlock) > 0 {
						methodBlock["post"] = []map[string]interface{}{postBlock}
					}
				}

				if len(methodBlock) > 0 {
					httpBlock["method"] = []map[string]interface{}{methodBlock}
				}
			}

			// TODO: Parse headers, auth, tls_config

			d.Set("http_probe", []map[string]interface{}{httpBlock})
		}

		if template.ProbeProperties.CmdProbe != nil {
			cmdBlock := map[string]interface{}{
				"command": template.ProbeProperties.CmdProbe.Command,
			}
			if template.ProbeProperties.CmdProbe.Source != "" {
				cmdBlock["source"] = template.ProbeProperties.CmdProbe.Source
			}

			// Parse comparator
			if template.ProbeProperties.CmdProbe.Comparator != nil {
				comparatorBlock := map[string]interface{}{
					"type":     template.ProbeProperties.CmdProbe.Comparator.Type_,
					"criteria": template.ProbeProperties.CmdProbe.Comparator.Criteria,
					"value":    template.ProbeProperties.CmdProbe.Comparator.Value,
				}
				cmdBlock["comparator"] = []map[string]interface{}{comparatorBlock}
			}

			// Parse env variables
			if len(template.ProbeProperties.CmdProbe.Env) > 0 {
				envList := make([]map[string]interface{}, len(template.ProbeProperties.CmdProbe.Env))
				for i, env := range template.ProbeProperties.CmdProbe.Env {
					envList[i] = map[string]interface{}{
						"name":  env.Name,
						"value": env.Value,
					}
				}
				cmdBlock["env"] = envList
			}

			d.Set("cmd_probe", []map[string]interface{}{cmdBlock})
		}

		if template.ProbeProperties.K8sProbe != nil {
			k8sBlock := map[string]interface{}{
				"version":  template.ProbeProperties.K8sProbe.Version,
				"resource": template.ProbeProperties.K8sProbe.Resource,
			}
			if template.ProbeProperties.K8sProbe.Group != "" {
				k8sBlock["group"] = template.ProbeProperties.K8sProbe.Group
			}
			if template.ProbeProperties.K8sProbe.Namespace != "" {
				k8sBlock["namespace"] = template.ProbeProperties.K8sProbe.Namespace
			}
			if template.ProbeProperties.K8sProbe.FieldSelector != "" {
				k8sBlock["field_selector"] = template.ProbeProperties.K8sProbe.FieldSelector
			}
			if template.ProbeProperties.K8sProbe.LabelSelector != "" {
				k8sBlock["label_selector"] = template.ProbeProperties.K8sProbe.LabelSelector
			}
			if template.ProbeProperties.K8sProbe.ResourceNames != "" {
				k8sBlock["resource_names"] = template.ProbeProperties.K8sProbe.ResourceNames
			}
			if template.ProbeProperties.K8sProbe.Operation != "" {
				k8sBlock["operation"] = template.ProbeProperties.K8sProbe.Operation
			}
			d.Set("k8s_probe", []map[string]interface{}{k8sBlock})
		}

		if template.ProbeProperties.ApmProbe != nil {
			apmBlock := map[string]interface{}{}

			// Set APM type
			if template.ProbeProperties.ApmProbe.Type_ != nil {
				apmBlock["apm_type"] = string(*template.ProbeProperties.ApmProbe.Type_)
			}

			// Parse comparator
			if template.ProbeProperties.ApmProbe.Comparator != nil {
				comparatorBlock := map[string]interface{}{
					"type":     template.ProbeProperties.ApmProbe.Comparator.Type_,
					"criteria": template.ProbeProperties.ApmProbe.Comparator.Criteria,
					"value":    template.ProbeProperties.ApmProbe.Comparator.Value,
				}
				apmBlock["comparator"] = []map[string]interface{}{comparatorBlock}
			}

			// Parse Prometheus inputs
			if template.ProbeProperties.ApmProbe.PrometheusProbeInputs != nil {
				promBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.PrometheusProbeInputs.ConnectorID,
					"query":        template.ProbeProperties.ApmProbe.PrometheusProbeInputs.Query,
				}
				// Parse TLS config
				if template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig != nil {
					tlsBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.CaCrt != nil {
						tlsBlock["ca_cert_secret"] = template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.CaCrt.Identifier
					}
					if template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.ClientCrt != nil {
						tlsBlock["client_cert_secret"] = template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.ClientCrt.Identifier
					}
					if template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.Key != nil {
						tlsBlock["client_key_secret"] = template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.Key.Identifier
					}
					if template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.InsecureSkipVerify != nil {
						tlsBlock["insecure_skip_verify"] = *template.ProbeProperties.ApmProbe.PrometheusProbeInputs.TlsConfig.InsecureSkipVerify
					}
					if len(tlsBlock) > 0 {
						promBlock["tls_config"] = []map[string]interface{}{tlsBlock}
					}
				}
				apmBlock["prometheus_inputs"] = []map[string]interface{}{promBlock}
			}

			// Parse Datadog inputs
			if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs != nil {
				datadogBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.ConnectorID,
				}
				if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.Query != "" {
					datadogBlock["query"] = template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.Query
				}
				if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.DurationInMin != nil {
					if durationVal, ok := (*template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.DurationInMin).(float64); ok {
						datadogBlock["duration_in_min"] = int(durationVal)
					}
				}
				// Parse synthetics test
				if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.SyntheticsTest != nil {
					syntheticsBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.SyntheticsTest.PublicId != "" {
						syntheticsBlock["public_id"] = template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.SyntheticsTest.PublicId
					}
					if template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.SyntheticsTest.TestType != nil {
						syntheticsBlock["test_type"] = string(*template.ProbeProperties.ApmProbe.DatadogApmProbeInputs.SyntheticsTest.TestType)
					}
					if len(syntheticsBlock) > 0 {
						datadogBlock["synthetics_test"] = []map[string]interface{}{syntheticsBlock}
					}
				}
				apmBlock["datadog_inputs"] = []map[string]interface{}{datadogBlock}
			}

			// Parse Dynatrace inputs
			if template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs != nil {
				dynatraceBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.ConnectorID,
				}
				if template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.DurationInMin != nil {
					if durationVal, ok := (*template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.DurationInMin).(float64); ok {
						dynatraceBlock["duration_in_min"] = int(durationVal)
					}
				}
				// Parse metrics
				if template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.Metrics != nil {
					metricsBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.Metrics.EntitySelector != "" {
						metricsBlock["entity_selector"] = template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.Metrics.EntitySelector
					}
					if template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.Metrics.MetricsSelector != "" {
						metricsBlock["metrics_selector"] = template.ProbeProperties.ApmProbe.DynatraceApmProbeInputs.Metrics.MetricsSelector
					}
					if len(metricsBlock) > 0 {
						dynatraceBlock["metrics"] = []map[string]interface{}{metricsBlock}
					}
				}
				apmBlock["dynatrace_inputs"] = []map[string]interface{}{dynatraceBlock}
			}

			// Parse AppDynamics inputs
			if template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs != nil {
				appDynamicsBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.ConnectorID,
				}
				// Parse appd_metrics
				if template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics != nil {
					appdMetricsBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.ApplicationName != "" {
						appdMetricsBlock["application_name"] = template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.ApplicationName
					}
					if template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.MetricsFullPath != "" {
						appdMetricsBlock["metrics_full_path"] = template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.MetricsFullPath
					}
					if template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.DurationInMin != nil {
						if durationVal, ok := (*template.ProbeProperties.ApmProbe.AppDynamicsProbeInputs.AppdMetrics.DurationInMin).(float64); ok {
							appdMetricsBlock["duration_in_min"] = int(durationVal)
						}
					}
					if len(appdMetricsBlock) > 0 {
						appDynamicsBlock["appd_metrics"] = []map[string]interface{}{appdMetricsBlock}
					}
				}
				apmBlock["app_dynamics_inputs"] = []map[string]interface{}{appDynamicsBlock}
			}

			// Parse NewRelic inputs
			if template.ProbeProperties.ApmProbe.NewRelicProbeInputs != nil {
				newRelicBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.NewRelicProbeInputs.ConnectorID,
				}
				// Parse new_relic_metric
				if template.ProbeProperties.ApmProbe.NewRelicProbeInputs.NewRelicMetric != nil {
					newRelicMetricBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.NewRelicProbeInputs.NewRelicMetric.Query != "" {
						newRelicMetricBlock["query"] = template.ProbeProperties.ApmProbe.NewRelicProbeInputs.NewRelicMetric.Query
					}
					if template.ProbeProperties.ApmProbe.NewRelicProbeInputs.NewRelicMetric.QueryMetric != "" {
						newRelicMetricBlock["query_metric"] = template.ProbeProperties.ApmProbe.NewRelicProbeInputs.NewRelicMetric.QueryMetric
					}
					if len(newRelicMetricBlock) > 0 {
						newRelicBlock["new_relic_metric"] = []map[string]interface{}{newRelicMetricBlock}
					}
				}
				apmBlock["new_relic_inputs"] = []map[string]interface{}{newRelicBlock}
			}

			// Parse Splunk Observability inputs
			if template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs != nil {
				splunkBlock := map[string]interface{}{
					"connector_id": template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.ConnectorID,
				}
				// Parse splunk_observability_metrics
				if template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics != nil {
					splunkMetricsBlock := map[string]interface{}{}
					if template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.Query != "" {
						splunkMetricsBlock["query"] = template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.Query
					}
					if template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.DurationInMin != nil {
						if durationVal, ok := (*template.ProbeProperties.ApmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.DurationInMin).(float64); ok {
							splunkMetricsBlock["duration_in_min"] = int(durationVal)
						}
					}
					if len(splunkMetricsBlock) > 0 {
						splunkBlock["splunk_observability_metrics"] = []map[string]interface{}{splunkMetricsBlock}
					}
				}
				apmBlock["splunk_observability_inputs"] = []map[string]interface{}{splunkBlock}
			}

			// Parse GCP Cloud Monitoring inputs
			if template.ProbeProperties.ApmProbe.GcpCloudMonitoringProbeInputs != nil {
				gcpBlock := map[string]interface{}{
					"project_id":          template.ProbeProperties.ApmProbe.GcpCloudMonitoringProbeInputs.ProjectID,
					"query":               template.ProbeProperties.ApmProbe.GcpCloudMonitoringProbeInputs.Query,
					"service_account_key": template.ProbeProperties.ApmProbe.GcpCloudMonitoringProbeInputs.ServiceAccountKey,
				}
				apmBlock["gcp_cloud_monitoring_inputs"] = []map[string]interface{}{gcpBlock}
			}

			d.Set("apm_probe", []map[string]interface{}{apmBlock})
		}
	}

	// Parse run properties (simplified)
	if template.RunProperties != nil {
		runPropsBlock := map[string]interface{}{}

		// String fields - only set if non-empty
		if template.RunProperties.InitialDelay != "" {
			runPropsBlock["initial_delay"] = template.RunProperties.InitialDelay
		}
		if template.RunProperties.Interval != "" {
			runPropsBlock["interval"] = template.RunProperties.Interval
		}
		if template.RunProperties.Timeout != "" {
			runPropsBlock["timeout"] = template.RunProperties.Timeout
		}
		if template.RunProperties.PollingInterval != "" {
			runPropsBlock["polling_interval"] = template.RunProperties.PollingInterval
		}
		if template.RunProperties.Verbosity != "" {
			runPropsBlock["verbosity"] = template.RunProperties.Verbosity
		}

		// Boolean field - only set if true (non-default)
		// Following Terraform best practice: don't set default values to avoid drift
		if template.RunProperties.StopOnFailure {
			runPropsBlock["stop_on_failure"] = true
		}

		// Integer fields - only set if non-zero (non-default)
		if template.RunProperties.Attempt != nil {
			attemptVal := getIntFromInterface(template.RunProperties.Attempt)
			if attemptVal > 0 {
				runPropsBlock["attempt"] = attemptVal
			}
		}
		if template.RunProperties.Retry != nil {
			retryVal := getIntFromInterface(template.RunProperties.Retry)
			if retryVal > 0 {
				runPropsBlock["retry"] = retryVal
			}
		}

		// Only set run_properties if there are non-default values
		if len(runPropsBlock) > 0 {
			d.Set("run_properties", []map[string]interface{}{runPropsBlock})
		}
	}

	// Parse variables
	if len(template.Variables) > 0 {
		varsList := make([]map[string]interface{}, len(template.Variables))
		for i, v := range template.Variables {
			varMap := map[string]interface{}{
				"name": v.Name,
			}
			if v.Value != nil {
				varMap["value"] = *v.Value
			}
			if v.Description != "" {
				varMap["description"] = v.Description
			}
			if v.Type_ != nil {
				varMap["type"] = strings.ToLower(string(*v.Type_))
			}
			// Only set required if true (non-default)
			// Following Terraform best practice: don't set default values to avoid drift
			if v.Required {
				varMap["required"] = true
			}
			varsList[i] = varMap
		}
		d.Set("variables", varsList)
	}

	return nil
}

func buildProbePropertiesSimplified(d *schema.ResourceData, req *chaos.ChaosprobetemplateProbeTemplate) error {
	probeType := d.Get("type").(string)

	switch probeType {
	case "httpProbe":
		if v, ok := d.GetOk("http_probe"); ok && len(v.([]interface{})) > 0 {
			httpConfig := v.([]interface{})[0].(map[string]interface{})
			httpProbe := &chaos.ProbeHttpProbeTemplate{
				Url: httpConfig["url"].(string),
			}

			// Build method object
			if methodList, ok := httpConfig["method"].([]interface{}); ok && len(methodList) > 0 {
				methodConfig := methodList[0].(map[string]interface{})
				method := &chaos.ProbeMethodTemplate{}

				// GET method
				if getList, ok := methodConfig["get"].([]interface{}); ok && len(getList) > 0 {
					getConfig := getList[0].(map[string]interface{})
					method.Get = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeGet{}
					if criteria, ok := getConfig["criteria"].(string); ok && criteria != "" {
						method.Get.Criteria = criteria
					}
					if responseBody, ok := getConfig["response_body"].(string); ok && responseBody != "" {
						method.Get.ResponseBody = responseBody
					}
					if responseCode, ok := getConfig["response_code"].(string); ok && responseCode != "" {
						method.Get.ResponseCode = responseCode
					}
				}

				// POST method
				if postList, ok := methodConfig["post"].([]interface{}); ok && len(postList) > 0 {
					postConfig := postList[0].(map[string]interface{})
					method.Post = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbePost{}
					if criteria, ok := postConfig["criteria"].(string); ok && criteria != "" {
						method.Post.Criteria = criteria
					}
					if responseBody, ok := postConfig["response_body"].(string); ok && responseBody != "" {
						method.Post.ResponseBody = responseBody
					}
					if responseCode, ok := postConfig["response_code"].(string); ok && responseCode != "" {
						method.Post.ResponseCode = responseCode
					}
					if body, ok := postConfig["body"].(string); ok && body != "" {
						method.Post.Body = body
					}
					if bodyPath, ok := postConfig["body_path"].(string); ok && bodyPath != "" {
						method.Post.BodyPath = bodyPath
					}
					if contentType, ok := postConfig["content_type"].(string); ok && contentType != "" {
						method.Post.ContentType = contentType
					}
				}

				httpProbe.Method = method
			} else {
				// Default to GET method with empty response code if no method specified
				httpProbe.Method = &chaos.ProbeMethodTemplate{
					Get: &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeGet{
						ResponseCode: "",
					},
				}
			}

			// TODO: Add headers, auth, tls_config support

			req.ProbeProperties = &chaos.ProbeProbeTemplateProperties{
				HttpProbe: httpProbe,
			}
		} else {
			return fmt.Errorf("http_probe block is required when type is 'httpProbe'")
		}

	case "cmdProbe":
		if v, ok := d.GetOk("cmd_probe"); ok && len(v.([]interface{})) > 0 {
			cmdConfig := v.([]interface{})[0].(map[string]interface{})
			cmdProbe := &chaos.ProbeCmdProbeTemplate{
				Command: cmdConfig["command"].(string),
			}
			if source, ok := cmdConfig["source"].(string); ok && source != "" {
				cmdProbe.Source = source
			}

			// Build comparator
			if comparatorList, ok := cmdConfig["comparator"].([]interface{}); ok && len(comparatorList) > 0 {
				comparatorConfig := comparatorList[0].(map[string]interface{})
				cmdProbe.Comparator = &chaos.ProbeComparatorTemplate{
					Type_:    comparatorConfig["type"].(string),
					Criteria: comparatorConfig["criteria"].(string),
					Value:    comparatorConfig["value"].(string),
				}
			}

			// Build env variables
			if envList, ok := cmdConfig["env"].([]interface{}); ok && len(envList) > 0 {
				envVars := make([]chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeEnv, len(envList))
				for i, envItem := range envList {
					envConfig := envItem.(map[string]interface{})
					envVars[i] = chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeEnv{
						Name:  envConfig["name"].(string),
						Value: envConfig["value"].(string),
					}
				}
				cmdProbe.Env = envVars
			}

			req.ProbeProperties = &chaos.ProbeProbeTemplateProperties{
				CmdProbe: cmdProbe,
			}
		} else {
			return fmt.Errorf("cmd_probe block is required when type is 'cmdProbe'")
		}

	case "k8sProbe":
		if v, ok := d.GetOk("k8s_probe"); ok && len(v.([]interface{})) > 0 {
			k8sConfig := v.([]interface{})[0].(map[string]interface{})
			k8sProbe := &chaos.ProbeK8SProbeTemplate{
				Version:  k8sConfig["version"].(string),
				Resource: k8sConfig["resource"].(string),
			}
			if group, ok := k8sConfig["group"].(string); ok && group != "" {
				k8sProbe.Group = group
			}
			if namespace, ok := k8sConfig["namespace"].(string); ok && namespace != "" {
				k8sProbe.Namespace = namespace
			}
			if fieldSelector, ok := k8sConfig["field_selector"].(string); ok && fieldSelector != "" {
				k8sProbe.FieldSelector = fieldSelector
			}
			if labelSelector, ok := k8sConfig["label_selector"].(string); ok && labelSelector != "" {
				k8sProbe.LabelSelector = labelSelector
			}
			if resourceNames, ok := k8sConfig["resource_names"].(string); ok && resourceNames != "" {
				k8sProbe.ResourceNames = resourceNames
			}
			if operation, ok := k8sConfig["operation"].(string); ok && operation != "" {
				k8sProbe.Operation = operation
			}
			req.ProbeProperties = &chaos.ProbeProbeTemplateProperties{
				K8sProbe: k8sProbe,
			}
		} else {
			return fmt.Errorf("k8s_probe block is required when type is 'k8sProbe'")
		}

	case "apmProbe":
		if v, ok := d.GetOk("apm_probe"); ok && len(v.([]interface{})) > 0 {
			apmConfig := v.([]interface{})[0].(map[string]interface{})
			apmProbe := &chaos.ProbeApmProbeTemplate{}

			// Set APM type
			if apmType, ok := apmConfig["apm_type"].(string); ok && apmType != "" {
				apmTypeEnum := chaos.ProbeApmProbeType(apmType)
				apmProbe.Type_ = &apmTypeEnum
			}

			// Build comparator
			if comparatorList, ok := apmConfig["comparator"].([]interface{}); ok && len(comparatorList) > 0 {
				comparatorConfig := comparatorList[0].(map[string]interface{})
				apmProbe.Comparator = &chaos.ProbeComparatorTemplate{
					Type_:    comparatorConfig["type"].(string),
					Criteria: comparatorConfig["criteria"].(string),
					Value:    comparatorConfig["value"].(string),
				}
			}

			// Build Prometheus inputs
			if promList, ok := apmConfig["prometheus_inputs"].([]interface{}); ok && len(promList) > 0 {
				promConfig := promList[0].(map[string]interface{})
				apmProbe.PrometheusProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbePrometheusProbeInputs{
					ConnectorID: promConfig["connector_id"].(string),
					Query:       promConfig["query"].(string),
				}
				// Build TLS config
				if tlsList, ok := promConfig["tls_config"].([]interface{}); ok && len(tlsList) > 0 {
					tlsConfig := tlsList[0].(map[string]interface{})
					apmProbe.PrometheusProbeInputs.TlsConfig = &chaos.ProbeTlsConfigSm{}
					if caCert, ok := tlsConfig["ca_cert_secret"].(string); ok && caCert != "" {
						apmProbe.PrometheusProbeInputs.TlsConfig.CaCrt = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeSecretManager{
							Identifier: caCert,
						}
					}
					if clientCert, ok := tlsConfig["client_cert_secret"].(string); ok && clientCert != "" {
						apmProbe.PrometheusProbeInputs.TlsConfig.ClientCrt = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeSecretManager{
							Identifier: clientCert,
						}
					}
					if clientKey, ok := tlsConfig["client_key_secret"].(string); ok && clientKey != "" {
						apmProbe.PrometheusProbeInputs.TlsConfig.Key = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeSecretManager{
							Identifier: clientKey,
						}
					}
					if insecureSkip, ok := tlsConfig["insecure_skip_verify"].(bool); ok {
						var insecureInterface interface{} = insecureSkip
						apmProbe.PrometheusProbeInputs.TlsConfig.InsecureSkipVerify = &insecureInterface
					}
				}
			}

			// Build Datadog inputs
			if datadogList, ok := apmConfig["datadog_inputs"].([]interface{}); ok && len(datadogList) > 0 {
				datadogConfig := datadogList[0].(map[string]interface{})
				apmProbe.DatadogApmProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeDatadogApmProbeInputs{
					ConnectorID: datadogConfig["connector_id"].(string),
				}
				if query, ok := datadogConfig["query"].(string); ok && query != "" {
					apmProbe.DatadogApmProbeInputs.Query = query
				}
				if duration, ok := datadogConfig["duration_in_min"].(int); ok && duration > 0 {
					var durationInterface interface{} = duration
					apmProbe.DatadogApmProbeInputs.DurationInMin = &durationInterface
				}
				// Build synthetics test
				if syntheticsList, ok := datadogConfig["synthetics_test"].([]interface{}); ok && len(syntheticsList) > 0 {
					syntheticsConfig := syntheticsList[0].(map[string]interface{})
					apmProbe.DatadogApmProbeInputs.SyntheticsTest = &chaos.ProbeSyntheticsTestTemplate{}
					if publicId, ok := syntheticsConfig["public_id"].(string); ok && publicId != "" {
						apmProbe.DatadogApmProbeInputs.SyntheticsTest.PublicId = publicId
					}
					if testType, ok := syntheticsConfig["test_type"].(string); ok && testType != "" {
						testTypeEnum := chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeDatadogSyntheticsTestType(testType)
						apmProbe.DatadogApmProbeInputs.SyntheticsTest.TestType = &testTypeEnum
					}
				}
			}

			// Build Dynatrace inputs
			if dynatraceList, ok := apmConfig["dynatrace_inputs"].([]interface{}); ok && len(dynatraceList) > 0 {
				dynatraceConfig := dynatraceList[0].(map[string]interface{})
				apmProbe.DynatraceApmProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeDynatraceApmProbeInputs{
					ConnectorID: dynatraceConfig["connector_id"].(string),
				}
				if duration, ok := dynatraceConfig["duration_in_min"].(int); ok && duration > 0 {
					var durationInterface interface{} = duration
					apmProbe.DynatraceApmProbeInputs.DurationInMin = &durationInterface
				}
				// Build metrics
				if metricsList, ok := dynatraceConfig["metrics"].([]interface{}); ok && len(metricsList) > 0 {
					metricsConfig := metricsList[0].(map[string]interface{})
					apmProbe.DynatraceApmProbeInputs.Metrics = &chaos.ProbeDynatraceMetricsTemplate{}
					if entitySelector, ok := metricsConfig["entity_selector"].(string); ok && entitySelector != "" {
						apmProbe.DynatraceApmProbeInputs.Metrics.EntitySelector = entitySelector
					}
					if metricsSelector, ok := metricsConfig["metrics_selector"].(string); ok && metricsSelector != "" {
						apmProbe.DynatraceApmProbeInputs.Metrics.MetricsSelector = metricsSelector
					}
				}
			}

			// Build AppDynamics inputs
			if appDynamicsList, ok := apmConfig["app_dynamics_inputs"].([]interface{}); ok && len(appDynamicsList) > 0 {
				appDynamicsConfig := appDynamicsList[0].(map[string]interface{})
				apmProbe.AppDynamicsProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeAppDynamicsProbeInputs{
					ConnectorID: appDynamicsConfig["connector_id"].(string),
				}
				// Build appd_metrics
				if appdMetricsList, ok := appDynamicsConfig["appd_metrics"].([]interface{}); ok && len(appdMetricsList) > 0 {
					appdMetricsConfig := appdMetricsList[0].(map[string]interface{})
					apmProbe.AppDynamicsProbeInputs.AppdMetrics = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeAppdMetrics{}
					if appName, ok := appdMetricsConfig["application_name"].(string); ok && appName != "" {
						apmProbe.AppDynamicsProbeInputs.AppdMetrics.ApplicationName = appName
					}
					if metricsPath, ok := appdMetricsConfig["metrics_full_path"].(string); ok && metricsPath != "" {
						apmProbe.AppDynamicsProbeInputs.AppdMetrics.MetricsFullPath = metricsPath
					}
					if duration, ok := appdMetricsConfig["duration_in_min"].(int); ok && duration > 0 {
						var durationInterface interface{} = duration
						apmProbe.AppDynamicsProbeInputs.AppdMetrics.DurationInMin = &durationInterface
					}
				}
			}

			// Build NewRelic inputs
			if newRelicList, ok := apmConfig["new_relic_inputs"].([]interface{}); ok && len(newRelicList) > 0 {
				newRelicConfig := newRelicList[0].(map[string]interface{})
				apmProbe.NewRelicProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeNewRelicProbeInputs{
					ConnectorID: newRelicConfig["connector_id"].(string),
				}
				// Build new_relic_metric
				if newRelicMetricList, ok := newRelicConfig["new_relic_metric"].([]interface{}); ok && len(newRelicMetricList) > 0 {
					newRelicMetricConfig := newRelicMetricList[0].(map[string]interface{})
					apmProbe.NewRelicProbeInputs.NewRelicMetric = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeNewRelicMetric{}
					if query, ok := newRelicMetricConfig["query"].(string); ok && query != "" {
						apmProbe.NewRelicProbeInputs.NewRelicMetric.Query = query
					}
					if queryMetric, ok := newRelicMetricConfig["query_metric"].(string); ok && queryMetric != "" {
						apmProbe.NewRelicProbeInputs.NewRelicMetric.QueryMetric = queryMetric
					}
				}
			}

			// Build Splunk Observability inputs
			if splunkList, ok := apmConfig["splunk_observability_inputs"].([]interface{}); ok && len(splunkList) > 0 {
				splunkConfig := splunkList[0].(map[string]interface{})
				apmProbe.SplunkObservabilityProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeSplunkObservabilityProbeInputs{
					ConnectorID: splunkConfig["connector_id"].(string),
				}
				// Build splunk_observability_metrics
				if splunkMetricsList, ok := splunkConfig["splunk_observability_metrics"].([]interface{}); ok && len(splunkMetricsList) > 0 {
					splunkMetricsConfig := splunkMetricsList[0].(map[string]interface{})
					apmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeSplunkObservabilityMetrics{}
					if query, ok := splunkMetricsConfig["query"].(string); ok && query != "" {
						apmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.Query = query
					}
					if duration, ok := splunkMetricsConfig["duration_in_min"].(int); ok && duration > 0 {
						var durationInterface interface{} = duration
						apmProbe.SplunkObservabilityProbeInputs.SplunkObservabilityMetrics.DurationInMin = &durationInterface
					}
				}
			}

			// Build GCP Cloud Monitoring inputs
			if gcpList, ok := apmConfig["gcp_cloud_monitoring_inputs"].([]interface{}); ok && len(gcpList) > 0 {
				gcpConfig := gcpList[0].(map[string]interface{})
				apmProbe.GcpCloudMonitoringProbeInputs = &chaos.GithubComHarnessHceSaasHceSdkTemplateSchemaProbeGcpCloudMonitoringProbeInputs{
					ProjectID:         gcpConfig["project_id"].(string),
					Query:             gcpConfig["query"].(string),
					ServiceAccountKey: gcpConfig["service_account_key"].(string),
				}
			}

			req.ProbeProperties = &chaos.ProbeProbeTemplateProperties{
				ApmProbe: apmProbe,
			}
		} else {
			return fmt.Errorf("apm_probe block is required when type is 'apmProbe'")
		}

	default:
		return fmt.Errorf("unsupported probe type: %s (only httpProbe, cmdProbe, k8sProbe, apmProbe are currently supported)", probeType)
	}

	return nil
}

func buildRunPropertiesSimplified(d *schema.ResourceData, req *chaos.ChaosprobetemplateProbeTemplate) error {
	if v, ok := d.GetOk("run_properties"); ok && len(v.([]interface{})) > 0 {
		runConfig := v.([]interface{})[0].(map[string]interface{})
		runProps := &chaos.ProbeProbeTemplateRunProperties{}

		if initialDelay, ok := runConfig["initial_delay"].(string); ok && initialDelay != "" {
			runProps.InitialDelay = initialDelay
		}
		if interval, ok := runConfig["interval"].(string); ok && interval != "" {
			runProps.Interval = interval
		}
		if timeout, ok := runConfig["timeout"].(string); ok && timeout != "" {
			runProps.Timeout = timeout
		}
		if pollingInterval, ok := runConfig["polling_interval"].(string); ok && pollingInterval != "" {
			runProps.PollingInterval = pollingInterval
		}
		if stopOnFailure, ok := runConfig["stop_on_failure"].(bool); ok {
			runProps.StopOnFailure = stopOnFailure
		}
		if verbosity, ok := runConfig["verbosity"].(string); ok && verbosity != "" {
			runProps.Verbosity = verbosity
		}
		// Integer fields - convert to interface{} pointer
		if attempt, ok := runConfig["attempt"].(int); ok && attempt > 0 {
			var attemptInterface interface{} = attempt
			runProps.Attempt = &attemptInterface
		}
		if retry, ok := runConfig["retry"].(int); ok && retry > 0 {
			var retryInterface interface{} = retry
			runProps.Retry = &retryInterface
		}

		req.RunProperties = runProps
	}

	return nil
}

func buildVariablesSimplified(d *schema.ResourceData, req *chaos.ChaosprobetemplateProbeTemplate) {
	if v, ok := d.GetOk("variables"); ok {
		varsList := v.([]interface{})
		if len(varsList) > 0 {
			variables := make([]chaos.TemplateVariable, len(varsList))

			for i, v := range varsList {
				varMap := v.(map[string]interface{})

				name := varMap["name"].(string)
				value := varMap["value"].(string)

				// Value field is *interface{}, convert string to interface{}
				var valueInterface interface{} = value
				variable := chaos.TemplateVariable{
					Name:  name,
					Value: &valueInterface,
				}

				if desc, ok := varMap["description"].(string); ok && desc != "" {
					variable.Description = desc
				}

				if varType, ok := varMap["type"].(string); ok && varType != "" {
					vType := chaos.TemplateVariableType(varType)
					variable.Type_ = &vType
				}

				if required, ok := varMap["required"].(bool); ok {
					variable.Required = required
				}

				variables[i] = variable
			}

			req.Variables = variables
		}
	}
}

// getIntFromInterface safely extracts int value from interface{} pointer
// Returns 0 if nil or cannot convert
func getIntFromInterface(val *interface{}) int {
	if val == nil {
		return 0
	}

	switch v := (*val).(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	case float32:
		return int(v)
	default:
		return 0
	}
}
