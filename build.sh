#!/bin/sh
version=0.90.12
source=registry.terraform.io/harness/harness
platform=darwin_amd64
mkdir -p ~/.terraform.d/plugins/$source/$version/$platform/
echo "HI"
cp terraform-provider-harness ~/.terraform.d/plugins/$source/$version/$platform/terraform-provider-harness
echo "Done"