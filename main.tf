terraform {
  required_providers {
    cidrblock = {
      source = "amilevskiy/cidrblock"
    }
  }
}

provider "cidrblock" {}

data "cidrblock_summarization" "vpc" {
  cidr_blocks = [
    "10.192.0.0/23",
    "10.192.2.0/22",
    "10.192.6.0/22",
    "10.100.0.0/24",
    "10.100.1.0/24",
    "10.100.4.0/23",
    "10.100.6.0/24",
  ]
}

#expected output
#10.192.0.0/21, 10.100.0.0/23, 10.100.4.0/23, 10.100.6.0/24

#10.192.0.0/21, 10.100.0.0/23, 10.100.4.0/23, 10.100.6.0/24

output "summarized_cidr_blocks" {
  value = data.cidrblock_summarization.vpc.summarized_cidr_blocks
}


