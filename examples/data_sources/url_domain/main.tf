terraform {
  required_providers {
    url = {
      source = "registry.terraform.io/gmeligio/url"
    }
  }
}

provider "url" {}

data "url_domain" "example" {
  host = "example.com"
}

output "domain" {
  value = data.url_domain.example
}