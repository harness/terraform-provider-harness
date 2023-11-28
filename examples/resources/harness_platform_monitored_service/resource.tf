#Sample template for Elastic Search Log Health Source
resource "harness_platform_monitored_service" "example" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "name"
      identifier = "identifier"
      type       = "ElasticSearch"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name      = "name"
            query     = "query"
            index     = "index"
            groupName = "Logs_Group"
            queryParams = {
              index                = "index"
              serviceInstanceField = "serviceInstanceIdentifier"
              timeStampIdentifier  = "timeStampIdentifier"
              timeStampFormat      = "timeStampFormat"
              messageIdentifier    = "messageIdentifier"
            }
          },
          {
            name      = "name2"
            query     = "query2"
            index     = "index2"
            groupName = "Logs_Group"
            queryParams = {
              index                = "index"
              serviceInstanceField = "serviceInstanceIdentifier"
              timeStampIdentifier  = "timeStampIdentifier"
              timeStampFormat      = "timeStampFormat"
              messageIdentifier    = "messageIdentifier"
            }
          }
        ]
      })
    }
    change_sources {
      name       = "csName1"
      identifier = "harness_cd_next_gen"
      type       = "HarnessCDNextGen"
      enabled    = true
      spec = jsonencode({
      })
      category = "Deployment"
    }
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref"
      enabled               = true
    }
    notification_rule_refs {
      notification_rule_ref = "notification_rule_ref1"
      enabled               = false
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}
#Sample template for Sumologic Metrics Health Source
resource "harness_platform_monitored_service" "example1" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "sumologicmetrics"
      identifier = "sumo_metric_identifier"
      type       = "SumologicMetrics"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "metric_cpu"
            identifier = "metric_cpu"
            query      = "metric=cpu"
            groupName  = "g1"
            queryParams = {
            }
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            liveMonitoringEnabled         = "true"
            continuousVerificationEnabled = "true"
            sliEnabled                    = "false"
            metricThresholds: [
              {
                type: "IgnoreThreshold",
                spec: {
                  action: "Ignore"
                },
                criteria: {
                  type: "Absolute",
                  spec: {
                    greaterThan: 100
                  }
                },
                metricType: "Custom",
                metricName: "metric_cpu"
              }
            ]
          },
          {
            name       = "name2"
            identifier = "identifier2"
            groupName  = "g2"
            query      = "metric=memory"
            queryParams = {
            }
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            liveMonitoringEnabled         = "false"
            continuousVerificationEnabled = "false"
            sliEnabled                    = "false"
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}
#Sample template for Sumologic Log Health Source
resource "harness_platform_monitored_service" "example2" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "sumologic"
      identifier = "sumo_metric_identifier"
      type       = "SumologicLogs"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "log1"
            identifier = "log1"
            query      = "*"
            groupName  = "Logs Group"
            queryParams = {
              serviceInstanceField = "_sourcehost"
            }
          },
          {
            name       = "log2"
            identifier = "identifier2"
            groupName  = "g2"
            query      = "error"
            queryParams = {
              serviceInstanceField = "_sourcehost"
            }
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}

#Sample template for Splunk Signal FX Health Source
resource "harness_platform_monitored_service" "example3" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "signalfxmetrics"
      identifier = "signalfxmetrics"
      type       = "SplunkSignalFXMetrics"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "metric_infra_cpu"
            identifier = "metric_infra_cpu"
            query      = "***"
            groupName  = "g"
            riskProfile = {
              riskCategory = "Errors"
              thresholdTypes = [
                "ACT_WHEN_HIGHER",
                "ACT_WHEN_LOWER"
              ]
            }
            liveMonitoringEnabled         = "true"
            continuousVerificationEnabled = "true"
            sliEnabled                    = "false"
          },
          {
            name       = "name2"
            identifier = "identifier2"
            groupName  = "g2"
            query      = "*"
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            liveMonitoringEnabled         = "true"
            continuousVerificationEnabled = "false"
            sliEnabled                    = "false"
            metricThresholds: [
              {
                type: "IgnoreThreshold",
                spec: {
                  action: "Ignore"
                },
                criteria: {
                  type: "Absolute",
                  spec: {
                    greaterThan: 100
                  }
                },
                metricType: "Custom",
                metricName: "identifier2"
              }
            ]
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}

#Sample template for Grafana Loki Log Health Source
resource "harness_platform_monitored_service" "example4" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "Test"
      identifier = "Test"
      type       = "GrafanaLokiLogs"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "Demo"
            identifier = "Demo"
            query      = "{job=~\".+\"}"
            groupName  = "Log_Group"
            queryParams = {
              serviceInstanceField = "job"
            }
          },
          {
            name       = "log2"
            identifier = "identifier2"
            groupName  = "g2"
            query      = "error"
            queryParams = {
              serviceInstanceField = "_sourcehost"
            }
            liveMonitoringEnabled         = "false"
            continuousVerificationEnabled = "false"
            sliEnabled                    = "false"
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}

