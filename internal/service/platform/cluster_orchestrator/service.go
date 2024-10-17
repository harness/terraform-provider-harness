package cluster_orchestrator

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildClusterOrch(d *schema.ResourceData) nextgen.CreateClusterOrchestratorDto {

	clusterOrch := &nextgen.CreateClusterOrchestratorDto{}

	if attr, ok := d.GetOk("name"); ok {
		clusterOrch.Name = attr.(string)
	}
	if attr, ok := d.GetOk("k8s_connector_id"); ok {
		clusterOrch.K8sConnID = attr.(string)
	}
	userCfg := nextgen.ClusterOrchestratorUserConfig{}

	if attr, ok := d.GetOk("cluster_endpoint"); ok {
		userCfg.ClusterEndPoint = attr.(string)
	}
	clusterOrch.UserConfig = userCfg

	return *clusterOrch

}

func setId(d *schema.ResourceData, id string) {
	d.SetId(id)
	d.Set("identifier", id)
}
