package as_rule

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

const (
	Database   = "database"
	Instance   = "instance"
	ECS        = "containers"
	K8s        = "k8s"
	ScaleGroup = "clusters"
)

func resourceASRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	ruleId, err := strconv.ParseFloat(d.Id(), 64)
	if err != nil {
		return diag.Errorf("invalid rule id")
	}
	resp, httpResp, err := c.CloudCostAutoStoppingRulesV2Api.GetAutoStoppingRuleV2(ctx, c.AccountId, ruleId, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		readASRule(d, resp.Response.Service)
		setDependencies(d, resp.Response.Deps)
	}

	return nil
}

func resourceASRuleCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, rule nextgen.SaveServiceRequestV2) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RuleResponse
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, rule, c.AccountId, c.AccountId)
	} else {
		resp, httpResp, err = c.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, rule, c.AccountId, c.AccountId, id)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		readASRule(d, resp.Response)
		// Set dependencies from the input rule since RuleResponse doesn't include deps
		setDependencies(d, rule.Deps)
	}

	return nil
}

func resourceASRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ruleId, err := strconv.ParseFloat(d.Id(), 64)
	if err != nil {
		return diag.Errorf("invalid rule id")
	}
	httpResp, err := c.CloudCostAutoStoppingRulesApi.DeleteAutoStoppingRule(ctx, ruleId, c.AccountId, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}
	return nil
}

func buildASRule(d *schema.ResourceData, kind string, accountId string) nextgen.SaveServiceRequestV2 {
	serviceV2 := &nextgen.ServiceV2{}
	serviceV2.AccountIdentifier = accountId
	serviceV2.Kind = kind
	if attr, ok := d.GetOk("name"); ok {
		serviceV2.Name = attr.(string)
	}

	if attr, ok := d.GetOk("cloud_connector_id"); ok {
		serviceV2.CloudAccountId = attr.(string)
	}
	serviceV2.Fulfilment = "ondemand"
	if attr, ok := d.GetOk("use_spot"); ok {
		onDemand := attr.(bool)
		if onDemand {
			serviceV2.Fulfilment = "ondemand"
		} else {
			serviceV2.Fulfilment = "spot"
		}
	}
	serviceV2.IdleTimeMins = 15
	if attr, ok := d.GetOk("idle_time_mins"); ok {
		serviceV2.IdleTimeMins = attr.(int)
	}

	if attr, ok := d.GetOk("custom_domains"); ok {
		domains := make([]string, 0)
		for _, v := range attr.([]interface{}) {
			domains = append(domains, v.(string))
		}
		serviceV2.CustomDomains = domains
	}

	opts := &nextgen.Opts{}
	if attr, ok := d.GetOk("dry_run"); ok {
		opts.DryRun = attr.(bool)
	}
	serviceV2.Opts = opts

	routingData := &nextgen.RoutingDataV2{}
	httpProxy, tcpProxy, healthCheck := getRoutingConfigurations(d)
	if httpProxy != nil {
		routingData.Http = httpProxy
	}
	if tcpProxy != nil {
		routingData.Tcp = tcpProxy
	}
	if healthCheck != nil {
		serviceV2.HealthCheck = healthCheck
	}

	rdsDatabase := getDatabaseConfig(d)
	if rdsDatabase != nil {
		routingData.Database = rdsDatabase
	}

	containerSvc := getContainerConfig(d)
	if containerSvc != nil {
		routingData.ContainerSvc = containerSvc
	}

	if attr, ok := d.GetOk("filter"); ok {
		filter := &nextgen.FilterObject{}
		filterObj := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := filterObj["vm_ids"]; ok {
			vmIds := make([]string, 0)
			for _, v := range attr.([]interface{}) {
				vmIds = append(vmIds, v.(string))
			}
			filter.Ids = vmIds
		}
		if attr, ok := filterObj["regions"]; ok {
			regions := make([]string, 0)
			for _, v := range attr.([]interface{}) {
				regions = append(regions, v.(string))
			}
			filter.Regions = regions
		}
		if attr, ok := filterObj["zones"]; ok {
			zones := make([]string, 0)
			for _, v := range attr.([]interface{}) {
				zones = append(zones, v.(string))
			}
			filter.Zones = zones
		}
		if attr, ok := filterObj["tags"]; ok {
			filter.Tags = make(map[string]string)
			for _, tag := range attr.([]interface{}) {
				item := tag.(map[string]interface{})
				filter.Tags[item["key"].(string)] = item["value"].(string)
			}
		}
		routingData.Instance = &nextgen.InstanceBasedRoutingDataV2{}
		routingData.Instance.Filter = filter

	}
	scaleGroup := getScaleGroupConfig(d)
	if scaleGroup != nil {
		if routingData.Instance == nil {
			routingData.Instance = &nextgen.InstanceBasedRoutingDataV2{}
		}
		routingData.Instance.ScaleGroup = scaleGroup
	}

	k8sConfig := getK8sConfig(d)
	if k8sConfig != nil {
		routingData.K8s = k8sConfig
		serviceV2.Fulfilment = "kubernetes"
	}

	serviceV2.Routing = routingData
	deps := getDependencies(d)
	saveServiceRequestV2 := &nextgen.SaveServiceRequestV2{
		Service: serviceV2,
		Deps:    deps,
	}
	return *saveServiceRequestV2
}

