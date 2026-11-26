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

func TestSetProjectDetails_WithDestinationServiceAccounts(t *testing.T) {
	d := testResourceSchema()
	project := &nextgen.AppprojectsAppProject{
		Metadata: &nextgen.V1ObjectMeta{
			Name:      "test-project",
			Namespace: "argocd",
		},
		Spec: &nextgen.AppprojectsAppProjectSpec{
			SourceRepos: []string{"*"},
			DestinationServiceAccounts: []nextgen.AppprojectsApplicationDestinationServiceAccount{
				{
					Server:                "https://kubernetes.default.svc",
					Namespace:             "myapp-*",
					DefaultServiceAccount: "argocd-deployer",
				},
				{
					Server:                "https://staging.example.com",
					Namespace:             "staging",
					DefaultServiceAccount: "staging-deployer",
				},
			},
		},
	}

	assert.NotPanics(t, func() {
		setProjectDetails(d, "account123", project)
	})

	assert.Equal(t, "https://kubernetes.default.svc", d.Get("project.0.spec.0.destination_service_accounts.0.server"))
	assert.Equal(t, "myapp-*", d.Get("project.0.spec.0.destination_service_accounts.0.namespace"))
	assert.Equal(t, "argocd-deployer", d.Get("project.0.spec.0.destination_service_accounts.0.default_service_account"))
	assert.Equal(t, "https://staging.example.com", d.Get("project.0.spec.0.destination_service_accounts.1.server"))
	assert.Equal(t, "staging", d.Get("project.0.spec.0.destination_service_accounts.1.namespace"))
	assert.Equal(t, "staging-deployer", d.Get("project.0.spec.0.destination_service_accounts.1.default_service_account"))
}

func TestCreateAndUpdateRequestBody_DestinationServiceAccounts(t *testing.T) {
	d := testResourceSchema()
	err := d.Set("project", []interface{}{
		map[string]interface{}{
			"metadata": []interface{}{
				map[string]interface{}{
					"name":      "test-project",
					"namespace": "argocd",
				},
			},
			"spec": []interface{}{
				map[string]interface{}{
					"destination_service_accounts": []interface{}{
						map[string]interface{}{
							"server":                  "https://kubernetes.default.svc",
							"namespace":               "myapp-*",
							"default_service_account": "argocd-deployer",
						},
					},
				},
			},
		},
	})
	assert.NoError(t, err)

	created := createRequestBody(d)
	assert.NotNil(t, created.Project)
	assert.NotNil(t, created.Project.Spec)
	assert.Equal(t, []nextgen.AppprojectsApplicationDestinationServiceAccount{
		{
			Server:                "https://kubernetes.default.svc",
			Namespace:             "myapp-*",
			DefaultServiceAccount: "argocd-deployer",
		},
	}, created.Project.Spec.DestinationServiceAccounts)

	updated := updateRequestBody(d)
	assert.NotNil(t, updated.Project)
	assert.NotNil(t, updated.Project.Spec)
	assert.Equal(t, created.Project.Spec.DestinationServiceAccounts, updated.Project.Spec.DestinationServiceAccounts)
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
