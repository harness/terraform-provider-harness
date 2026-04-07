package project

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func testResourceSchema() *schema.ResourceData {
	r := ResourceProject()
	d := r.TestResourceData()
	return d
}

func TestSetProjectDetails_NilOrphanedResources(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos:       []string{"*"},
			OrphanedResources: nil,
		},
	}

	// This used to panic with nil pointer dereference before the fix
	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "test-project", d.Id())
	assert.Equal(t, "account123", d.Get("account_id"))
}

func TestSetProjectDetails_WithOrphanedResources(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos: []string{"*"},
			OrphanedResources: &nextgen.AppprojectsOrphanedResourcesMonitorSettings{
				Warn: true,
			},
		},
	}

	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "test-project", d.Id())

	// Verify orphaned_resources warn value is set correctly via the schema
	warn := d.Get("project.0.spec.0.orphaned_resources.0.warn").(bool)
	assert.True(t, warn)
}

func TestSetProjectDetails_OrphanedResourcesWarnFalse(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos: []string{"*"},
			OrphanedResources: &nextgen.AppprojectsOrphanedResourcesMonitorSettings{
				Warn: false,
			},
		},
	}

	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	warn := d.Get("project.0.spec.0.orphaned_resources.0.warn").(bool)
	assert.False(t, warn)
}

func TestSetProjectDetails_NilSpec(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: nil,
	}

	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "test-project", d.Id())
}

func TestSetProjectDetails_MultipleDestinations(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos: []string{"*"},
			Destinations: []nextgen.AppprojectsApplicationDestination{
				{
					Namespace: "default",
					Server:    "https://kubernetes.default.svc",
					Name:      "in-cluster",
				},
				{
					Namespace: "staging",
					Server:    "https://staging.example.com",
					Name:      "staging-cluster",
				},
				{
					Namespace: "production",
					Server:    "https://prod.example.com",
					Name:      "prod-cluster",
				},
			},
		},
	}

	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "test-project", d.Id())

	// Verify all three destinations are stored, not just the first one
	dest0NS := d.Get("project.0.spec.0.destinations.0.namespace").(string)
	assert.Equal(t, "default", dest0NS)
	dest0Name := d.Get("project.0.spec.0.destinations.0.name").(string)
	assert.Equal(t, "in-cluster", dest0Name)
	dest0Server := d.Get("project.0.spec.0.destinations.0.server").(string)
	assert.Equal(t, "https://kubernetes.default.svc", dest0Server)

	dest1NS := d.Get("project.0.spec.0.destinations.1.namespace").(string)
	assert.Equal(t, "staging", dest1NS)
	dest1Name := d.Get("project.0.spec.0.destinations.1.name").(string)
	assert.Equal(t, "staging-cluster", dest1Name)
	dest1Server := d.Get("project.0.spec.0.destinations.1.server").(string)
	assert.Equal(t, "https://staging.example.com", dest1Server)

	dest2NS := d.Get("project.0.spec.0.destinations.2.namespace").(string)
	assert.Equal(t, "production", dest2NS)
	dest2Name := d.Get("project.0.spec.0.destinations.2.name").(string)
	assert.Equal(t, "prod-cluster", dest2Name)
	dest2Server := d.Get("project.0.spec.0.destinations.2.server").(string)
	assert.Equal(t, "https://prod.example.com", dest2Server)
}

func TestSetProjectDetails_FullProjectWithoutOrphanedResources(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos: []string{"*"},
			Destinations: []nextgen.AppprojectsApplicationDestination{
				{
					Namespace: "default",
					Server:    "https://kubernetes.default.svc",
					Name:      "in-cluster",
				},
			},
			ClusterResourceWhitelist: []nextgen.V1GroupKind{
				{Group: "*", Kind: "Namespace"},
			},
			OrphanedResources: nil,
		},
	}

	// Should not panic even with multiple fields set and OrphanedResources nil
	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "test-project", d.Id())
	assert.Equal(t, "account123", d.Get("account_id"))

	// Verify other spec fields are set correctly
	server := d.Get("project.0.spec.0.destinations.0.server").(string)
	assert.Equal(t, "https://kubernetes.default.svc", server)
}