func getDatabaseConfig(d *schema.ResourceData) *nextgen.RdsDatabase {
	var rdsDatabase *nextgen.RdsDatabase
	if attr, ok := d.GetOk("database"); ok {
		rdsDatabase = &nextgen.RdsDatabase{}
		databaseObj := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := databaseObj["id"]; ok {
			rdsDatabase.Id = attr.(string)
		}
		if attr, ok := databaseObj["region"]; ok {
			rdsDatabase.Region = attr.(string)
		}
	}
	return rdsDatabase
}

func getContainerConfig(d *schema.ResourceData) *nextgen.ContainerSvc {
	var containerSvc *nextgen.ContainerSvc
	if attr, ok := d.GetOk("container"); ok {
		containerSvc = &nextgen.ContainerSvc{}
		databaseObj := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := databaseObj["cluster"]; ok {
			containerSvc.Cluster = attr.(string)
		}
		if attr, ok := databaseObj["region"]; ok {
			containerSvc.Region = attr.(string)
		}
		if attr, ok := databaseObj["service"]; ok {
			containerSvc.Service = attr.(string)
		}
		containerSvc.TaskCount = 1
		if attr, ok := databaseObj["task_count"]; ok {
			desCount, ok := attr.(int64)
			if ok {
				containerSvc.TaskCount = float64(desCount)
			}
		}
	}
	return containerSvc
}

func getScaleGroupConfig(d *schema.ResourceData) *nextgen.AsgMinimal {
	var scaleGroup *nextgen.AsgMinimal
	if attr, ok := d.GetOk("scale_group"); ok {
		scaleGroup = &nextgen.AsgMinimal{}
		scaleGroupObj := attr.([]interface{})[0].(map[string]interface{})

		// Handle string fields
		if attr, ok := scaleGroupObj["id"]; ok && attr.(string) != "" {
			scaleGroup.Id = attr.(string)
		}
		if attr, ok := scaleGroupObj["name"]; ok && attr.(string) != "" {
			scaleGroup.Name = attr.(string)
		}
		if attr, ok := scaleGroupObj["region"]; ok && attr.(string) != "" {
			scaleGroup.Region = attr.(string)
		}
		if attr, ok := scaleGroupObj["zone"]; ok && attr.(string) != "" {
			scaleGroup.AvailabilityZones = []string{attr.(string)}
		}

		// Handle numeric fields
		if attr, ok := scaleGroupObj["desired"]; ok && attr != nil {
			if desired, ok := attr.(int); ok {
				scaleGroup.Desired = int32(desired)
			} else if desired, ok := attr.(int64); ok {
				scaleGroup.Desired = int32(desired)
			}
		}
		if attr, ok := scaleGroupObj["min"]; ok && attr != nil {
			if min, ok := attr.(int); ok {
				scaleGroup.Min = int32(min)
			} else if min, ok := attr.(int64); ok {
				scaleGroup.Min = int32(min)
			}
		}
		if attr, ok := scaleGroupObj["max"]; ok && attr != nil {
			if max, ok := attr.(int); ok {
				scaleGroup.Max = int32(max)
			} else if max, ok := attr.(int64); ok {
				scaleGroup.Max = int32(max)
			}
		}
		if attr, ok := scaleGroupObj["on_demand"]; ok && attr != nil {
			if onDemand, ok := attr.(int); ok {
				scaleGroup.OnDemand = int32(onDemand)
			} else if onDemand, ok := attr.(int64); ok {
				scaleGroup.OnDemand = int32(onDemand)
			}
		}
	}
	return scaleGroup
}

