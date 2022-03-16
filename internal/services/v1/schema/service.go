package schema

import (
	"context"
	"fmt"
	"strconv"
	"time"

	services "github.com/francois-poidevin/briefly/internal/services/v1"
	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	sf "github.com/sony/sonyflake"
	"google.golang.org/grpc"
)

type service struct {
	sonyFlake *sf.Sonyflake
}

// New services instance
func New() services.Schema {
	sfSettings := sf.Settings{
		StartTime: time.Now(),
	}
	sfl := sf.NewSonyflake(sfSettings)
	return &service{
		sonyFlake: sfl,
	}
}

// Get the unshortcoded URL.
func (s *service) GetUnShortCodedURL(ctx context.Context, in *schemav1.GetUnShortCodedURLRequest, opts ...grpc.CallOption) (*schemav1.GetUnShortCodedURLResponse, error) {
	//TODO: code here
	resp := schemav1.GetUnShortCodedURLResponse{
		Url: "https://www.google.com",
	}
	return &resp, nil
}

// Get the Hash of the URL shortcoded.
func (s *service) GetShortCodeHash(ctx context.Context, in *schemav1.GetShortCodeHashRequest, opts ...grpc.CallOption) (*schemav1.GetShortCodeHashResponse, error) {
	//TODO: code here
	id, err := s.sonyFlake.NextID()
	if err != nil {
		fmt.Println("Error dude !")
	}
	resp := schemav1.GetShortCodeHashResponse{
		Hash: strconv.FormatUint(id, 10),
	}
	return &resp, nil
}
