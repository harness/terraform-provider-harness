package nextgen

import (
	"context"

	"github.com/antihax/optional"
)

type GetServiceByNameOpts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (s *ServicesApiService) GetServiceByName(ctx context.Context, accountId string, name string, opts GetServiceByNameOpts) (*ServiceResponseDetails, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, _, err := s.GetServiceList(ctx, accountId, &ServicesApiGetServiceListOpts{
			OrgIdentifier:     opts.OrgIdentifier,
			ProjectIdentifier: opts.ProjectIdentifier,
			SearchTerm:        optional.NewString(name),
			Page:              optional.NewInt32(pageIndex),
			Size:              optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, nil
		}

		for _, svc := range resp.Data.Content {
			if svc.Service.Name == name {
				return svc.Service, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil
}