func getK8sConfig(d *schema.ResourceData) *nextgen.RoutingDataV2K8s {
	if attr, ok := d.GetOk("rule_yaml"); !ok || attr.(string) == "" {
		return nil
	}
	yamlStr := d.Get("rule_yaml").(string)
	jsonStr, err := yamlToJSON(yamlStr)
	if err != nil {
		return nil
	}
	k8s := &nextgen.RoutingDataV2K8s{}
	k8s.RuleJson = jsonStr
	if attr, ok := d.GetOk("k8s_connector_id"); ok && attr.(string) != "" {
		k8s.ConnectorID = attr.(string)
	}
	if attr, ok := d.GetOk("k8s_namespace"); ok && attr.(string) != "" {
		k8s.Namespace = attr.(string)
	}
	return k8s
}

// yamlToJSON converts a YAML string to JSON. Input may be YAML or JSON (JSON is a subset of YAML).
func yamlToJSON(yamlStr string) (string, error) {
	var v interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &v); err != nil {
		return "", err
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// jsonToYAML converts a JSON string to YAML for storing in rule_yaml on read.
func jsonToYAML(jsonStr string) (string, error) {
	var v interface{}
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		return "", err
	}
	return canonicalRuleYAML(v)
}

// canonicalRuleYAML marshals v to YAML with 2-space indent so state has consistent format
// and avoids drift from indentation or key-order differences.
func canonicalRuleYAML(v interface{}) (string, error) {
	var buf strings.Builder
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(v); err != nil {
		return "", err
	}
	if err := enc.Close(); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

func getDependencies(d *schema.ResourceData) []nextgen.ServiceDep {
	dependencies := d.Get("depends").([]interface{})
	dependencyList := make([]nextgen.ServiceDep, 0)
	if len(dependencies) == 0 {
		return dependencyList
	}
	for _, dep := range dependencies {
		id := dep.(map[string]interface{})["rule_id"].(int)
		delay := dep.(map[string]interface{})["delay_in_sec"].(int)
		dependency := &nextgen.ServiceDep{
			DepId:     int64(id),
			DelaySecs: int32(delay),
		}
		dependencyList = append(dependencyList, *dependency)
	}
	return dependencyList
}

// setDependencies sets the rule dependencies in Terraform state from the API response
func setDependencies(d *schema.ResourceData, deps []nextgen.ServiceDep) {
	dependsList := make([]map[string]interface{}, 0, len(deps))
	for _, dep := range deps {
		dependsList = append(dependsList, map[string]interface{}{
			"rule_id":      int(dep.DepId),
			"delay_in_sec": int(dep.DelaySecs),
		})
	}
	d.Set("depends", dependsList)
}

func getRoutingConfigurations(d *schema.ResourceData) (*nextgen.HttpProxy, *nextgen.TcpProxy, *nextgen.HealthCheck) {
	var httpProxy *nextgen.HttpProxy
	var tcpProxy *nextgen.TcpProxy
	var healthCheck *nextgen.HealthCheck
	if attr, ok := d.GetOk("http"); ok {
		httpProxy = &nextgen.HttpProxy{}
		httpRoutingObj := attr.([]interface{})[0].(map[string]interface{})

		if attr, ok := httpRoutingObj["proxy_id"]; ok {
			proxy := &nextgen.Proxy{
				Id: attr.(string),
			}
			httpProxy.Proxy = proxy
		}
		if attr, ok := httpRoutingObj["routing"]; ok {
			routingConfigs := attr.([]interface{})
			portConfigsList := make([]nextgen.PortConfig, 0)
			for _, routingConfig := range routingConfigs {
				routingObj := routingConfig.(map[string]interface{})
				portConfig := &nextgen.PortConfig{}
				if attr, ok := routingObj["source_protocol"]; ok {
					portConfig.Protocol = attr.(string)
				}
				if attr, ok := routingObj["target_protocol"]; ok {
					portConfig.TargetProtocol = attr.(string)
				}
				if attr, ok := routingObj["source_port"]; ok {
					portConfig.Port = attr.(int)
				}
				if attr, ok := routingObj["target_port"]; ok {
					portConfig.TargetPort = attr.(int)
				}
				if attr, ok := routingObj["action"]; ok {
					portConfig.Action = attr.(string)
				}
				portConfig.RoutingRules = []nextgen.RoutingRule{}

				portConfigsList = append(portConfigsList, *portConfig)
			}
			httpProxy.Ports = portConfigsList
		}
		if attr, ok := httpRoutingObj["health"]; ok {
			healthCheck = &nextgen.HealthCheck{}
			healthConfig := attr.([]interface{})[0].(map[string]interface{})
			if attr, ok := healthConfig["protocol"]; ok {
				healthCheck.Protocol = attr.(string)
			}
			if attr, ok := healthConfig["port"]; ok {
				healthCheck.Port = attr.(int)
			}
			if attr, ok := healthConfig["path"]; ok {
				healthCheck.Path = attr.(string)
			}
			if attr, ok := healthConfig["timeout"]; ok {
				healthCheck.Timeout = attr.(int)
			}
			if attr, ok := healthConfig["status_code_from"]; ok {
				healthCheck.StatusCodeFrom = attr.(int)
			}
			if attr, ok := healthConfig["status_code_to"]; ok {
				healthCheck.StatusCodeTo = attr.(int)
			}
		}
	}

	if attr, ok := d.GetOk("tcp"); ok {
		tcpProxy = &nextgen.TcpProxy{}
		tcpRoutingObj := attr.([]interface{})[0].(map[string]interface{})

		if attr, ok := tcpRoutingObj["proxy_id"]; ok {
			proxy := &nextgen.Proxy{
				Id: attr.(string),
			}
			tcpProxy.Proxy = proxy
		}
		if attr, ok := tcpRoutingObj["ssh"]; ok {
			if valArr, ok := attr.([]interface{}); ok && len(valArr) > 0 {
				sshConfig := valArr[0].(map[string]interface{})
				ssh := &nextgen.ServiceRoutingTcpPort{}
				if attr, ok := sshConfig["connect_on"]; ok {
					ssh.Source = attr.(int)
				}
				ssh.Target = 22
				if attr, ok := sshConfig["port"]; ok {
					ssh.Target = attr.(int)
				}
				tcpProxy.SshConf = ssh
			}
		}
		if attr, ok := tcpRoutingObj["rdp"]; ok {
			if valArr, ok := attr.([]interface{}); ok && len(valArr) > 0 {
				rdpConfig := valArr[0].(map[string]interface{})
				rdp := &nextgen.ServiceRoutingTcpPort{}
				if attr, ok := rdpConfig["connect_on"]; ok {
					rdp.Source = attr.(int)
				}
				rdp.Target = 3389
				if attr, ok := rdpConfig["port"]; ok {
					rdp.Target = attr.(int)
				}
				tcpProxy.RdpConf = rdp
			}
		}
		if attr, ok := tcpRoutingObj["forward_rule"]; ok {
			forwardRules := attr.([]interface{})
			tcpPortList := make([]nextgen.ServiceRoutingTcpPort, 0)
			for _, forwardRule := range forwardRules {
				rule := forwardRule.(map[string]interface{})
				tcpPort := &nextgen.ServiceRoutingTcpPort{}
				if attr, ok := rule["connect_on"]; ok {
					tcpPort.Source = attr.(int)
				}
				if attr, ok := rule["port"]; ok {
					tcpPort.Target = attr.(int)
				}
				tcpPortList = append(tcpPortList, *tcpPort)
			}
			tcpProxy.CustomPorts = tcpPortList
		}
	}

	return httpProxy, tcpProxy, healthCheck
}

// setRoutingConfig sets the HTTP and TCP routing configurations in Terraform state from the API response.
// When setConnect is true (VM rule), also sets the computed connect block from TCP ssh/rdp source ports.
// Always sets both http and tcp to ensure stale data is cleared when configs are removed
func setRoutingConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2, healthCheck *nextgen.HealthCheck) {
	// Set HTTP routing config (or clear it if absent)
	if routing != nil && routing.Http != nil {
		httpConfig := make(map[string]interface{})

		// Set proxy_id
		if routing.Http.Proxy != nil && routing.Http.Proxy.Id != "" {
			httpConfig["proxy_id"] = routing.Http.Proxy.Id
		}

		// Set routing (port configs)
		if len(routing.Http.Ports) > 0 {
			routingList := make([]map[string]interface{}, 0, len(routing.Http.Ports))
			for _, portConfig := range routing.Http.Ports {
				routingEntry := map[string]interface{}{
					"source_protocol": portConfig.Protocol,
					"target_protocol": portConfig.TargetProtocol,
					"source_port":     portConfig.Port,
					"target_port":     portConfig.TargetPort,
					"action":          portConfig.Action,
				}
				routingList = append(routingList, routingEntry)
			}
			httpConfig["routing"] = routingList
		}

		// Set health check config (nested inside http)
		if healthCheck != nil {
			healthConfig := []map[string]interface{}{
				{
					"protocol":         healthCheck.Protocol,
					"port":             healthCheck.Port,
					"path":             healthCheck.Path,
					"timeout":          healthCheck.Timeout,
					"status_code_from": healthCheck.StatusCodeFrom,
					"status_code_to":   healthCheck.StatusCodeTo,
				},
			}
			httpConfig["health"] = healthConfig
		}

		d.Set("http", []map[string]interface{}{httpConfig})
	} else {
		// Clear http config if not present in API response
		d.Set("http", []map[string]interface{}{})
	}

	// Set TCP routing config (or clear it if absent)
	if routing != nil && routing.Tcp != nil {
		tcpConfig := make(map[string]interface{})

		// Set proxy_id
		if routing.Tcp.Proxy != nil && routing.Tcp.Proxy.Id != "" {
			tcpConfig["proxy_id"] = routing.Tcp.Proxy.Id
		}

		// Set SSH config
		if routing.Tcp.SshConf != nil {
			sshConfig := []map[string]interface{}{
				{
					"connect_on": routing.Tcp.SshConf.Source,
					"port":       routing.Tcp.SshConf.Target,
				},
			}
			tcpConfig["ssh"] = sshConfig
		}

		// Set RDP config
		if routing.Tcp.RdpConf != nil {
			rdpConfig := []map[string]interface{}{
				{
					"connect_on": routing.Tcp.RdpConf.Source,
					"port":       routing.Tcp.RdpConf.Target,
				},
			}
			tcpConfig["rdp"] = rdpConfig
		}

		// Set forward rules (custom ports)
		if len(routing.Tcp.CustomPorts) > 0 {
			forwardRules := make([]map[string]interface{}, 0, len(routing.Tcp.CustomPorts))
			for _, tcpPort := range routing.Tcp.CustomPorts {
				forwardRule := map[string]interface{}{
					"connect_on": tcpPort.Source,
					"port":       tcpPort.Target,
				}
				forwardRules = append(forwardRules, forwardRule)
			}
			tcpConfig["forward_rule"] = forwardRules
		}

		d.Set("tcp", []map[string]interface{}{tcpConfig})
		connectBlock := buildConnectBlock(routing)
		_ = d.Set("connect", []map[string]interface{}{connectBlock})
	} else {
		// Clear tcp config if not present in API response
		d.Set("tcp", []map[string]interface{}{})
		_ = d.Set("connect", []map[string]interface{}{})
	}
}

