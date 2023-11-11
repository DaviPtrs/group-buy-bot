variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "tenancy_ocid" {}

variable "compartment_ocid" {
}
variable "region" {
}

variable "vcn_id" {
}
variable "subnet_id" {
}

variable "ad_number" {
  type    = string
  default = "1"
}
variable "ssh_pub_key" {
}
variable "instance_shape" {
}
variable "instance_ocpus" {
}
variable "instance_memory_in_gbs" {
}
variable "instance_image_ocid" {
}
variable "preserve_boot_volume" {
  type    = bool
  default = false
}

variable "datadisk_device_path" {
  default = "/dev/oracleoci/oraclevdc"
}
variable "datadisk_size_in_gbs" {
  default = "100"
}