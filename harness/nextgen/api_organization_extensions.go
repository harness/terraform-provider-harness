package nextgen

import (
	"context"

	"github.com/antihax/optional"
)

func (o *OrganizationApiService) GetOrganizationByName(ctx context.Context, accountId string, name string) (*OrganizationResponse, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, _, err := o.GetOrganizationList(ctx, accountId, &OrganizationApiGetOrganizationListOpts{
			SearchTerm: optional.NewString(name),
			PageIndex:  optional.NewInt32(pageIndex),
			PageSize:   optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, nil
		}

		for _, org := range resp.Data.Content {
			if org.Organization.Name == name {
				return &org, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil
}
