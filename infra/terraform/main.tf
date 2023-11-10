# terraform {
#   backend "remote" {
#     organization = "DaviPtrs"
#     workspaces {
#       name = "group-buy-bot"
#     }
#   }
# }


terraform {
  cloud {
    organization = "DaviPtrs"
    workspaces {
      name = "group-buy-bot"
    }
  }
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

