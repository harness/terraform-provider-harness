package cluster_orchestrator

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
func buildClusterOrchConfig(d *schema.ResourceData) nextgen.ClusterOrchConfig {
	config := &nextgen.ClusterOrchConfig{}
	if attr, ok := d.GetOk("distribution.0.strategy"); ok {
		config.DistributionStrategy = nextgen.ClusterOrchNodeDistributionStrategy(attr.(string))
	}
	if attr, ok := d.GetOk("distribution.0.base_ondemand_capacity"); ok {
		config.BaseOnDemandCapacity = attr.(int)
	}

	attr := d.Get("distribution.0.ondemand_replica_percentage").(float64)
	config.OnDemandSplit = int(attr)
	config.SpotSplit = 100 - config.OnDemandSplit

	if attr, ok := d.GetOk("distribution.0.selector"); ok {
		config.SpotDistribution = nextgen.ClusterOrchDistributionSelector(attr.(string))
	}
	if _, ok := d.GetOk("binpacking.0.pod_eviction"); ok {
		config.Consolidation.PodEvictor.Enabled = true
		if attr, ok := d.GetOk("binpacking.0.pod_eviction.0.threshold.0.cpu"); ok {
			config.Consolidation.PodEvictor.MinCPU = attr.(float64)
		}
		if attr, ok := d.GetOk("binpacking.0.pod_eviction.0.threshold.0.memory"); ok {
			config.Consolidation.PodEvictor.MinMem = attr.(float64)
		}
	}
	if attr, ok := d.GetOk("binpacking.0.disruption.0.criteria"); ok {
		config.Consolidation.ConsolidationPolicy = nextgen.ConsolidationPolicy(attr.(string))
	}
	if attr, ok := d.GetOk("binpacking.0.disruption.0.delay"); ok {
		config.Consolidation.ConsolidateAfter = attr.(string)
	}
	if attr, ok := d.GetOk("binpacking.0.disruption.0.budget"); ok {
		budgetDtos := attr.([]interface{})
		budgets := []nextgen.DisruptionBudget{}
		for _, budgetDto := range budgetDtos {
			budget := budgetDto.(map[string]interface{})
			b := nextgen.DisruptionBudget{
				Reasons: getDisruptionBudgetReasons(budget),
				Nodes:   budget["nodes"].(string),
			}
			if len(budget["schedule"].([]interface{})) > 0 {
				frequency := budget["schedule"].([]interface{})[0].(map[string]interface{})["frequency"].(string)
				duration := budget["schedule"].([]interface{})[0].(map[string]interface{})["duration"].(string)
				if frequency != "" && duration != "" {
					b.Schedule = &frequency
					b.Duration = duration

				}
			}
			budgets = append(budgets, b)
		}
		config.Consolidation.Budgets = budgets
	}
	if _, ok := d.GetOk("node_preferences"); ok {
		if attr, ok := d.GetOk("node_preferences.0.ttl"); ok {
			expiry := attr.(string)
			config.Consolidation.NodeExpiry = &expiry
		}
		if attr, ok := d.GetOk("node_preferences.0.reverse_fallback_interval"); ok {
			intrvl := attr.(string)
			config.ReverseFallback = &nextgen.ReverseFallback{
				Enabled:       true,
				RetryInterval: intrvl,
			}
		}
	}
	populateCommitmentIntegrationConfig(d, config)
	populateReplacementWindowConfig(d, config)
	return *config
}
func getDisruptionBudgetReasons(b map[string]interface{}) []string {
	reasons := b["reasons"].([]interface{})
	if len(reasons) == 0 {
		return []string{
			"Drifted", "Underutilized", "Empty",
		}
	}
	reasonList := []string{}
	for _, reason := range reasons {
		reasonList = append(reasonList, reason.(string))
	}
	return reasonList
}

func populateCommitmentIntegrationConfig(d *schema.ResourceData, config *nextgen.ClusterOrchConfig) {
	if _, ok := d.GetOk("commitment_integration"); ok {
		config.CommitmentIntegration = &nextgen.CommitmentIntegration{
			Enabled:        d.Get("commitment_integration.0.enabled").(bool),
			CloudAccountID: d.Get("commitment_integration.0.master_account_id").(string),
		}
	}
}

