package split

import (
	"fmt"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
)

// ruleBasedSegmentResolveTrafficTypeID resolves traffic type ID for import using
// client.RuleBasedSegments.List workspace metadata (trafficTypeId and/or nested trafficType).
func ruleBasedSegmentResolveTrafficTypeID(client *splitsdk.APIClient, wsID, segName string) (string, error) {
	metaList, err := client.RuleBasedSegments.List(wsID)
	if err != nil {
		return "", fmt.Errorf("rule-based segment %q: workspace list: %w", segName, err)
	}
	for _, m := range metaList {
		if m.Name != segName {
			continue
		}
		if m.TrafficTypeID != "" {
			return m.TrafficTypeID, nil
		}
		if m.TrafficType != nil && m.TrafficType.ID != "" {
			return m.TrafficType.ID, nil
		}
		return "", fmt.Errorf("rule-based segment %q matched workspace list but metadata had no trafficTypeId or trafficType.id", segName)
	}
	return "", fmt.Errorf("rule-based segment %q not found in workspace list", segName)
}
