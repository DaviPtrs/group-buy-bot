output "instance_public_ip" {
  value = oci_core_instance.instance.public_ip
}

output "datadisk_path" {
  value = oci_core_volume_attachment.instance-datadisk.device
}