# Example of a Virtual Registry
resource "harness_platform_har_registry" "virtual_registry" {
  identifier    = "virtual_docker_registry"
  description   = "Virtual Docker Registry"
  space_ref     = "accountId/orgId/projectId"
  package_type  = "DOCKER"

  config {
    type = "VIRTUAL"
    upstream_proxies = ["registry1", "registry2"]
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Registry with Authentication
resource "harness_platform_har_registry" "upstream_registry" {
  identifier    = "upstream_helm_registry"
  description   = "Upstream Helm Registry"
  space_ref     = "accountId/orgId/projectId"
  package_type  = "HELM"

  config {
    type   = "UPSTREAM"
    source = "Custom"
    url    = "https://helm.sh"
    auth {
      auth_type         = "UserPassword"
      user_name         = "registry_user"
      secret_identifier = "registry_password"
      secret_space_path = "accountId/orgId/projectId"
    }
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Go Registry (GoProxy source needs no url)
resource "harness_platform_har_registry" "go_upstream" {
  identifier   = "upstream_go_registry"
  description  = "Upstream Go Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "GO"

  config {
    type      = "UPSTREAM"
    source    = "GoProxy"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Conda Registry (Anaconda source needs no url)
resource "harness_platform_har_registry" "conda_upstream" {
  identifier   = "upstream_conda_registry"
  description  = "Upstream Conda Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "CONDA"

  config {
    type      = "UPSTREAM"
    source    = "Anaconda"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Helm HTTP Registry (HelmChartRepo source requires url)
resource "harness_platform_har_registry" "helm_http_upstream" {
  identifier   = "upstream_helm_http_registry"
  description  = "Upstream Helm HTTP Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "HELM_HTTP"

  config {
    type      = "UPSTREAM"
    source    = "HelmChartRepo"
    url       = "https://charts.bitnami.com/bitnami"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Debian Registry with Debian-specific configuration
resource "harness_platform_har_registry" "debian_virtual" {
  identifier   = "virtual_debian_registry"
  description  = "Virtual Debian Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "DEBIAN"

  config {
    type              = "VIRTUAL"
    upstream_proxies  = ["debian_upstream_registry"]
    debian_config {
      remote_indexed_architectures       = ["amd64", "arm64"]
      optional_index_compression_formats = [".xz"]
    }
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Custom Debian Registry (Debian source needs url)
resource "harness_platform_har_registry" "debian_upstream" {
  identifier   = "upstream_debian_registry"
  description  = "Upstream Debian Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "DEBIAN"

  config {
    type      = "UPSTREAM"
    source    = "Custom"
    url       = "http://deb.debian.org/debian"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Python Registry with a custom remote URL suffix
resource "harness_platform_har_registry" "python_upstream" {
  identifier   = "upstream_python_registry"
  description  = "Upstream Python Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "PYTHON"

  config {
    type              = "UPSTREAM"
    source            = "Custom"
    url               = "https://pypi.example.com"
    remote_url_suffix = "simple"
    auth_type         = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Conan Registry (ConanCenter source needs no url)
resource "harness_platform_har_registry" "conan_upstream" {
  identifier   = "upstream_conan_registry"
  description  = "Upstream Conan Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "CONAN"

  config {
    type      = "UPSTREAM"
    source    = "ConanCenter"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Ruby Registry (RubyGems source needs no url)
resource "harness_platform_har_registry" "rubygems_upstream" {
  identifier   = "upstream_ruby_registry"
  description  = "Upstream Ruby Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "RUBY"

  config {
    type      = "UPSTREAM"
    source    = "RubyGems"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Terraform Registry
resource "harness_platform_har_registry" "terraform_virtual" {
  identifier   = "virtual_terraform_registry"
  description  = "Virtual Terraform Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "TERRAFORM"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual CRAN Registry
resource "harness_platform_har_registry" "cran_virtual" {
  identifier   = "virtual_cran_registry"
  description  = "Virtual CRAN Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "CRAN"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream CRAN Registry (CRAN source needs no url)
resource "harness_platform_har_registry" "cran_upstream" {
  identifier   = "upstream_cran_registry"
  description  = "Upstream CRAN Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "CRAN"

  config {
    type      = "UPSTREAM"
    source    = "CRAN"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Alpine Registry
resource "harness_platform_har_registry" "alpine_virtual" {
  identifier   = "virtual_alpine_registry"
  description  = "Virtual Alpine Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "ALPINE"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Alpine Registry (Alpine source needs no url)
resource "harness_platform_har_registry" "alpine_upstream" {
  identifier   = "upstream_alpine_registry"
  description  = "Upstream Alpine Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "ALPINE"

  config {
    type      = "UPSTREAM"
    source    = "Alpine"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Wolfi Registry
resource "harness_platform_har_registry" "wolfi_virtual" {
  identifier   = "virtual_wolfi_registry"
  description  = "Virtual Wolfi Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "WOLFI"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream Wolfi Registry (Wolfi source needs no url)
resource "harness_platform_har_registry" "wolfi_upstream" {
  identifier   = "upstream_wolfi_registry"
  description  = "Upstream Wolfi Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "WOLFI"

  config {
    type      = "UPSTREAM"
    source    = "Wolfi"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Composer Registry
resource "harness_platform_har_registry" "composer_virtual" {
  identifier   = "virtual_composer_registry"
  description  = "Virtual Composer Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "COMPOSER"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Dart Registry
resource "harness_platform_har_registry" "dart_virtual" {
  identifier   = "virtual_dart_registry"
  description  = "Virtual Dart Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "DART"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Swift Registry
resource "harness_platform_har_registry" "swift_virtual" {
  identifier   = "virtual_swift_registry"
  description  = "Virtual Swift Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "SWIFT"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual APM Registry
resource "harness_platform_har_registry" "apm_virtual" {
  identifier   = "virtual_apm_registry"
  description  = "Virtual APM Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "APM"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of an Upstream APM Registry
# APM has no built-in default source - use Custom with the upstream APM registry url.
resource "harness_platform_har_registry" "apm_upstream" {
  identifier   = "upstream_apm_registry"
  description  = "Upstream APM Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "APM"

  config {
    type      = "UPSTREAM"
    source    = "Custom"
    url       = "https://apm.example.com"
    auth_type = "Anonymous"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Hugging Face Registry
# HUGGINGFACE supports VIRTUAL registries only - upstream is not supported.
resource "harness_platform_har_registry" "huggingface_virtual" {
  identifier   = "virtual_huggingface_registry"
  description  = "Virtual Hugging Face Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "HUGGINGFACE"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}

# Example of a Virtual Terraform Backend Registry
# TERRAFORM_BACKEND supports VIRTUAL registries only - upstream is not supported.
resource "harness_platform_har_registry" "terraform_backend_virtual" {
  identifier   = "virtual_terraform_backend_registry"
  description  = "Virtual Terraform Backend Registry"
  space_ref    = "accountId/orgId/projectId"
  package_type = "TERRAFORM_BACKEND"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "accountId/orgId/projectId"
}
