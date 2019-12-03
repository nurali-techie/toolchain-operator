package olm

import (
	"context"
	"testing"

	orgv1 "github.com/eclipse/che-operator/pkg/apis/org/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CheClusterAssertion struct {
	t              *testing.T
	client         client.Reader
	namespacedName types.NamespacedName
	cheCluster     *orgv1.CheCluster
}

func AssertThatCheCluster(t *testing.T, ns, name string, client client.Reader) *CheClusterAssertion {
	return &CheClusterAssertion{
		t:              t,
		client:         client,
		namespacedName: types.NamespacedName{Namespace: ns, Name: name},
	}
}

func (a *CheClusterAssertion) Exists() *CheClusterAssertion {
	err := a.loadCheClusterAssertion()
	require.NoError(a.t, err)
	return a
}

func (a *CheClusterAssertion) HasRunningStatus(want string) *CheClusterAssertion {
	a.Exists()
	assert.Equal(a.t, want, a.cheCluster.Status.CheClusterRunning)
	return a
}

func (a *CheClusterAssertion) loadCheClusterAssertion() error {
	if a.cheCluster != nil {
		return nil
	}
	cheCluster := &orgv1.CheCluster{}
	err := a.client.Get(context.TODO(), a.namespacedName, cheCluster)
	a.cheCluster = cheCluster
	return err
}
