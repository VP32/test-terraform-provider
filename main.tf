terraform {
  required_providers {
    vpfoo = {
      version = "~> 0.2"
      source  = "hashicorp.com/edu/vpfoo"
    }
  }
}

provider "vpfoo" {
  host = "http://localhost:8080"
}

resource "vpfoo_fooserver" "my-server1" {
  items {
    server {
      region = "honduras"
      type   = "automatic"
    }
  }
}

output "items" {
  value = vpfoo_fooserver.my-server1.items
}


data "vpfoo_fooserver" "item" {
  items {
    server {
      name = "N8081TserverRMS"
    }
  }
}

output "found_server" {
  value = data.vpfoo_fooserver.item
}

