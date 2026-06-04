package folders

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// folderListResponse matches the API response for listing folders (top level may nest via SubFolders).
type folderListResponse struct {
	Resource []nextgen.Folder `json:"resource"`
}

func DataSourceDashboardFolders() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a list of Harness Custom Dashboard Folders.",

		ReadContext: dataSourceFoldersListRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of a specific folder to filter the list by (optional).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"folders": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Identifier of the folder.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the folder.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"created_at": {
							Description: "Created DateTime of the folder.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	return resource
}

func dataSourceFoldersListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, _ := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	allFolders, httpResp, err := listDashboardFolders(c)
	if err != nil {
		if httpResp != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		return diag.FromErr(fmt.Errorf("failed to list dashboard folders: %w", err))
	}

	// Flatten nested sub_folders so the data source returns "all"
	flat := flattenFolders(allFolders)

	// Optional name filter (exact match)
	if name, ok := d.GetOk("name"); ok && name.(string) != "" {
		filtered := make([]nextgen.Folder, 0)
		for _, f := range flat {
			if f.Name == name.(string) {
				filtered = append(filtered, f)
			}
		}
		flat = filtered
	}

	folderMaps := make([]map[string]interface{}, 0, len(flat))
	for _, f := range flat {
		folderMaps = append(folderMaps, map[string]interface{}{
			"id":         f.Id,
			"name":       f.Name,
			"created_at": f.CreatedAt,
		})
	}

	d.SetId(c.AccountId)
	if err := d.Set("folders", folderMaps); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// listDashboardFolders performs a raw GET to /dashboard/folders because the harness-go-sdk
// (v0.7.30) does not expose a ListFolders / GetFolders method on DashboardsFolderApi.
func listDashboardFolders(c *nextgen.APIClient) ([]nextgen.Folder, *http.Response, error) {
	if c == nil {
		return nil, nil, errors.New("client is nil")
	}
	base := strings.TrimRight(c.Endpoint, "/")
	// accountId query param is required by the API (same as other folder ops)
	u := fmt.Sprintf("%s/dashboard/folders?accountId=%s", base, c.AccountId)

	req, err := retryablehttp.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Create a retry client similar to the provider's getHttpClient (simple config is sufficient for reads)
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 3

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer httpResp.Body.Close()

	body, readErr := io.ReadAll(httpResp.Body)
	if readErr != nil {
		return nil, httpResp, readErr
	}

	if httpResp.StatusCode >= 300 {
		// Let caller use HandleApiError with the resp; return body in err for visibility if needed
		return nil, httpResp, fmt.Errorf("API error: status=%s body=%s", httpResp.Status, string(body))
	}

	var listResp folderListResponse
	if jsonErr := json.Unmarshal(body, &listResp); jsonErr == nil && listResp.Resource != nil {
		return listResp.Resource, httpResp, nil
	}

	// Fallback: response might be a bare array of folders
	var direct []nextgen.Folder
	if jsonErr := json.Unmarshal(body, &direct); jsonErr == nil {
		return direct, httpResp, nil
	}

	return nil, httpResp, fmt.Errorf("unexpected list folders response shape: %s", string(body))
}

func flattenFolders(input []nextgen.Folder) []nextgen.Folder {
	var out []nextgen.Folder
	for _, f := range input {
		out = append(out, f)
		if len(f.SubFolders) > 0 {
			out = append(out, flattenFolders(f.SubFolders)...)
		}
	}
	return out
}
