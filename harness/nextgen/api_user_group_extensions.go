package nextgen

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
)

type UserGroupApiGetUserGroupByNameOpts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (a *UserGroupApiService) GetUserGroupByName(ctx context.Context, accountIdentifier string, name string, opts *UserGroupApiGetUserGroupByNameOpts) (*UserGroup, *http.Response, error) {

	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		resp, httpResp, err := a.GetUserGroupList(ctx, accountIdentifier, &UserGroupApiGetUserGroupListOpts{
			OrgIdentifier:     opts.OrgIdentifier,
			ProjectIdentifier: opts.ProjectIdentifier,
			SearchTerm:        optional.NewString(name),
			PageIndex:         optional.NewInt32(pageIndex),
			PageSize:          optional.NewInt32(pageSize),
		})

		if err != nil {
			return nil, httpResp, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, httpResp, nil
		}

		for _, userGroup := range resp.Data.Content {
			if userGroup.Name == name {
				return &userGroup, httpResp, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil, nil
}
