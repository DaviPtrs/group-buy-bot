data "oci_identity_availability_domains" "AvailabilityDomains" {
  compartment_id = var.compartment_ocid
}

data "oci_core_volume_backup_policies" "volume_backup_policies" {
}


locals {
  availability_domain_id = data.oci_identity_availability_domains.AvailabilityDomains.availability_domains["${var.ad_number}" - 1]["name"]
  backup_policy_map = {
    for policy in data.oci_core_volume_backup_policies.volume_backup_policies.volume_backup_policies :
    policy.display_name => policy.id
  }
}

