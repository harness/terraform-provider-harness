package nextgen

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
)

func (p *ProjectApiService) GetProjectByName(ctx context.Context, accountId string, organizationId string, name string) (*ProjectResponse, *http.Response, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, httpResp, err := p.GetProjectList(ctx, accountId, &ProjectApiGetProjectListOpts{
			SearchTerm:    optional.NewString(name),
			OrgIdentifier: optional.NewString(organizationId),
			PageIndex:     optional.NewInt32(pageIndex),
			PageSize:      optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, httpResp, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, httpResp, nil
		}

		for _, project := range resp.Data.Content {
			if project.Project.Name == name {
				return &project, httpResp, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil, nil
}
