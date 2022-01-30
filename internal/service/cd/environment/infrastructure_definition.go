package environment

import (
	"context"
	"fmt"
	"log"
	"strings"

	sdk "github.com/harness-io/harness-go-sdk"
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	DefaultHostnameConvention = "${host.ec2Instance.privateDnsName.split('\\.')[0]}"
	DefaultK8sReleaseName     = "release-${infra.kubernetes.infraId}"
	DefaultHelmReleaseName    = "${infra.kubernetes.infraId}"
)

var infraDetailTypes = []string{
	"kubernetes",
	"kubernetes_gcp",
	"aws_ssh",
	"aws_ami",
	"aws_ecs",
	"aws_lambda",
	"aws_winrm",
	"azure_vmss",
	"azure_webapp",
	"tanzu",
	"datacenter_winrm",
	"datacenter_ssh",
}

func ResourceInfraDefinition() *schema.Resource {

	return &schema.Resource{
		Description:   utils.ConfigAsCodeDescription("Resource for creating am infrastructure definition."),
		CreateContext: resourceInfraDefinitionCreateOrUpdate,
		ReadContext:   resourceInfraDefinitionRead,
		UpdateContext: resourceInfraDefinitionCreateOrUpdate,
		DeleteContext: resourceInfraDefinitionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The unique id of the infrastructure definition.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the infrastructure definition",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"app_id": {
				Description: "The id of the application the infrastructure definition belongs to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"env_id": {
				Description: "The id of the environment the infrastructure definition belongs to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"cloud_provider_type": {
				Description:  fmt.Sprintf("The type of the cloud provider to connect with. Valid options are %s", strings.Join(cac.CloudProviderTypesSlice, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cac.CloudProviderTypesSlice, false),
			},
			"deployment_type": {
				Description:  fmt.Sprintf("The type of the deployment to use. Valid options are %s", strings.Join(cac.DeploymenTypesSlice, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cac.DeploymenTypesSlice, false),
			},
			"provisioner_name": {
				Description: "The name of the infrastructure provisioner to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployment_template_uri": {
				Description: "The URI of the deployment template to use. Only used if deployment_type is `CUSTOM`.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"scoped_services": {
				Description: "The list of service names to scope this infrastructure definition to.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"kubernetes": {
				Description:   "The configuration details for Kubernetes deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsK8sDirectSchema(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "kubernetes"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"kubernetes_gcp": {
				Description:   "The configuration details for Kubernetes on GCP deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsK8sGcp(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "kubernetes_gcp"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"aws_ssh": {
				Description:   "The configuration details for AWS SSH deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAwsSSH(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "aws_ssh"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"aws_ami": {
				Description:   "The configuration details for Aws AMI deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAwsAmi(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "aws_ami"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"aws_ecs": {
				Description:   "The configuration details for Aws AMI deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAwsEcs(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "aws_ecs"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"aws_lambda": {
				Description:   "The configuration details for Aws Lambda deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAwsLambda(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "aws_lambda"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"aws_winrm": {
				Description:   "The configuration details for AWS WinRM deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAwsWinRM(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "aws_winrm"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"azure_vmss": {
				Description:   "The configuration details for Azure VMSS deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAzureVmss(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "azure_vmss"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"azure_webapp": {
				Description:   "The configuration details for Azure WebApp deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsAzureWebApp(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "azure_webapp"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"tanzu": {
				Description:   "The configuration details for PCF deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsTanzu(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "tanzu"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"datacenter_winrm": {
				Description:   "The configuration details for WinRM datacenter deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsDatacenterWinRM(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "datacenter_winrm"),
				ExactlyOneOf:  infraDetailTypes,
			},
			"datacenter_ssh": {
				Description:   "The configuration details for SSH datacenter deployments.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          infraDetailsDatacenterSSH(),
				ConflictsWith: utils.GetConflictsWithSlice(infraDetailTypes, "datacenter_ssh"),
				ExactlyOneOf:  infraDetailTypes,
			},
		},
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				// <app_id>/<env_id>/<id>
				parts := strings.Split(d.Id(), "/")

				d.Set("app_id", parts[0])
				d.Set("env_id", parts[1])
				d.SetId(parts[2])

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceInfraDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)
	envId := d.Get("env_id").(string)

	log.Printf("[DEBUG] Terraform: Read infrastructure definition %s", id)
	infraDef, err := c.CDClient.ConfigAsCodeClient.GetInfraDefinitionById(appId, envId, id)
	if err != nil {
		return diag.FromErr(err)
	} else if infraDef == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readInfraDefinition(d, infraDef)

	return nil
}

func resourceInfraDefinitionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)
	envId := d.Get("env_id").(string)

	log.Printf("[DEBUG] Terraform: Delete infrastructure definition %s", id)
	err := c.CDClient.ConfigAsCodeClient.DeleteInfraDefinition(appId, envId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceInfraDefinitionCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*sdk.Session)

	var input *cac.InfrastructureDefinition
	var err error

	if d.IsNewResource() {
		log.Printf("[DEBUG] Terraform: Create infrastructure definition %s", d.Get("name"))
		input = cac.NewEntity(cac.ObjectTypes.InfrastructureDefinition).(*cac.InfrastructureDefinition)
	} else {
		id := d.Get("id").(string)
		appId := d.Get("app_id").(string)
		envId := d.Get("env_id").(string)
		log.Printf("[DEBUG] Terraform: Updating infrastructure definition %s", d.Get("name"))
		input, err = c.CDClient.ConfigAsCodeClient.GetInfraDefinitionById(appId, envId, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if attr := d.Get("app_id"); attr != "" {
		input.ApplicationId = attr.(string)
	}

	if attr := d.Get("env_id"); attr != "" {
		input.EnvironmentId = attr.(string)
	}

	if attr := d.Get("name"); attr != "" {
		input.Name = attr.(string)
	}

	if attr := d.Get("cloud_provider_type"); attr != "" {
		input.CloudProviderType = cac.CloudProviderType(attr.(string))
	}

	if attr := d.Get("deployment_type"); attr != "" {
		input.DeploymentType = cac.DeploymentType(attr.(string))
	}

	if attr := d.Get("provisioner_name"); attr != "" {
		input.Provisioner = attr.(string)
	}

	if attr := d.Get("deployment_template_uri"); attr != "" {
		input.DeploymentTypeTemplateUri = attr.(string)
	}

	expandScopedServices(d.Get("scoped_services").(*schema.Set).List(), input)
	expandKubernetesConfiguration(d.Get("kubernetes").([]interface{}), input)
	expandKubernetesGcpConfiguration(d.Get("kubernetes_gcp").([]interface{}), input)
	expandAwsSSHConfiguration(d.Get("aws_ssh").([]interface{}), input)
	expandAwsAmiConfiguration(d.Get("aws_ami").([]interface{}), input)
	expandAwsEcsConfiguration(d.Get("aws_ecs").([]interface{}), input)
	expandAwsLambdaConfiguration(d.Get("aws_lambda").([]interface{}), input)
	expandAwsWinRMConfiguration(d.Get("aws_winrm").([]interface{}), input)
	expandTanzuConfiguration(d.Get("tanzu").([]interface{}), input)
	expandAzureWebAppConfiguration(d.Get("azure_webapp").([]interface{}), input)

	infraDef, err := c.CDClient.ConfigAsCodeClient.UpsertInfraDefinition(input)
	if err != nil {
		return diag.FromErr(err)
	}

	readInfraDefinition(d, infraDef)

	return nil
}

func readInfraDefinition(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) {
	d.SetId(infraDef.Id)
	d.Set("app_id", infraDef.ApplicationId)
	d.Set("env_id", infraDef.EnvironmentId)
	d.Set("name", infraDef.Name)
	d.Set("cloud_provider_type", infraDef.CloudProviderType)
	d.Set("deployment_type", infraDef.DeploymentType)
	d.Set("provisioner_name", infraDef.Provisioner)
	d.Set("deployment_template_uri", infraDef.DeploymentTypeTemplateUri)

	if services := flattenScopedServices(d, infraDef); len(services) > 0 {
		d.Set("scoped_services", services)
	}

	if config := flattenKubernetesConfiguration(d, infraDef); len(config) > 0 {
		d.Set("kubernetes", config)
	}

	if config := flattenKubernetesGcpConfiguration(d, infraDef); len(config) > 0 {
		d.Set("kubernetes_gcp", config)
	}

	if config := flattenAwsSSHConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ssh", config)
	}

	if config := flattenAwsAmiConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ami", config)
	}

	if config := flattenAwsEcsConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_ecs", config)
	}

	if config := flattenAwsLambdaConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_lambda", config)
	}

	if config := flattenAwsWinRMConfiguration(d, infraDef); len(config) > 0 {
		d.Set("aws_winrm", config)
	}

	if config := flattenTanzuConfiguration(d, infraDef); len(config) > 0 {
		d.Set("tanzu", config)
	}

	if config := flattenAzureWebAppConfiguration(d, infraDef); len(config) > 0 {
		d.Set("azure_webapp", config)
	}

}

func flattenScopedServices(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	for _, v := range infraDef.ScopedServices {
		results = append(results, v)
	}

	return results
}

func expandScopedServices(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	results := []string{}

	for _, v := range d {
		results = append(results, v.(string))
	}

	infraDef.ScopedServices = results
}
