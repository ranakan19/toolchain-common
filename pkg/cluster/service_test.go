package cluster_test

import (
	"testing"

	"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	"github.com/codeready-toolchain/toolchain-common/pkg/cluster"
	"github.com/codeready-toolchain/toolchain-common/pkg/test"
	"github.com/codeready-toolchain/toolchain-common/pkg/test/verify"
)

func TestAddToolchainClusterAsMember(t *testing.T) {
	// given & then
	verify.AddToolchainClusterAsMember(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		return service.AddOrUpdateToolchainCluster(toolchainCluster)
	})

}

func TestAddToolchainClusterAsHost(t *testing.T) {
	// given & then
	verify.AddToolchainClusterAsHost(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		return service.AddOrUpdateToolchainCluster(toolchainCluster)
	})
}

func TestAddToolchainClusterFailsBecauseOfMissingSecret(t *testing.T) {
	// given & then
	verify.AddToolchainClusterFailsBecauseOfMissingSecret(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		return service.AddOrUpdateToolchainCluster(toolchainCluster)
	})
}

func TestAddToolchainClusterFailsBecauseOfEmptySecret(t *testing.T) {
	// given & then
	verify.AddToolchainClusterFailsBecauseOfEmptySecret(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		return service.AddOrUpdateToolchainCluster(toolchainCluster)
	})
}

func TestUpdateToolchainCluster(t *testing.T) {
	// given & then
	verify.UpdateToolchainCluster(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		return service.AddOrUpdateToolchainCluster(toolchainCluster)
	})
}

func TestDeleteToolchainClusterWhenDoesNotExist(t *testing.T) {
	// given & then
	verify.DeleteToolchainCluster(t, func(toolchainCluster *v1alpha1.ToolchainCluster, cl *test.FakeClient, service cluster.ToolchainClusterService) error {
		// when
		service.DeleteToolchainCluster("east")
		return nil
	})
}
