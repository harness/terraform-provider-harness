package nextgen

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
)

type GetEnvironmentByNameOpts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (e *EnvironmentsApiService) GetEnvironmentByName(ctx context.Context, accountId string, name string, opts GetEnvironmentByNameOpts) (*EnvironmentResponseDetails, *http.Response, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, httpResp, err := e.GetEnvironmentList(ctx, accountId, &EnvironmentsApiGetEnvironmentListOpts{
			OrgIdentifier:     opts.OrgIdentifier,
			ProjectIdentifier: opts.ProjectIdentifier,
			SearchTerm:        optional.NewString(name),
			Page:              optional.NewInt32(pageIndex),
			Size:              optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, httpResp, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, httpResp, nil
		}

		for _, svc := range resp.Data.Content {
			if svc.Environment.Name == name {
				return svc.Environment, httpResp, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil, nil
}
