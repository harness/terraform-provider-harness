package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetConnectorByName(t *testing.T) {
	c, ctx := getClientWithContext()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	conn := Connector{
		Connector: &ConnectorInfo{
			Name:       name,
			Identifier: name,
			Type_:      ConnectorTypes.K8sCluster,
			K8sCluster: &KubernetesClusterConfig{
				Credential: &KubernetesCredential{
					Type_: KubernetesCredentialTypes.InheritFromDelegate,
				},
				DelegateSelectors: []string{"primary"},
			},
		},
	}
	connector, _, err := c.ConnectorsApi.CreateConnector(ctx, conn, c.AccountId, &ConnectorsApiCreateConnectorOpts{})

	defer func() {
		c.ConnectorsApi.DeleteConnector(ctx, c.AccountId, name, &ConnectorsApiDeleteConnectorOpts{})
	}()

	require.NoError(t, err)
	require.NotNil(t, connector)

	foundConnector, err := c.ConnectorsApi.GetConnectorByName(ctx, c.AccountId, name, ConnectorTypes.K8sCluster, ConnectorsApiGetConnectorByNameOpts{})
	require.NoError(t, err)
	require.NotNil(t, foundConnector)
}

func TestDeleteConnectorForceDelete(t *testing.T) {
	c, ctx := getClientWithContext()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	conn := Connector{
		Connector: &ConnectorInfo{
			Name:       name,
			Identifier: name,
			Type_:      ConnectorTypes.K8sCluster,
			K8sCluster: &KubernetesClusterConfig{
				Credential: &KubernetesCredential{
					Type_: KubernetesCredentialTypes.InheritFromDelegate,
				},
				DelegateSelectors: []string{"primary"},
			},
		},
	}

	connector, _, err := c.ConnectorsApi.CreateConnector(ctx, conn, c.AccountId, &ConnectorsApiCreateConnectorOpts{})
	require.NoError(t, err)
	require.NotNil(t, connector)

	r, _, err := c.ConnectorsApi.DeleteConnector(ctx, c.AccountId, name, &ConnectorsApiDeleteConnectorOpts{ForceDelete: optional.NewBool(true)})
	require.NoError(t, err)
	require.NotNil(t, r)
	require.Equal(t, r.Data, true)
}
