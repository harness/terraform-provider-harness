# Example of a Kubernetes Security Governance Condition
resource "harness_chaos_security_governance_condition" "k8s_condition" {
  org_id      = var.org_id
  project_id  = var.project_id
  name        = "k8s-security-condition"
  description = "Security governance condition for Kubernetes workloads"
  infra_type  = "KubernetesV2"
  
  fault_spec {
    operator = "NOT_EQUAL_TO"
    
    faults {
      fault_type = "FAULT"
      name       = "pod-delete"
    }
    
    faults {
      fault_type = "FAULT"
      name       = "pod-dns"
    }
  }
  
  k8s_spec {
    infra_spec {
      operator  = "EQUAL_TO"
      infra_ids = [var.k8s_infra_id]
    }
    
    application_spec {
      operator = "EQUAL_TO"
      
      workloads {
        namespace = "default"
        kind      = "deployment"
        label     = "app=nginx"
        services  = ["nginx-service"]
        application_map_id = "nginx-app"
      }
    }
    
    chaos_service_account_spec {
      operator = "EQUAL_TO"
      service_accounts = ["default", "chaos-service-account"]
    }
  }
  
  tags = ["env:prod", "team:security", "platform:k8s"]
}

# Example of a Windows Security Governance Condition
resource "harness_chaos_security_governance_condition" "windows_condition" {
  org_id      = var.org_id
  project_id  = var.project_id
  name        = "windows-security-condition"
  description = "Security governance condition for Windows hosts"
  infra_type  = "Windows"
  
  fault_spec {
    operator = "NOT_EQUAL_TO"
    
    faults {
      fault_type = "FAULT"
      name       = "process-kill"
    }
    
    faults {
      fault_type = "FAULT"
      name       = "cpu-hog"
    }
  }
  
  machine_spec {
    infra_spec {
      operator  = "EQUAL_TO"
      infra_ids = [var.windows_infra_id]
    }
  }
  
  tags = ["env:prod", "team:security", "platform:windows"]
}

# Example of a Linux Security Governance Condition
resource "harness_chaos_security_governance_condition" "linux_condition" {
  org_id      = var.org_id
  project_id  = var.project_id
  name        = "linux-security-condition"
  description = "Security governance condition for Linux hosts"
  infra_type  = "Linux"
  
  fault_spec {
    operator = "NOT_EQUAL_TO"
    
    faults {
      fault_type = "FAULT"
      name       = "process-kill"
    }
    
    faults {
      fault_type = "FAULT"
      name       = "memory-hog"
    }
  }
  
  machine_spec {
    infra_spec {
      operator  = "EQUAL_TO"
      infra_ids = [var.linux_infra_id]
    }
  }
  
  tags = ["env:prod", "team:security", "platform:linux"]
}

# Output the created conditions
output "k8s_condition_id" {
  value = harness_chaos_security_governance_condition.k8s_condition.id
}

output "windows_condition_id" {
  value = harness_chaos_security_governance_condition.windows_condition.id
}

output "linux_condition_id" {
  value = harness_chaos_security_governance_condition.linux_condition.id
}
