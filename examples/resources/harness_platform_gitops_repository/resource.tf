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
    accessKey = "{  \"type\": \"service_account\",  \"project_id\": \"google-project-id\",  \"private_key_id\": \"5ef370c719dd12674be7be2312313caaab31231\",  \"private_key\": \"-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDQXwd7zfou3NDB\nP9TwgQMWyHYpWRJENC7RTpCaQ1KqBElEHPUBzPpv2L/ayn50qpsmtxYTWdQwR7OL\n8qK+8hrYfSvwSTph/PMOtesrZpWrloIG2virgw1y1zeUDP3DKQDemerVrpve189h\nxxTLTriLS0JMsDniE0HHzeeMtxTbi0NcojDufSq3gTZ4TJ80jPtdIURYAwfW4jHA\nOW1nq8EY3VjOkUDHZ++xtHen9FT4OzpcTSkGyovaYZfLUPFDrtZqmcdS5IjFmTpE\nHfIibB+roT0jynhBBPwuBUWtbJEjXg8Gw2hRquSChCMxKkD9PjttyQMjmlhVrLqR\nG1p4vBSDAgMBAAECggEAS8nUsf4oOjVMpI1wCQ4Troy5Fa71CuOkB7M4uzMzdO1c\nLK8PmlkQ2e+PUKgIOKz5A6riF6W7nNfngUZ+VU8/3nAgtCQeXReg3D/kyoNkeuWi\nY5Xvjop7MMMAzxOulPZr/4siNBhvTy1Vm63KbWwziU6VTclnNEhmy6KjzrWkm3ky\ndcYsnGr5eWbQvmSAE38EeSMw6+OoiEmPYk/hoRg85lVouQ72d4FHaNjU3NNCR1Y0\nUEdj8r9K19nZ7ICTxJ6AZp5rc3Z2rG8MrBY2UEhGhFLZmad8hUtyo5Ol93kOwSHW\n+baGXRasBhWN3uZ3PYCh3NzEJeHEVX0HT4FYUF8NQQKBgQD1OMOVAF9XWESPTuyw\nz5/4S+kRWdplXkP9dEacQ21hIwgX5PFWZFHD1VhwSVniIPWao2Fa/QMZK4np+g9d\nnEkgPlFatPsLT0q3/QT59oHEIAIEorOdz0RXxAkU9xo0RfXQWsXFDgcwcKd9yeet\nxbiQO/LscNomM/CcWW63O1l5IwKBgQDZh596QU9Y3z07OfF9pl86X+QIQlEY0nxr\nx2L+JspVXWnIHoVGlODOoP/EmCfS23oJdZZC7TWLSS9GDCsTC4UPcHW5I0cFFT69\n9M0ZvP2P6oCf2Jg7QOX8DIamcv6wI0MQKdUFDW+wtf01hiS/6lwEQL8xFBhw2+xq\njIKdkoOdIQKBgQCN6Z7OURvb6Xor0UoK/O0f/ZZQ80X/mfEQ8cSXVDItn99kLJs6\nGu5yvbnjqZ95zQc1yc1iob+0Rk0W+h8AVpy/KzFbpBcQsX+VQLkri2wHu1pPonT+\nI9/yRsHWvzYMAFzEinOfmYGxl9BmbH1GRIGN/xOTn6+voilh4iO/qHocLwKBgCNy\n7pJFwmCBQME+GBSZ4DrrFYYjCIQ7CPunaoJwX9i5eFucXau650fFBOlMwnCiQ6j2\n+J2/elJQgtuvb/WSdqSJFyYskY5KgAcEtcfT/J5PYNarvWMqmFAS2n6Vjtu1Y2Bm\n8Mf6AJGTlsf6LFL6JjSrOH0PAUyjCkvyyfZTwgw3BAoGBAOrOYrOC6zigjC5Kmve3\nORnw318hPOV5oo7a7NpztSwwY1/7xZuOJZLaflZXnYCO1BXY+PosshI1cdfrv6PT\niEr+SQ+mbaaxcFxtJUP6Y4GBI4ayeHnmqafuVwPEd//rnPD6YA5RRFF/dfI619Hu\nAt9fAayERhb7iptxMQw6wpbF\n-----END PRIVATE KEY-----\n\",  \"client_email\": \"xxxxxxxx-compute@developer.gserviceaccount.com\",  \"client_id\": \"xxxxxxxxxxx0161940\",  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",  \"token_uri\": \"https://oauth2.googleapis.com/token\",  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/511111111911-compute@developer.gserviceaccount.com\",  \"universe_domain\": \"googleapis.com\"}"
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