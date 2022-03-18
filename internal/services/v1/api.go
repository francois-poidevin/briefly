package v1

import (
	"context"

	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
)

type Schema interface {
	// Get the unshortcoded URL.
	GetUnShortCodedURL(ctx context.Context, req *schemav1.GetUnShortCodedURLRequest) (*schemav1.GetUnShortCodedURLResponse, error)
	// GetUnShortCodedURL(ctx context.Context, in *schemav1.GetUnShortCodedURLRequest, opts ...grpc.CallOption) (*schemav1.GetUnShortCodedURLResponse, error)
	// Get the Hash of the URL shortcoded.
	GetShortCodeHash(context.Context, *schemav1.GetShortCodeHashRequest) (*schemav1.GetShortCodeHashResponse, error)
	// GetShortCodeHash(ctx context.Context, in *schemav1.GetShortCodeHashRequest, opts ...grpc.CallOption) (*schemav1.GetShortCodeHashResponse, error)
}
