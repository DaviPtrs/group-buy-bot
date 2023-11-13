terraform {
  cloud {
    organization = "DaviPtrs"
    workspaces {
      project = "group-buy-bot"
      name    = "main"
    }
  }
}
provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid    = var.user_ocid
  fingerprint  = var.fingerprint
  private_key  = var.private_key
  region       = var.region
}