func populateReplacementWindowConfig(d *schema.ResourceData, config *nextgen.ClusterOrchConfig) {
	replacementWindow := nextgen.ReplacementWindow{}
	_, ok := d.GetOk("replacement_schedule")
	if !ok {
		return
	}
	replacementWindow.AppliesTo = &nextgen.WindowAppliesTo{
		HarnessPodEviction: d.Get("replacement_schedule.0.applies_to.0.harness_pod_eviction").(bool),
		Consolidation:      d.Get("replacement_schedule.0.applies_to.0.consolidation").(bool),
		ReverseFallback:    d.Get("replacement_schedule.0.applies_to.0.reverse_fallback").(bool),
	}
	replacementWindow.ReplacementWindowType = nextgen.ReplacementWindowType(d.Get("replacement_schedule.0.window_type").(string))
	details, ok := d.GetOk("replacement_schedule.0.window_details")
	if ok {
		windowDetails := details.([]interface{})
		days := windowDetails[0].(map[string]interface{})["days"].([]interface{})
		var weekDays []time.Weekday
		for _, day := range days {
			idx, validDay := dayIndex[day.(string)]
			if validDay {
				weekDays = append(weekDays, idx)
			}
		}
		timeZone := windowDetails[0].(map[string]interface{})["time_zone"].(string)
		replacementWindow.WindowDetails = &nextgen.WindowDetails{
			Days:     weekDays,
			TimeZone: timeZone,
		}
		allDay, allDayOk := windowDetails[0].(map[string]interface{})["all_day"]
		startTime, startTimeOk := windowDetails[0].(map[string]interface{})["start_time"]
		endTime, endTimeOk := windowDetails[0].(map[string]interface{})["end_time"]
		if allDayOk {
			replacementWindow.WindowDetails.AllDay = allDay.(bool)
		}
		if startTimeOk {
			startTimeStr, _ := startTime.(string)
			replacementWindow.WindowDetails.StartTime = parseTimeInDay(startTimeStr, nextgen.TimeInDayForWindow{Hour: 0, Min: 0})
		}
		if endTimeOk {
			endTimeStr, _ := endTime.(string)
			replacementWindow.WindowDetails.EndTime = parseTimeInDay(endTimeStr, nextgen.TimeInDayForWindow{Hour: 23, Min: 59})
		}
	}
	config.ReplacementWindow = &replacementWindow
	return
}

func parseTimeInDay(timeInDayStr string, defaultTimeInDay nextgen.TimeInDayForWindow) *nextgen.TimeInDayForWindow {
	timeParts := strings.Split(strings.TrimSpace(timeInDayStr), ":")
	timeInDay := defaultTimeInDay
	if len(timeParts) == 2 {
		timeHr, err := strconv.ParseInt(timeParts[0], 10, 64)
		if err == nil {
			timeInDay.Hour = int(timeHr)
		}
		timeMin, err := strconv.ParseInt(timeParts[1], 10, 64)
		if err == nil {
			timeInDay.Min = int(timeMin)
		}
	}
	return &timeInDay
}

func readClusterOrchConfig(d *schema.ResourceData, orch *nextgen.ClusterOrchestrator) {
	d.SetId(orch.ID)
	d.Set("disabled", orch.Disabled)
	d.Set("distribution", []interface{}{map[string]interface{}{
		"base_ondemand_capacity":      orch.Config.BaseOnDemandCapacity,
		"ondemand_replica_percentage": orch.Config.OnDemandSplit,
		"selector":                    orch.Config.SpotDistribution,
		"strategy":                    orch.Config.DistributionStrategy,
	}})
	d.Set("binpacking", []interface{}{map[string]interface{}{
		"pod_eviction": getPodEvictionConfig(orch),
		"disruption":   getDisruptionConfig(orch),
	}})
	d.Set("node_preferences", []interface{}{map[string]interface{}{
		"ttl":                       orch.Config.Consolidation.NodeExpiry,
		"reverse_fallback_interval": getReverseFBInterval(orch),
	}})

	if commitmentIntegration := getCommitmentIntegrationIfExists(orch.Config); len(commitmentIntegration) > 0 {
		d.Set("commitment_integration", commitmentIntegration)
	}

	if replacementWindow := getReplacementWindowIfExists(orch.Config); len(replacementWindow) > 0 {
		d.Set("replacement_schedule", replacementWindow)
	}
}

