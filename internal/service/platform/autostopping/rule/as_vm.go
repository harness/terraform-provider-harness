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

func ResourceVMRule() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness Variables.",

		ReadContext:   resourceVMRuleRead,
		UpdateContext: resourceVMRuleCreateOrUpdate,
		DeleteContext: resourceVMRuleDelete,
		CreateContext: resourceVMRuleCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"name": {
				Description: "Name of the rule",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_connector_id": {
				Description: "Id of the cloud connector",
				Type:        schema.TypeString,
				Required:    true,
			},
			"idle_time_mins": {
				Description: "Idle time in minutes. This is the time that the AutoStopping rule waits before stopping the idle instances.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"use_spot": {
				Description: "Boolean that indicates whether the selected instances should be converted to spot vm",
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
			},
			"custom_domains": {
				Description: "Custom URLs used to access the instances",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_ids": {
							Description: "Ids of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Description: "Tags of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"regions": {
							Description: "Regions of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"zones": {
							Description: "Zones of instances that needs to be managed using the AutoStopping rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"http": {
				Description: "Http routing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Id of the proxy",
							Type:        schema.TypeString,
							Required:    true,
						},
						"routing": {
							Description: "Routing configuration used to access the instances",
							Type:        schema.TypeList,
							MinItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_protocol": {
										Description: "Source protocol of the proxy can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"target_protocol": {
										Description: "Target protocol of the instance can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"source_port": {
										Description: "Port on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"target_port": {
										Description: "Port on the VM",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"action": {
										Description: "Organization Identifier for the Entity",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"health": {
							Description: "Health Check Details",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Description: "Protocol can be http or https",
										Type:        schema.TypeString,
										Required:    true,
									},
									"port": {
										Description: "Health check port on the VM",
										Type:        schema.TypeInt,
										Required:    true,
									},
									"path": {
										Description: "API path to use for health check",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"timeout": {
										Description: "Health check timeout",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_from": {
										Description: "Lower limit for acceptable status code",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"status_code_to": {
										Description: "Upper limit for acceptable status code",
										Type:        schema.TypeInt,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"tcp": {
				Description: "TCP routing configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_id": {
							Description: "Id of the Proxy",
							Type:        schema.TypeString,
							Required:    true,
						},
						"ssh": {
							Description: "SSH configuration",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     22,
									},
								},
							},
						},
						"rdp": {
							Description: "RDP configuration",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3389,
									},
								},
							},
						},
						"forward_rule": {
							Description: "Additional tcp forwarding rules",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connect_on": {
										Description: "Port to listen on the proxy",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"port": {
										Description: "Port to listen on the vm",
										Type:        schema.TypeInt,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"depends": {
				Description: "Dependent rules",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Description: "Rule id of the dependent rule",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"delay_in_sec": {
							Description: "Number of seconds the rule should wait after warming up the dependent rule",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceVMRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		readVMRule(d, resp.Response.Service.Id)
	}

	return nil
}

func resourceVMRuleCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.RuleResponse
	var httpResp *http.Response

	id := d.Id()
	saveServiceRequestV2 := buildVMRule(d, c.AccountId)

	if id == "" {
		resp, httpResp, err = c.CloudCostAutoStoppingRulesV2Api.CreateAutoStoppingRuleV2(ctx, saveServiceRequestV2, c.AccountId, c.AccountId)
	} else {
		resp, httpResp, err = c.CloudCostAutoStoppingRulesV2Api.UpdateAutoStoppingRuleV2(ctx, saveServiceRequestV2, c.AccountId, c.AccountId, id)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Response != nil {
		readVMRule(d, resp.Response.Id)
	}

	return nil
}

func resourceVMRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func buildVMRule(d *schema.ResourceData, accountId string) nextgen.SaveServiceRequestV2 {
	serviceV2 := &nextgen.ServiceV2{}
	serviceV2.AccountIdentifier = accountId
	serviceV2.Kind = "instance"
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

	routingData.Instance = &nextgen.InstanceBasedRoutingDataV2{}

	filter := &nextgen.FilterObject{}
	if attr, ok := d.GetOk("filter"); ok {
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

func readVMRule(d *schema.ResourceData, id int64) {
	identifier := strconv.Itoa(int(id))
	d.SetId(identifier)
	d.Set("identifier", identifier)
}
