package idp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"gopkg.in/yaml.v3"
)

func TestAccResourceCatalogEntity(t *testing.T) {
	description := t.Name()
	id := fmt.Sprintf("%s_%s", description, utils.RandStringBytes(5))
	updatedDescription := fmt.Sprintf("%s_updated", description)
	resourceName := "harness_platform_idp_catalog_entity.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCatalogEntityDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCatalogEntity(id, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					testAccEntityCheckYamlField(resourceName, "metadata.description", description),
				),
			},
			{
				Config: testAccResourceCatalogEntity(id, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					testAccEntityCheckYamlField(resourceName, "metadata.description", updatedDescription),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccCatalogEntityImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceCatalogEntity(id string, description string) string {
	str := fmt.Sprintf(`
		resource "harness_platform_idp_catalog_entity" "test" {
			identifier = "%[1]s"
			kind = "component"
			yaml = <<-EOT
	        apiVersion: harness.io/v1
	        kind: Component
	        name: Example Catalog
	        identifier: "%[1]s"
	        type: service
	        owner: user:account/admin@harness.io
	        spec:
	            lifecycle: prod
	        metadata:
	            tags:
		            - test
	            description: "%[2]s"
	        EOT
		}
	`, id, description)

	return str
}

func TestAccResourceRemoteCatalogEntity(t *testing.T) {
	description := t.Name()
	id := fmt.Sprintf("%s_%s", description, utils.RandStringBytes(5))
	resourceName := "harness_platform_idp_catalog_entity.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCatalogEntityDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRemoteCatalogEntity(id, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					testAccEntityCheckYamlField(resourceName, "metadata.description", description),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccCatalogEntityImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceRemoteCatalogEntity(id string, description string) string {
	str := fmt.Sprintf(`
		resource "harness_platform_idp_catalog_entity" "test" {
			identifier = "%[1]s"
			org_id = "default"
			project_id = "ssem"
			kind = "component"
			git_details {
		    store_type = "REMOTE"
			connector_ref = "demossem"
			repo_name = "catalog"
			branch_name = "main"
			file_path = "gitimport/%[1]s.yaml"
			}
			yaml = <<-EOT
	        apiVersion: harness.io/v1
	        kind: Component
	        orgIdentifier: default
	        projectIdentifier: ssem
	        name: Example Catalog
	        identifier: "%[1]s"
	        type: service
	        owner: user:account/admin@harness.io
	        spec:
	            lifecycle: prod
	        metadata:
	            tags:
		            - test
	            description: "%[2]s"
	        EOT
		}
	`, id, description)

	return str
}

func TestAccResourceImportRemoteCatalogEntity(t *testing.T) {

	resourceName := "harness_platform_idp_catalog_entity.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCatalogEntityDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceImportRemoteCatalogEntity(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "test_import"),
					testAccEntityCheckYamlField(resourceName, "kind", "Component"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccCatalogEntityImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceImportRemoteCatalogEntity() string {
	str := `
		resource "harness_platform_idp_catalog_entity" "test" {
			identifier = "test_import"
			org_id = "default"
			project_id = "ssem"
			import_from_git = true
			git_details {
		    store_type = "REMOTE"
			connector_ref = "demossem"
			repo_name = "catalog"
			branch_name = "main"
			file_path = "gitimport/test_import.yaml"
			}
		}
	`

	return str
}

func testAccGetCatalogEntity(resourceName string, state *terraform.State) (*idp.EntityResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetIDPClientWithContext()
	id := r.Primary.ID

	info, err := getCatalogEntityInfo(r.Primary)
	if err != nil {
		return nil, err
	}

	resp, _, err := c.EntitiesApi.GetEntity(ctx, info.Scope, info.Kind, id, &idp.EntitiesApiGetEntityOpts{
		OrgIdentifier:     info.OrgId,
		ProjectIdentifier: info.ProjectId,
		HarnessAccount:    optional.NewString(c.AccountId),
	})
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

func testAccCatalogEntityDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		catalogEntity, _ := testAccGetCatalogEntity(resourceName, state)
		if catalogEntity != nil {
			return fmt.Errorf("Found catalog entity: %s", catalogEntity.Identifier)
		}
		return nil
	}
}

func testAccEntityCheckYamlField(resourceName, key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r := acctest.TestAccGetResource(resourceName, state)
		yamlString := r.Primary.Attributes["yaml"]
		var yamlData map[string]any
		if err := yaml.Unmarshal([]byte(yamlString), &yamlData); err != nil {
			return err
		}

		parts := strings.Split(key, ".")

		var v any
		for i := range parts {
			var ok bool
			v, ok = yamlData[parts[i]]
			if !ok {
				return fmt.Errorf("field %s not found in yaml", parts[i])
			}

			next, ok := v.(map[string]any)
			if !ok {
				// If we've reached the end and it's not a map, we're done
				if i == len(parts)-1 {
					yamlData = nil
					break
				}
				return fmt.Errorf("field %s is not a map in yaml", parts[i])
			}
			yamlData = next
		}

		// At this point, v contains the value we're looking for
		if yamlData != nil {
			return fmt.Errorf("field %s is not a leaf node in yaml", key)
		}

		// Check if the final value matches what we expect
		if v != value {
			return fmt.Errorf("field %s expected %s, got %v", key, value, v)
		}

		return nil
	}
}

type catalogEntityInfo struct {
	Scope      string
	Kind       string
	Identifier string
	OrgId      optional.String
	ProjectId  optional.String
}

func getCatalogEntityInfo(d *terraform.InstanceState) (catalogEntityInfo, error) {
	kind := d.Attributes["kind"]
	identifier := d.Attributes["identifier"]
	orgId := d.Attributes["org_id"]
	projectId := d.Attributes["project_id"]

	catalogInfo := catalogEntityInfo{
		Kind:       kind,
		Scope:      "account",
		Identifier: identifier,
	}

	if orgId != "" {
		catalogInfo.OrgId = optional.NewString(orgId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, orgId)
	} else {
		catalogInfo.OrgId = optional.EmptyString()
	}

	if projectId != "" {
		catalogInfo.ProjectId = optional.NewString(projectId)
		catalogInfo.Scope = fmt.Sprintf("%s.%s", catalogInfo.Scope, projectId)
	} else {
		catalogInfo.ProjectId = optional.EmptyString()
	}

	return catalogInfo, nil
}

func testAccCatalogEntityImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		info, err := getCatalogEntityInfo(rs.Primary)
		if err != nil {
			return "", err
		}

		if info.Scope == "account" {
			return fmt.Sprintf("%s/%s", info.Kind, info.Identifier), nil
		}

		scope, _ := strings.CutPrefix(info.Scope, "account.")

		return fmt.Sprintf("%s/%s/%s", scope, info.Kind, info.Identifier), nil
	}
}
