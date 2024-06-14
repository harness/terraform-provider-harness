// Create a git repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "https://github.com/willycoll/argocd-example-apps.git"
    name            = "repo_name"
    insecure        = true
    connection_type = "HTTPS_ANONYMOUS"
  }
  upsert = true
}

// Create a ssh git repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "git@github.com:yourorg"
    name            = "repo_name"
    insecure        = false
    connection_type = "SSH"
    ssh_private_key = "----- BEGIN OPENSSH PRIVATE KEY-----\nXXXXX\nXXXXX\nXXXXX\n-----END OPENSSH PRIVATE KEY -----\n"
  }
  upsert = true
}


// Create a HELM repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "https://charts.helm.sh/stable"
    name            = "repo_name"
    insecure        = true
    connection_type = "HTTPS_ANONYMOUS"
    type_           = "helm"
  }
  upsert = true
}

// Create a OCI HELM repository at project level
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "ghcr.io/test-repo"
    name            = "repo_name"
    insecure        = false
    username        = "username"
    password        = "ghp_xxxxxxxx"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  upsert = true
}
resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "111222333444.dkr.ecr.us-west-1.amazonaws.com"
    name            = "repo_name"
    insecure        = false
    username        = "AWS"
    password        = "aws_ecr_token"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  gen_type = "AWS_ECR"
  ecr_gen {
    region = "us-west-1"
    secret_ref {
      aws_access_key_id     = "AWS_ACCESS_KEY_ID"
      aws_secret_access_key = "AWS_SECRET_ACCESS_KEY"
    }
  }
  refreshInterval = "1m"
  upsert          = false
}

resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "111222333444.dkr.ecr.us-west-1.amazonaws.com"
    name            = "repo_name"
    insecure        = false
    username        = "AWS"
    password        = "aws_ecr_token"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  gen_type = "AWS_ECR"
  ecr_gen {
    region = "us-west-1"
    jwt_auth {
      name      = "name"
      namespace = "namespace"
    }
  }
  refreshInterval = "1m"
  upsert          = false
}

resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "us.gcr.io/projectID/repo_name"
    name            = "repo_name"
    insecure        = false
    username        = "oauth2accesstoken"
    password        = "aws_ecr_token"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  gen_type = "GOOGLE_GCR"
  gcr_gen {
    projectID = "projectID"
    accessKey = "{  \"type\": \"service_account\",  \"project_id\": \"google-project-id\",  \"private_key_id\": \"xxxxxxx19dd12674be7be2312313caaabxxxx\",  \"private_key\": \"-----.......-----\n\",  \"client_email\": \"xxxxxxxx-compute@developer.gserviceaccount.com\",  \"client_id\": \"xxxxxxxxxxx0161940\",  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",  \"token_uri\": \"https://oauth2.googleapis.com/token\",  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/511111111911-compute@developer.gserviceaccount.com\",  \"universe_domain\": \"googleapis.com\"}"
  }
  refreshInterval = "1m"
  upsert          = false
}

resource "harness_platform_gitops_repository" "example" {
  identifier = "identifier"
  account_id = "account_id"
  project_id = "project_id"
  org_id     = "org_id"
  agent_id   = "agent_id"
  repo {
    repo            = "us.gcr.io/projectID/repo_name"
    name            = "repo_name"
    insecure        = false
    username        = "oauth2accesstoken"
    password        = "aws_ecr_token"
    connection_type = "HTTPS"
    type_           = "helm"
    enable_oci      = true
  }
  gen_type = "GOOGLE_GCR"
  gcr_gen {
    project_id = "projectID"
    workload_identity {
      cluster_location   = "GCPClusterLocation"
      cluster_name       = "GCPClusterName"
      cluster_project_id = "GCPClusterProjectID"
      service_account_ref {
        name      = "name"
        namespace = "namespace"
      }
    }
  }
  refreshInterval = "1m"
  upsert          = false
}