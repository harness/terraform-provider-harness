# Harness Terraform Provider

- Website: [harness.io](https://harness.io)

The Terraform provider for Harness is a plugin that allows you to manage resources in Harness CD.

## Support

If you have any questions please open a [new issue](https://github.com/harness/terraform-provider-harness/issues/new) or join our slack [channel](https://harnesscommunity.slack.com/archives/C02G9CUNF1S).

## Quick Starts

- [Example project](https://github.com/harness/terraform-demo)
- [Provider usage](https://registry.terraform.io/providers/harness/harness/latest/docs)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.17

## Documentation

Full, comprehensive documentation is available on the Terraform website:

<https://registry.terraform.io/providers/harness/harness/latest/docs>

## Building and Testing Locally

1. Clone the repo into your local directory. Run `git clone https://github.com/harness/terraform-provider-harness.git`
2. Run `go mod tidy`
3. Run `go build -o terraform-provider-harness`
4. Create a file called `local.sh` in the root directory of the repository and copy the following script to the bash file 

```SH
#!/bin/sh

version=0.40.2 #specify in this format 
source=registry.terraform.io/harness/harness
platform=darwin_amd64

mkdir -p ~/.terraform.d/plugins/$source/$version/$platform/

cp terraform-provider-harness ~/.terraform.d/plugins/$source/$version/$platform/terraform-provider-harness
```

5. Run the Bash Script `./local.sh`

## Point terraform script to local terraform-provider-harness build
1. Update the .terraform.rc file
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/harness/harness" = "{path}/terraform-provider-harness"
  }
  direct {}
}
```
2. Create build - `go build`
*Note: Please make sure the terraform provider version matches the version in the script*