func buildConnectBlock(routing *nextgen.RoutingDataV2) map[string]interface{} {
	connectBlock := map[string]interface{}{
		"ssh": 0,
		"rdp": 0,
	}
	if routing.Tcp.SshConf != nil {
		connectBlock["ssh"] = routing.Tcp.SshConf.Source
	}
	if routing.Tcp.RdpConf != nil {
		connectBlock["rdp"] = routing.Tcp.RdpConf.Source
	}
	return connectBlock
}

// setDatabaseConfig sets the database configuration in Terraform state from the API response
func setDatabaseConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2) {
	if routing == nil || routing.Database == nil {
		return
	}
	database := []map[string]interface{}{
		{
			"id":     routing.Database.Id,
			"region": routing.Database.Region,
		},
	}
	d.Set("database", database)
}

// setContainerConfig sets the container (ECS) configuration in Terraform state from the API response
func setContainerConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2) {
	if routing == nil || routing.ContainerSvc == nil {
		return
	}
	container := []map[string]interface{}{
		{
			"cluster":    routing.ContainerSvc.Cluster,
			"service":    routing.ContainerSvc.Service,
			"region":     routing.ContainerSvc.Region,
			"task_count": int(routing.ContainerSvc.TaskCount),
		},
	}
	d.Set("container", container)
}

