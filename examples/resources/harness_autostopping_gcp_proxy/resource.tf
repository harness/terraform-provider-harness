resource "harness_autostopping_gcp_proxy" "test" {
  name               = "name"
  cloud_connector_id = "cloud_connector_id"
  host_name          = "host_name"
  region             = "region"
  vpc                = "https://www.googleapis.com/compute/v1/projects/project_id/global/networks/netwok_id"
  zone               = "zone"
  security_groups    = ["http-server"]
  machine_type       = "e2-micro"
  subnet_id          = "https://www.googleapis.com/compute/v1/projects/project_id/regions/region/subnetworks/subnet_name"
  api_key            = ""
  allocate_static_ip = false
  certificates {
    key_secret_id  = "projects/project_id/secrets/secret_id/versions/1"
    cert_secret_id = "projects/project_id/secrets/secret_id/versions/1"
  }
  delete_cloud_resources_on_destroy = false
}

