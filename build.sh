#!/bin/sh
version=0.37.13
source=registry.terraform.io/harness/harness
platform=darwin_arm64
mkdir -p ~/.terraform.d/plugins/$source/$version/$platform/
echo "HI"
cp terraform-provider-harness ~/.terraform.d/plugins/$source/$version/$platform/terraform-provider-harness
echo "Done"