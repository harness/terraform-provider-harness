package nextgen

import (
	"context"

	"github.com/antihax/optional"
)

type ConnectorsApiGetConnectorByNameOpts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (a *ConnectorsApiService) GetConnectorByName(ctx context.Context, accountId string, name string, connectorType ConnectorType, opts ConnectorsApiGetConnectorByNameOpts) (*ConnectorInfo, error) {
	var pageIndex int32 = 0
	var pageSize int32 = 2

	for true {
		filters := ConnectorFilterProperties{
			ConnectorNames: []string{name},
			Types:          []string{connectorType.String()},
			FilterType:     ConnectorFilterTypes.Connector,
		}

		resp, _, err := a.GetConnectorListV2(ctx, filters, accountId, &ConnectorsApiGetConnectorListV2Opts{
			PageIndex:         optional.NewInt32(pageIndex),
			PageSize:          optional.NewInt32(pageSize),
			OrgIdentifier:     opts.OrgIdentifier,
			ProjectIdentifier: opts.ProjectIdentifier,
		})

		if err != nil {
			return nil, err
		}

		if len(resp.Data.Content) == 0 {
			return nil, nil
		}

		for _, svc := range resp.Data.Content {
			if svc.Connector.Name == name {
				return svc.Connector, nil
			}
		}

		pageIndex += pageSize
	}

	return nil, nil
}