#Sample template for Azure Metrics Health Source
resource "harness_platform_monitored_service" "example5" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "azure metrics verify step"
      identifier = "azure_metrics_verify_step"
      type       = "AzureMetrics"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "metric"
            identifier = "metric"
            query      = "default"
            groupName  = "g1"
            queryParams = {
              serviceInstanceField        = "host"
              index                       = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test",
              healthSourceMetricName      = "cpuUsagePercentage",
              healthSourceMetricNamespace = "insights.container/nodes",
              aggregationType             = "average"
            }
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            liveMonitoringEnabled         = "true"
            continuousVerificationEnabled = "true"
            sliEnabled                    = "false"
            metricThresholds: [
              {
                type: "IgnoreThreshold",
                spec: {
                  action: "Ignore"
                },
                criteria: {
                  type: "Absolute",
                  spec: {
                    greaterThan: 100
                  }
                },
                metricType: "Custom",
                metricName: "metric"
              }]
          },
          {
            name       = "name2"
            identifier = "identifier2"
            groupName  = "g2"
            queryParams = {
              serviceInstanceField        = "host"
              index                       = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test",
              healthSourceMetricName      = "cpuUsagePercentage",
              healthSourceMetricNamespace = "insights.container/nodes",
              aggregationType             = "average"
            }
            riskProfile = {
              riskCategory = "Performance_Other"
              thresholdTypes = [
                "ACT_WHEN_HIGHER"
              ]
            }
            liveMonitoringEnabled         = "false"
            continuousVerificationEnabled = "false"
            sliEnabled                    = "false"
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}
#Sample template for Azure Log Health Source
resource "harness_platform_monitored_service" "example6" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
    type            = "Application"
    description     = "description"
    service_ref     = "service_ref"
    environment_ref = "environment_ref"
    tags            = ["foo:bar", "bar:foo"]
    health_sources {
      name       = "Demo azure"
      identifier = "Demo_azure"
      type       = "AzureLogs"
      version    = "v2"
      spec = jsonencode({
        connectorRef = "connectorRef"
        queryDefinitions = [
          {
            name       = "name2"
            identifier = "identifier2"
            groupName  = "g2"
            query      = "*"
            queryParams = {
              serviceInstanceField = "Name",
              timeStampIdentifier  = "StartedTime",
              messageIdentifier    = "Image",
              index                = "/subscriptions/12d2db62-5aa9-471d-84bb-faa489b3e319/resourceGroups/srm-test/providers/Microsoft.ContainerService/managedClusters/srm-test"
            }
            liveMonitoringEnabled         = "false"
            continuousVerificationEnabled = "false"
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}
#Sample template for Prometheus Metrics Health Source
resource "harness_platform_monitored_service" "example7" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
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
                serviceInstanceFieldName = "pod_name"
              }
            }
            sli : {
              enabled = true
            }
            query         = "count(up{group=\"cv\",group=\"cv\"})"
            groupName     = "met"
            isManualQuery = true
          }
        ]
        metricPacks: [
          {
            identifier: "Custom",
            metricThresholds: [
              {
                type: "IgnoreThreshold",
                spec: {
                  action: "Ignore"
                },
                criteria: {
                  type: "Absolute",
                  spec: {
                    greaterThan: 100
                  }
                },
                metricType: "Custom",
                metricName: "Prometheus Metric"
              }
            ]
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}
#Sample template for Datadog Metrics Health Source
resource "harness_platform_monitored_service" "example8" {
  org_id     = "org_id"
  project_id = "project_id"
  identifier = "identifier"
  request {
    name            = "name"
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
        feature      = "Datadog Cloud Metrics"
        metricDefinitions = [
          {
            metricName            = "metric"
            metricPath            = "M1"
            identifier            = "metric"
            query                 = "avg:kubernetes.cpu.limits{*}.rollup(avg, 60);\navg:kubernetes.cpu.limits{*}.rollup(avg, 30);\n(a+b)/10"
            isManualQuery         = true
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
                serviceInstanceFieldName = "pod"
              }
            }
            sli : {
              enabled = true
            }
          },
          {
            metricName            = "dashboard_metric_cpu"
            identifier            = "metric_cpu"
            query                 = "avg:kubernetes.cpu.limits{*}.rollup(avg, 60);\navg:kubernetes.cpu.limits{*}.rollup(avg, 30);\n(a+b)/10"
            isManualQuery         = false
            dashboardName         = "dashboard"
            metricPath            = "M1"
            groupingQuery         = "avg:kubernetes.cpu.limits{*} by {host}.rollup(avg, 60)"
            metric                = "kubernetes.cpu.limits"
            aggregation           = "avg"
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
                serviceInstanceFieldName = "pod"
              }
            }
            sli : {
              enabled = true
            }
          }
        ]
        metricPacks: [
          {
            identifier: "Custom",
            metricThresholds: [
              {
                type: "IgnoreThreshold",
                spec: {
                  action: "Ignore"
                },
                criteria: {
                  type: "Absolute",
                  spec: {
                    greaterThan: 100
                  }
                },
                metricType: "Custom",
                metricName: "metric"
              }
            ]
          }
        ]
      })
    }
    template_ref  = "template_ref"
    version_label = "version_label"
  }
}