// setScaleGroupConfig sets the scale group configuration in Terraform state from the API response
func setScaleGroupConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2) {
	if routing == nil || routing.Instance == nil || routing.Instance.ScaleGroup == nil {
		return
	}
	sg := routing.Instance.ScaleGroup

	// Get zone from AvailabilityZones array (first element if present)
	zone := ""
	if len(sg.AvailabilityZones) > 0 {
		zone = sg.AvailabilityZones[0]
	}

	scaleGroup := []map[string]interface{}{
		{
			"id":        sg.Id,
			"name":      sg.Name,
			"region":    sg.Region,
			"zone":      zone,
			"desired":   int(sg.Desired),
			"min":       int(sg.Min),
			"max":       int(sg.Max),
			"on_demand": int(sg.OnDemand),
		},
	}
	d.Set("scale_group", scaleGroup)
}

// setFilterConfig sets the VM filter configuration in Terraform state from the API response
func setFilterConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2) {
	if routing == nil || routing.Instance == nil || routing.Instance.Filter == nil {
		return
	}
	filterObj := routing.Instance.Filter

	// Convert tags from map[string]string to []map[string]interface{} to match schema
	// Sort keys to ensure deterministic ordering and avoid spurious Terraform diffs
	var tagsList []map[string]interface{}
	tagKeys := make([]string, 0, len(filterObj.Tags))
	for key := range filterObj.Tags {
		tagKeys = append(tagKeys, key)
	}
	sort.Strings(tagKeys)
	for _, key := range tagKeys {
		tagsList = append(tagsList, map[string]interface{}{
			"key":   key,
			"value": filterObj.Tags[key],
		})
	}

	filter := []map[string]interface{}{
		{
			"vm_ids":  filterObj.Ids,
			"regions": filterObj.Regions,
			"zones":   filterObj.Zones,
			"tags":    tagsList,
		},
	}
	d.Set("filter", filter)
}

