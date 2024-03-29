---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_autostopping_rule_vm Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness Variable.
---

# harness_autostopping_rule_vm (Data Source)

Data source for retrieving a Harness Variable.

## Example Usage

```terraform
data "harness_autostopping_rule_vm" "example" {
  identifier = "identifier"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_connector_id` (String) Id of the cloud connector
- `filter` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--filter))
- `name` (String) Name of the rule

### Optional

- `custom_domains` (List of String) Custom URLs used to access the instances
- `depends` (Block List) Dependent rules (see [below for nested schema](#nestedblock--depends))
- `http` (Block List) Http routing configuration (see [below for nested schema](#nestedblock--http))
- `idle_time_mins` (Number) Idle time in minutes. This is the time that the AutoStopping rule waits before stopping the idle instances.
- `tcp` (Block List) TCP routing configuration (see [below for nested schema](#nestedblock--tcp))
- `use_spot` (Boolean) Boolean that indicates whether the selected instances should be converted to spot vm

### Read-Only

- `id` (String) The ID of this resource.
- `identifier` (Number) Unique identifier of the resource

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `vm_ids` (List of String) Ids of instances that needs to be managed using the AutoStopping rules

Optional:

- `regions` (List of String) Regions of instances that needs to be managed using the AutoStopping rules
- `tags` (Block List) Tags of instances that needs to be managed using the AutoStopping rules (see [below for nested schema](#nestedblock--filter--tags))
- `zones` (List of String) Zones of instances that needs to be managed using the AutoStopping rules

<a id="nestedblock--filter--tags"></a>
### Nested Schema for `filter.tags`

Required:

- `key` (String)
- `value` (String)



<a id="nestedblock--depends"></a>
### Nested Schema for `depends`

Required:

- `rule_id` (Number) Rule id of the dependent rule

Optional:

- `delay_in_sec` (Number) Number of seconds the rule should wait after warming up the dependent rule


<a id="nestedblock--http"></a>
### Nested Schema for `http`

Required:

- `proxy_id` (String) Id of the proxy

Optional:

- `health` (Block List) Health Check Details (see [below for nested schema](#nestedblock--http--health))
- `routing` (Block List) Routing configuration used to access the instances (see [below for nested schema](#nestedblock--http--routing))

<a id="nestedblock--http--health"></a>
### Nested Schema for `http.health`

Required:

- `port` (Number) Health check port on the VM
- `protocol` (String) Protocol can be http or https

Optional:

- `path` (String) API path to use for health check
- `status_code_from` (Number) Lower limit for acceptable status code
- `status_code_to` (Number) Upper limit for acceptable status code
- `timeout` (Number) Health check timeout


<a id="nestedblock--http--routing"></a>
### Nested Schema for `http.routing`

Required:

- `source_protocol` (String) Source protocol of the proxy can be http or https
- `target_protocol` (String) Target protocol of the instance can be http or https

Optional:

- `action` (String) Organization Identifier for the Entity
- `source_port` (Number) Port on the proxy
- `target_port` (Number) Port on the VM



<a id="nestedblock--tcp"></a>
### Nested Schema for `tcp`

Required:

- `proxy_id` (String) Id of the Proxy

Optional:

- `forward_rule` (Block List) Additional tcp forwarding rules (see [below for nested schema](#nestedblock--tcp--forward_rule))
- `rdp` (Block List) RDP configuration (see [below for nested schema](#nestedblock--tcp--rdp))
- `ssh` (Block List) SSH configuration (see [below for nested schema](#nestedblock--tcp--ssh))

<a id="nestedblock--tcp--forward_rule"></a>
### Nested Schema for `tcp.forward_rule`

Required:

- `port` (Number) Port to listen on the vm

Optional:

- `connect_on` (Number) Port to listen on the proxy


<a id="nestedblock--tcp--rdp"></a>
### Nested Schema for `tcp.rdp`

Optional:

- `connect_on` (Number) Port to listen on the proxy
- `port` (Number) Port to listen on the vm


<a id="nestedblock--tcp--ssh"></a>
### Nested Schema for `tcp.ssh`

Optional:

- `connect_on` (Number) Port to listen on the proxy
- `port` (Number) Port to listen on the vm
