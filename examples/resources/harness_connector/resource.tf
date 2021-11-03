# Create a kubernetes connector

resource "harness_connector" "test" {
  identifier = "k8s"
  name       = "Test Kubernetes Cluster"

  k8s_cluster {
    service_account {
      master_url                = "https://kubernetes.example.com"
      service_account_token_ref = "account.k8s_service_account_token"
    }
  }
}

# Create a docker registry connector

resource "harness_connector" "test" {
  identifier = "dockerhub"
  name       = "Docker Hub"

  docker_registry {
    type = "DockerHub"
    url  = "https://hub.docker.io"

    credentials {
      username     = "admin"
      password_ref = "account.docker_registry_password"
    }
  }
}
