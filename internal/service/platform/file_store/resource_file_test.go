package file_store_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFileStoreFile(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_file_store_node_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFileStore_File(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				Config: testAccResourceFileStore_File(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"file_content_path"},
				ImportStateIdFunc:       acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceFileStoreFileOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_file_store_node_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFileStore_FileOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceFileStore_FileOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"file_content_path"},
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceFileStoreFileProjectLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_file_store_node_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFileStore_FileProjectLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				Config: testAccResourceFileStore_FileProjectLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"file_content_path"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceFileStore_FileProjectLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}

	resource "harness_platform_file_store_node_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = "%[1]s"
		project_id  = "%[1]s"
		description = "test file"
		tags = ["foo:bar", "bar:foo"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path = "%[3]s"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "4s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

func testAccResourceFileStore_FileOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_file_store_node_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = "%[1]s"
		description = "test file"
		tags = ["foo:bar"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path = "%[3]s"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "4s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

func testAccResourceFileStore_File(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_file_store_node_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test file"
		tags = ["foo:bar", "bar:foo"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path =  "%[3]s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

// common methods for file and folder

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccFileStoreDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		fileNode, _ := testAccGetFile(resourceName, state)
		if fileNode != nil {
			return fmt.Errorf("Found file store node: %s", fileNode.Identifier)
		}
		return nil
	}
}

func testAccGetFile(resourceName string, state *terraform.State) (*nextgen.File, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.FileStoreApi.GetFile(ctx, id, c.AccountId, &nextgen.FileStoreApiGetFileOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
	})

	if resp.Data == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
