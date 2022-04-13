package delegate_test

import (
	"context"
	"sync"

	"github.com/harness/harness-go-sdk/harness/delegate"
)

var delegateImagePull sync.Once

func pullDelegateImage(ctx context.Context, cfg *delegate.DockerDelegateConfig) {
	delegateImagePull.Do(func() {
		delegate.PullDelegateImage(ctx, cfg)
	})
}
