package as_rule

import (
	"context"
	"net/http"
	"strconv"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	Database = "database"
	Instance = "instance"
)

func resourceASRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	ruleId, err := strconv.ParseFloat(d.Id(), 64)
	if err != nil {
		return diag.Errorf("invalid rule id")
	}
	resp, httpResp, err := c.CloudCostAutoStoppingRulesApi.AutoStoppingRuleDetails(ctx, c.AccountId, ruleId, c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		readASRule(d, resp.Response.Service.Id)
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
		readASRule(d, resp.Response.Id)
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
			containerSvc.TaskCount = attr.(float64)
		}
	}
	return containerSvc
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
			sshConfig := attr.([]interface{})[0].(map[string]interface{})
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
		if attr, ok := tcpRoutingObj["rdp"]; ok {
			rdpConfig := attr.([]interface{})[0].(map[string]interface{})
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

func readASRule(d *schema.ResourceData, id int64) {
	identifier := strconv.Itoa(int(id))
	d.SetId(identifier)
	d.Set("identifier", identifier)
}
