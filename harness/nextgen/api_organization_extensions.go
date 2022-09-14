package nextgen

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
)

func (o *OrganizationApiService) GetOrganizationByName(ctx context.Context, accountId string, name string) (*OrganizationResponse, *http.Response, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, httpResp, err := o.GetOrganizationList(ctx, accountId, &OrganizationApiGetOrganizationListOpts{
			SearchTerm: optional.NewString(name),
			PageIndex:  optional.NewInt32(pageIndex),
			PageSize:   optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, httpResp, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, httpResp, nil
		}

		for _, org := range resp.Data.Content {
			if org.Organization.Name == name {
				return &org, httpResp, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil, nil
}
