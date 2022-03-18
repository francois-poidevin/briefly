package schema

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	services "github.com/francois-poidevin/briefly/internal/services/v1"
	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	sf "github.com/sony/sonyflake"
)

type Service struct {
	sonyFlake *sf.Sonyflake
	schemav1.UnimplementedSchemaAPIServer
}

// New services instance
func New() services.Schema {
	sfSettings := sf.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			rand.Seed(time.Now().UnixNano())
			val := rand.Intn(65535 + 1)
			return uint16(val), nil
		},
	}
	sfl := sf.NewSonyflake(sfSettings)
	return &Service{
		sonyFlake: sfl,
	}
}

// Get the unshortcoded URL.
func (s *Service) GetUnShortCodedURL(ctx context.Context, req *schemav1.GetUnShortCodedURLRequest) (*schemav1.GetUnShortCodedURLResponse, error) {
	//TODO: code here
	resp := schemav1.GetUnShortCodedURLResponse{
		Url: "https://www.google.com",
	}
	return &resp, nil
}

// Get the Hash of the URL shortcoded.
func (s *Service) GetShortCodeHash(ctx context.Context, req *schemav1.GetShortCodeHashRequest) (*schemav1.GetShortCodeHashResponse, error) {
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
