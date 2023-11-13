resource "oci_core_volume" "instance-datadisk" {
  availability_domain = local.availability_domain_id
  compartment_id      = var.compartment_ocid
  display_name        = "group-buy-bot_datadisk"
  size_in_gbs         = var.datadisk_size_in_gbs
  vpus_per_gb         = "10"

  lifecycle {
    create_before_destroy = false
  }
}

resource "oci_core_volume_attachment" "instance-datadisk" {
  attachment_type = "paravirtualized"
  device          = var.datadisk_device_path
  instance_id     = oci_core_instance.instance.id
  volume_id       = oci_core_volume.instance-datadisk.id
}

resource "oci_core_volume_backup_policy_assignment" "datadisk_backup" {

  asset_id  = oci_core_volume.instance-datadisk.id
  policy_id = local.backup_policy_map["gold"]
}
