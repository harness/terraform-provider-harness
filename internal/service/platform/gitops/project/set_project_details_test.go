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
