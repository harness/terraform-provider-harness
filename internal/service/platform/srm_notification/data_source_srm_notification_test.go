package srm_notification_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSrmNotification(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6)) //Add with muliptle logs and metrics
	name := id
	resourceName := "data.harness_platform_monitored_service.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccELKDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicMetricDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicLogDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSplunkSignalFXDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGrafanaLokiLogsDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureMetricsDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureLogsDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPrometheusMetricsDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
	id = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDatadogMetricsDataSourceMonitoredService(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccELKDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "name"
				identifier = "identifier"
				type = "ElasticSearch"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "name"
                    identifier = "identifier"
					query = "query"
                    groupName = "Logs Group"
                    queryParams = {
					  index = "index"
					  serviceInstanceField = "serviceInstanceIdentifier"
					  timeStampIdentifier = "timeStampIdentifier"
					  timeStampFormat = "timeStampFormat"
					  messageIdentifier = "messageIdentifier"
                    }
					},
					{
					name  = "name2"
                    identifier = "identifier2"
                    groupName = "Logs Group"
                    query = "query"
					queryParams = {
					  index = "index"
					  serviceInstanceField = "serviceInstanceIdentifier"
					  timeStampIdentifier = "timeStampIdentifier"
					  timeStampFormat = "timeStampFormat"
					  messageIdentifier = "messageIdentifier"
                    }
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccSumologicMetricDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "sumologicmetrics"
				identifier = "sumo_metric_identifier"
				type = "SumologicMetrics"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "metric_cpu"
                    identifier = "metric_cpu"
					query = "metric=cpu"
                    groupName = "g1"
                    queryParams = {
                    }
                    riskProfile = {
                    riskCategory = "Performance_Other"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER"
                    ]
                    }
                    liveMonitoringEnabled = "true"
                    continuousVerificationEnabled = "true"
                    sliEnabled = "false"
					},
					{
					name  = "name2"
                    identifier = "identifier2"
                    groupName = "g2"
                    query = "metric=memory"
					queryParams = {
                    }
                    riskProfile = {
                    riskCategory = "Performance_Other"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER"
                    ]
                    }
                    liveMonitoringEnabled = "false"
                    continuousVerificationEnabled = "false"
                    sliEnabled = "false"
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccSumologicLogDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "sumologic"
				identifier = "sumo_metric_identifier"
				type = "SumologicLogs"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "log1"
                    identifier = "log1"
					query = "*"
                    groupName = "Logs Group"
                    queryParams = {
                    serviceInstanceField = "_sourcehost"
                    }
					},
					{
					name  = "log2"
                    identifier = "identifier2"
                    groupName = "g2"
                    query = "error"
                    queryParams = {
                    serviceInstanceField = "_sourcehost"
                    }
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccSplunkSignalFXDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "signalfxmetrics"
				identifier = "signalfxmetrics"
				type = "SplunkSignalFXMetrics"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "metric_infra_cpu"
                    identifier = "metric_infra_cpu"
					query = "***"
                    groupName = "g"
                    riskProfile = {
                    riskCategory = "Errors"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER",
                    "ACT_WHEN_LOWER"
                    ]
                    }
                    liveMonitoringEnabled = "true"
                    continuousVerificationEnabled = "true"
                    sliEnabled = "false"
					},
					{
					name  = "name2"
                    identifier = "identifier2"
                    groupName = "g2"
                    query = "*"
                    riskProfile = {
                    riskCategory = "Performance_Other"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER"
                    ]
                    }
                    liveMonitoringEnabled = "true"
                    continuousVerificationEnabled = "false"
                    sliEnabled = "false"
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccGrafanaLokiLogsDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "Test"
				identifier = "Test"
				type = "GrafanaLokiLogs"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "Demo"
                    identifier = "Demo"
					query = "{job=~\".+\"}"
                    groupName = "Log_Group"
                    queryParams = {
                    serviceInstanceField = "job"
                    }
					},
					{
					name  = "log2"
                    identifier = "identifier2"
                    groupName = "g2"
                    query = "error"
                    queryParams = {
                    serviceInstanceField = "_sourcehost"
                    }
                    liveMonitoringEnabled = "false"
                    continuousVerificationEnabled = "false"
                    sliEnabled = "false"
					}
				]})
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccAzureMetricsDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = ["foo:bar", "bar:foo"]
			health_sources {
				name = "azure metrics verify step"
				identifier = "azure_metrics_verify_step"
				type = "AzureMetrics"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name = "metric"
                    identifier = "metric"
					query = "default"
                    groupName = "g1"
                    queryParams = {
                    serviceInstanceField = "host"
                    index = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test",
                    healthSourceMetricName = "cpuUsagePercentage",
                    healthSourceMetricNamespace = "insights.container/nodes",
                    aggregationType = "average"
                    }
                    riskProfile = {
                    riskCategory = "Performance_Other"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER"
                    ]
                    }
                    liveMonitoringEnabled = "true"
                    continuousVerificationEnabled = "true"
                    sliEnabled = "false"
					},
					{
					name  = "name2"
                    identifier = "identifier2"
                    groupName = "g2"
                    queryParams = {
                    serviceInstanceField = "host"
                    index = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test",
                    healthSourceMetricName = "cpuUsagePercentage",
                    healthSourceMetricNamespace = "insights.container/nodes",
                    aggregationType = "average"
                    }
                    riskProfile = {
                    riskCategory = "Performance_Other"
                    thresholdTypes = [
                    "ACT_WHEN_HIGHER"
                    ]
                    }
                    liveMonitoringEnabled = "false"
                    continuousVerificationEnabled = "false"
                    sliEnabled = "false"
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccAzureLogsDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_monitored_service" "test" {
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		identifier = "%[1]s"
		request {
			name = "%[2]s"
			type = "Application"
			description = "description"
			service_ref = "service_ref"
			environment_ref = "environment_ref"
			tags = []
			health_sources {
				name = "Demo azure"
				identifier = "Demo_azure"
				type = "AzureLogs"
                version = "v2"
				spec = jsonencode({
				connectorRef = "connectorRef"
				queryDefinitions = [
					{
					name  = "name2"
                    identifier = "identifier2"
                    groupName = "g2"
                    query = "*"
					queryParams = {
                    serviceInstanceField = "Name",
                    timeStampIdentifier = "StartedTime",
                    messageIdentifier = "Image",
                    index = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test"
                    }
                    liveMonitoringEnabled = "false"
                    continuousVerificationEnabled = "false"
					}
				]})
			}
			change_sources {
				name = "csName1"
				identifier = "harness_cd_next_gen"
				type = "HarnessCDNextGen"
				enabled = true
				spec = jsonencode({
				})
				category = "Deployment"
			}
			template_ref = "template_ref"
			version_label = "version_label"
		}
	}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccPrometheusMetricsDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

