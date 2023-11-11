data "oci_identity_availability_domains" "AvailabilityDomains" {
  compartment_id = var.compartment_ocid
}

locals {
  availability_domain_id = data.oci_identity_availability_domains.AvailabilityDomains.availability_domains["${var.ad_number}" - 1]["name"]
}

