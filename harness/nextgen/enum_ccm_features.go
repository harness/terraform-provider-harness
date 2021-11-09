package nextgen

type CCMFeature string

var CCMFeatures = struct {
	Billing      CCMFeature
	Optimization CCMFeature
	Visibility   CCMFeature
}{
	Billing:      "BILLING",
	Optimization: "OPTIMIZATION",
	Visibility:   "VISIBILITY",
}

var CCMFeaturesSlice = []string{
	CCMFeatures.Billing.String(),
	CCMFeatures.Optimization.String(),
	CCMFeatures.Visibility.String(),
}

func (e CCMFeature) String() string {
	return string(e)
}
