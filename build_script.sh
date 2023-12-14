#!/bin/sh

version=8.7.72
source=registry.terraform.io/harness/harness
platform=darwin_arm64

mkdir -p ~/.terraform.d/plugins/$source/$version/$platform/ 
cp terraform-provider-harness ~/.terraform.d/plugins/$source/$version/$platform/terraform-provider-harness