resource "harness_platform_monitored_service" "test" {
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  identifier = "%[1]s"
  request {
    name            = "%[2]s"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "prometheus metrics verify step"
      identifier = "prometheus_metrics"
      type       = "Prometheus"
      spec = jsonencode({
        connectorRef = "connectorRef"
        metricDefinitions = [
          {
            identifier = "Prometheus_Metric",
            metricName = "Prometheus Metric",
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            analysis = {
              liveMonitoring = {
                enabled = true
              }
              deploymentVerification = {
                enabled                  = true
              }
            }
            sli : {
              enabled = true
            }
            query                    = "count(up{group=\"cv\",group=\"cv\"})"
            groupName                = "met"
            serviceInstanceFieldName = "pod_name"
            isManualQuery            = true
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}

func testAccDatadogMetricsDataSourceMonitoredService(id string, name string) string {
	return fmt.Sprintf(
		`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

resource "harness_platform_monitored_service" "test" {
  org_id     = harness_platform_organization.test.id
  project_id = harness_platform_project.test.id
  identifier = "%[1]s"
  request {
    name            = "%[2]s"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "ddm"
      identifier = "ddm"
      type       = "DatadogMetrics"
      spec = jsonencode({
        connectorRef = "connectorRef"
        feature = "Datadog Cloud Metrics"
        metricDefinitions = [
          {
            metricName       = "metric"
            identifier = "metric"
            query      = "avg:kubernetes.cpu.limits{*}.rollup(avg, 60);\navg:kubernetes.cpu.limits{*}.rollup(avg, 30);\n(a+b)/10"
            isManualQuery = true
            isCustomCreatedMetric = true
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            analysis = {
              liveMonitoring = {
                enabled = true
              }
              deploymentVerification = {
                enabled                  = true
                serviceInstanceFieldName = "group"
              }
            }
            sli : {
              enabled = true
            }
          },
          {
            metricName    = "dashboard_metric_cpu"
            identifier    = "metric_cpu"
            query         = "avg:kubernetes.cpu.limits{*}.rollup(avg, 60);\navg:kubernetes.cpu.limits{*}.rollup(avg, 30);\n(a+b)/10"
            isManualQuery = false
            dashboardName = "dashboard"
            metricPath = "M1"
            groupingQuery = "avg:kubernetes.cpu.limits{*} by {host}.rollup(avg, 60)"
            metric = "kubernetes.cpu.limits"
            aggregation = "avg"
            isCustomCreatedMetric = true
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            analysis = {
              liveMonitoring = {
                enabled = true
              }
              deploymentVerification = {
                enabled                  = true
                serviceInstanceFieldName = "group"
              }
            }
            sli : {
              enabled = true
            }
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}

	data "harness_platform_monitored_service" "test" {
		identifier = harness_platform_monitored_service.test.identifier
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}`,
		id, name)
}
