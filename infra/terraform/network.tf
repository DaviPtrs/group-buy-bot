data "oci_core_vcn" "vcn" {
  vcn_id = var.vcn_id
}

data "oci_core_subnet" "subnet" {
  subnet_id = var.subnet_id
}

resource "oci_core_network_security_group" "nsg" {
  compartment_id = var.compartment_ocid

  display_name = "group-buy-bot_NSG"
  vcn_id       = data.oci_core_vcn.vcn.id
}

resource "oci_core_network_security_group_security_rule" "ssh-rule" {
  description               = "Allow SSH access"
  direction                 = "INGRESS"
  network_security_group_id = oci_core_network_security_group.nsg.id
  protocol                  = "6"
  source                    = "0.0.0.0/0"
  source_type               = "CIDR_BLOCK"

  tcp_options {
    destination_port_range {
      max = "22"
      min = "22"
    }
  }
}
