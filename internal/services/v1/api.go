package v1

import (
	"context"

	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	"google.golang.org/grpc"
)

type Schema interface {
	// Get the unshortcoded URL.
	GetUnShortCodedURL(ctx context.Context, in *schemav1.GetUnShortCodedURLRequest, opts ...grpc.CallOption) (*schemav1.GetUnShortCodedURLResponse, error)
	// Get the Hash of the URL shortcoded.
	GetShortCodeHash(ctx context.Context, in *schemav1.GetShortCodeHashRequest, opts ...grpc.CallOption) (*schemav1.GetShortCodeHashResponse, error)
}
