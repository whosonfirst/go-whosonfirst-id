package id

import (
	_ "github.com/aaronland/go-uid-whosonfirst"
)

import (
	"context"
	"fmt"
	"github.com/aaronland/go-uid"
	_ "github.com/aaronland/go-uid-proxy"
)

type Provider interface {
	NewID(context.Context) (int64, error)
}

type WOFProvider struct {
	Provider
	uid_provider uid.Provider
}

func NewProvider(ctx context.Context) (Provider, error) {

	uri := "proxy://?provider=whosonfirst://"
	return NewProviderWithURI(ctx, uri)
}

func NewProviderWithURI(ctx context.Context, uri string) (Provider, error) {

	uid_pr, err := uid.NewProvider(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new provider, %w", err)
	}

	wof_pr := &WOFProvider{
		uid_provider: uid_pr,
	}

	return wof_pr, nil
}

func (wof_pr *WOFProvider) NewID(ctx context.Context) (int64, error) {

	v, err := wof_pr.uid_provider.UID(ctx)

	if err != nil {
		return -1, fmt.Errorf("Failed to generate ID, %w", err)
	}

	id, ok := uid.AsInt64(v)

	if !ok {
		return -1, fmt.Errorf("Provider return invalid value")
	}

	return id, nil
}
