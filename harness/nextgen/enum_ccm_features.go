package nextgen

type CCMFeature string

var CCMFeatures = struct {
	Billing      CCMFeature
	Optimization CCMFeature
	Visibility   CCMFeature
	Governance   CCMFeature
}{
	Billing:      "BILLING",
	Optimization: "OPTIMIZATION",
	Visibility:   "VISIBILITY",
	Governance:   "GOVERNANCE",
}

var CCMFeaturesSlice = []string{
	CCMFeatures.Billing.String(),
	CCMFeatures.Optimization.String(),
	CCMFeatures.Visibility.String(),
	CCMFeatures.Governance.String(),
}

func (e CCMFeature) String() string {
	return string(e)
}
