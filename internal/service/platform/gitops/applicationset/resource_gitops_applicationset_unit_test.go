package applicationset

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBuildApplicationSetGeneratorSelector covers CDS-129262, where a generator level
// `selector` was accepted by the schema but never included in the create/update payload.
func TestBuildApplicationSetGeneratorSelector(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsApplicationSet().Schema, map[string]interface{}{
		"applicationset": []interface{}{
			map[string]interface{}{
				"metadata": []interface{}{
					map[string]interface{}{"name": "test-appset"},
				},
				"spec": []interface{}{
					map[string]interface{}{
						"generator": []interface{}{
							map[string]interface{}{
								"git": []interface{}{
									map[string]interface{}{
										"repo_url": "https://github.com/argoproj/argocd-example-apps.git",
										"revision": "HEAD",
										"file": []interface{}{
											map[string]interface{}{"path": "platform/**/harness-config.yaml"},
										},
									},
								},
								"selector": []interface{}{
									map[string]interface{}{
										"match_labels": map[string]interface{}{"team": "platform"},
										"match_expressions": []interface{}{
											map[string]interface{}{
												"key":      "chart",
												"operator": "Exists",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	appset := buildApplicationSet(d)

	require.NotNil(t, appset)
	require.NotNil(t, appset.Spec)
	require.Len(t, appset.Spec.Generators, 1)

	selector := appset.Spec.Generators[0].Selector
	require.NotNil(t, selector, "generator selector must be sent to the API")
	assert.Equal(t, map[string]string{"team": "platform"}, selector.MatchLabels)
	require.Len(t, selector.MatchExpressions, 1)
	assert.Equal(t, "chart", selector.MatchExpressions[0].Key)
	assert.Equal(t, "Exists", selector.MatchExpressions[0].Operator)
	assert.Empty(t, selector.MatchExpressions[0].Values)
}

// TestSetApplicationSetGeneratorSelector asserts the generator selector returned by the API is
// written back into state. Without this the config drifts and every plan re-adds the selector.
func TestSetApplicationSetGeneratorSelector(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsApplicationSet().Schema, map[string]interface{}{})

	appset := &nextgen.Servicev1ApplicationSet{
		Appset: &nextgen.ApplicationsApplicationSet{
			Metadata: &nextgen.V1ObjectMeta{Name: "test-appset"},
			Spec: &nextgen.ApplicationsApplicationSetSpec{
				Generators: []nextgen.ApplicationsApplicationSetGenerator{
					{
						Git: &nextgen.ApplicationsGitGenerator{
							RepoURL:  "https://github.com/argoproj/argocd-example-apps.git",
							Revision: "HEAD",
						},
						Selector: &nextgen.V1LabelSelector{
							MatchLabels: map[string]string{"team": "platform"},
							MatchExpressions: []nextgen.V1LabelSelectorRequirement{
								{Key: "chart", Operator: "Exists"},
							},
						},
					},
				},
			},
		},
	}

	require.NoError(t, setApplicationSet(d, appset))

	assert.Equal(t, "chart", d.Get("applicationset.0.spec.0.generator.0.selector.0.match_expressions.0.key"))
	assert.Equal(t, "Exists", d.Get("applicationset.0.spec.0.generator.0.selector.0.match_expressions.0.operator"))
	assert.Equal(t, "platform", d.Get("applicationset.0.spec.0.generator.0.selector.0.match_labels.team"))
}

// TestBuildApplicationSetNestedGeneratorSelector covers the same drop on matrix child generators.
func TestBuildApplicationSetNestedGeneratorSelector(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsApplicationSet().Schema, map[string]interface{}{
		"applicationset": []interface{}{
			map[string]interface{}{
				"metadata": []interface{}{
					map[string]interface{}{"name": "test-appset"},
				},
				"spec": []interface{}{
					map[string]interface{}{
						"generator": []interface{}{
							map[string]interface{}{
								"matrix": []interface{}{
									map[string]interface{}{
										"generator": []interface{}{
											map[string]interface{}{
												"clusters": []interface{}{
													map[string]interface{}{"enabled": true},
												},
											},
											map[string]interface{}{
												"git": []interface{}{
													map[string]interface{}{
														"repo_url": "https://github.com/argoproj/argocd-example-apps.git",
														"revision": "HEAD",
													},
												},
												"selector": []interface{}{
													map[string]interface{}{
														"match_expressions": []interface{}{
															map[string]interface{}{
																"key":      "env",
																"operator": "In",
																"values":   []interface{}{"prod"},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	appset := buildApplicationSet(d)

	require.NotNil(t, appset)
	require.Len(t, appset.Spec.Generators, 1)
	require.NotNil(t, appset.Spec.Generators[0].Matrix)
	require.Len(t, appset.Spec.Generators[0].Matrix.Generators, 2)

	selector := appset.Spec.Generators[0].Matrix.Generators[1].Selector
	require.NotNil(t, selector, "nested generator selector must be sent to the API")
	require.Len(t, selector.MatchExpressions, 1)
	assert.Equal(t, "env", selector.MatchExpressions[0].Key)
	assert.Equal(t, "In", selector.MatchExpressions[0].Operator)
	assert.Equal(t, []string{"prod"}, selector.MatchExpressions[0].Values)
}

// TestBuildApplicationSetWithoutSelector guards against sending an empty selector when the
// optional block is omitted, which would otherwise change the generated ApplicationSet.
func TestBuildApplicationSetWithoutSelector(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceGitopsApplicationSet().Schema, map[string]interface{}{
		"applicationset": []interface{}{
			map[string]interface{}{
				"metadata": []interface{}{
					map[string]interface{}{"name": "test-appset"},
				},
				"spec": []interface{}{
					map[string]interface{}{
						"generator": []interface{}{
							map[string]interface{}{
								"git": []interface{}{
									map[string]interface{}{
										"repo_url": "https://github.com/argoproj/argocd-example-apps.git",
										"revision": "HEAD",
									},
								},
							},
						},
					},
				},
			},
		},
	})

	appset := buildApplicationSet(d)

	require.NotNil(t, appset)
	require.Len(t, appset.Spec.Generators, 1)
	assert.Nil(t, appset.Spec.Generators[0].Selector)
}
