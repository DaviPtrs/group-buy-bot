data "oci_identity_availability_domains" "AvailabilityDomains" {
  compartment_id = var.compartment_ocid
}

locals {
  availability_domain_id = data.oci_identity_availability_domains.AvailabilityDomains.availability_domains["${var.ad_number}" - 1]["name"]
}

resource "oci_core_instance" "instance" {
  availability_domain = local.availability_domain_id
  compartment_id      = var.compartment_ocid
  display_name        = "group-buy-bot"
  create_vnic_details {
    assign_public_ip       = "true"
    display_name           = "group-buy-bot"
    hostname_label         = "group-buy-bot"
    skip_source_dest_check = "true"
    subnet_id              = data.oci_core_subnet.subnet.id
    nsg_ids = [
      oci_core_network_security_group.nsg.id
    ]
  }

  metadata = {
    "ssh_authorized_keys" = var.ssh_pub_key
  }

  shape = var.instance_shape
  shape_config {
    ocpus         = var.instance_ocpus
    memory_in_gbs = var.instance_memory_in_gbs
  }
  source_details {
    source_id   = var.instance_image_ocid
    source_type = "image"
  }
  preserve_boot_volume = var.preserve_boot_volume
}