// setK8sConfig sets the K8s routing configuration in Terraform state from the API response.
func setK8sConfig(d *schema.ResourceData, routing *nextgen.RoutingDataV2) {
	if routing == nil || routing.K8s == nil {
		return
	}
	k8s := routing.K8s
	d.Set("k8s_connector_id", k8s.ConnectorID)
	d.Set("k8s_namespace", k8s.Namespace)
	ruleYaml := k8s.RuleJson
	if yamlStr, err := jsonToYAML(k8s.RuleJson); err == nil {
		ruleYaml = yamlStr
	} else {
		var v interface{}
		if err := yaml.Unmarshal([]byte(k8s.RuleJson), &v); err == nil {
			if canonical, err := canonicalRuleYAML(v); err == nil {
				ruleYaml = canonical
			}
		}
	}
	d.Set("rule_yaml", ruleYaml)
}

func readASRule(d *schema.ResourceData, service *nextgen.ServiceV2) {
	if service == nil {
		return
	}
	identifier := strconv.Itoa(int(service.Id))
	d.SetId(identifier)
	d.Set("identifier", identifier)
	d.Set("name", service.Name)
	d.Set("cloud_connector_id", service.CloudAccountId)
	d.Set("idle_time_mins", service.IdleTimeMins)
	if service.Opts != nil {
		d.Set("dry_run", service.Opts.DryRun)
	}
	d.Set("use_spot", service.Fulfilment == "spot")
	d.Set("custom_domains", service.CustomDomains)

	// Set routing-related fields based on rule kind
	switch service.Kind {
	case Database:
		setDatabaseConfig(d, service.Routing)
	case ECS:
		setContainerConfig(d, service.Routing)
	case K8s:
		setK8sConfig(d, service.Routing)
	case ScaleGroup:
		setScaleGroupConfig(d, service.Routing)
	case Instance:
		setFilterConfig(d, service.Routing)
	}
	// Always call setRoutingConfig to ensure stale http/tcp configs are cleared
	setRoutingConfig(d, service.Routing, service.HealthCheck)
}