func getCommitmentIntegrationIfExists(cfg *nextgen.ClusterOrchConfig) []interface{} {
	if cfg.CommitmentIntegration == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"enabled":           cfg.CommitmentIntegration.Enabled,
		"master_account_id": cfg.CommitmentIntegration.CloudAccountID,
	}}
}

func getReplacementWindowIfExists(cfg *nextgen.ClusterOrchConfig) []interface{} {
	if cfg.ReplacementWindow == nil {
		return nil
	}
	replacementWindow := map[string]interface{}{
		"window_type": cfg.ReplacementWindow.ReplacementWindowType,
		"applies_to": []interface{}{
			map[string]interface{}{
				"harness_pod_eviction": cfg.ReplacementWindow.AppliesTo.HarnessPodEviction,
				"consolidation":        cfg.ReplacementWindow.AppliesTo.Consolidation,
				"reverse_fallback":     cfg.ReplacementWindow.AppliesTo.ReverseFallback,
			},
		},
	}
	if cfg.ReplacementWindow.WindowDetails != nil {
		replacementWindow["window_details"] = getWindowDetails(cfg)
	}
	return []interface{}{replacementWindow}
}

func getWindowDetails(cfg *nextgen.ClusterOrchConfig) []interface{} {
	var windowDetailsMap map[string]interface{}
	windowDetails := cfg.ReplacementWindow.WindowDetails
	if windowDetails == nil {
		return nil
	}
	var days []string
	for _, day := range windowDetails.Days {
		days = append(days, reverseDayIndex[day])
	}
	windowDetailsMap = map[string]interface{}{
		"days":      days,
		"all_day":   windowDetails.AllDay,
		"time_zone": windowDetails.TimeZone,
	}
	if windowDetails.StartTime != nil {
		windowDetailsMap["start_time"] = fmt.Sprintf("%02d:%02d", windowDetails.StartTime.Hour, windowDetails.StartTime.Min)
	}
	if windowDetails.EndTime != nil {
		windowDetailsMap["end_time"] = fmt.Sprintf("%02d:%02d", windowDetails.EndTime.Hour, windowDetails.EndTime.Min)
	}
	return []interface{}{windowDetailsMap}
}

func getPodEvictionConfig(orch *nextgen.ClusterOrchestrator) []interface{} {
	podEvictorCfg := orch.Config.Consolidation.PodEvictor
	if podEvictorCfg.Enabled {
		return []interface{}{
			map[string]interface{}{
				"threshold": map[string]interface{}{
					"cpu":    podEvictorCfg.MinCPU,
					"memory": podEvictorCfg.MinMem,
				},
			},
		}
	}
	return nil
}

func getDisruptionConfig(orch *nextgen.ClusterOrchestrator) []interface{} {
	disruptionCfg := orch.Config.Consolidation
	disruptionDto := map[string]interface{}{
		"criteria": disruptionCfg.ConsolidationPolicy,
		"delay":    disruptionCfg.ConsolidateAfter,
	}
	var budgets []interface{}
	for _, budget := range disruptionCfg.Budgets {
		budgets = append(budgets, map[string]interface{}{
			"reasons": budget.Reasons,
			"nodes":   budget.Nodes,
			"schedule": map[string]interface{}{
				"frequency": budget.Schedule,
				"duration":  budget.Duration,
			},
		})
	}
	if len(budgets) > 0 {
		disruptionDto["budgets"] = budgets
	}
	return []interface{}{disruptionDto}
}

func getReverseFBInterval(orch *nextgen.ClusterOrchestrator) string {
	if orch.Config.ReverseFallback != nil {
		return orch.Config.ReverseFallback.RetryInterval
	}
	return ""
}
