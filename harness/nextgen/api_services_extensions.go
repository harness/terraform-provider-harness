package nextgen

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
)

type GetServiceByNameOpts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (s *ServicesApiService) GetServiceByName(ctx context.Context, accountId string, name string, opts GetServiceByNameOpts) (*ServiceResponseDetails, *http.Response, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, httpResp, err := s.GetServiceList(ctx, accountId, &ServicesApiGetServiceListOpts{
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
			if svc.Service.Name == name {
				return svc.Service, httpResp, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil, nil
}
