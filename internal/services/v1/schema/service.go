package schema

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	services "github.com/francois-poidevin/briefly/internal/services/v1"
	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	"github.com/sirupsen/logrus"
	sf "github.com/sony/sonyflake"
)

type Service struct {
	ctx       context.Context
	log       *logrus.Logger
	sonyFlake *sf.Sonyflake
	schemav1.UnimplementedSchemaAPIServer
}

// New services instance
func New(ctx context.Context, log *logrus.Logger) services.Schema {
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
		ctx:       ctx,
		log:       log,
		sonyFlake: sfl,
	}
}

// Get the unshortcoded URL.
func (s *Service) GetUnShortCodedURL(ctx context.Context, req *schemav1.GetUnShortCodedURLRequest) (*schemav1.GetUnShortCodedURLResponse, error) {
	s.log.WithContext(s.ctx).WithFields(logrus.Fields{"RequestID ": req.Hash}).Debug("Request ID to search")

	//TODO code here the search in database
	url := "https://www.google.com"
	s.log.WithContext(s.ctx).WithFields(logrus.Fields{"ResponseURL ": url}).Debug("Response URL")

	resp := schemav1.GetUnShortCodedURLResponse{
		Url: url,
	}
	return &resp, nil
}

// Get the Hash of the URL shortcoded.
func (s *Service) GetShortCodeHash(ctx context.Context, req *schemav1.GetShortCodeHashRequest) (*schemav1.GetShortCodeHashResponse, error) {

	s.log.WithContext(s.ctx).WithFields(logrus.Fields{"RequestURL ": req.Url}).Debug("Request URL to encode")
	//TODO search in database the same URL, if already shortcoded, then return the associated ID if exist, otherwise do the following code

	id, err := s.sonyFlake.NextID()
	if err != nil {
		s.log.WithContext(s.ctx).WithFields(logrus.Fields{"Error": err}).Error("Error during ID generation")
		return nil, err
	}

	s.log.WithContext(s.ctx).WithFields(logrus.Fields{"id": id}).Debug("SnowFlake generated ID")

	resp := schemav1.GetShortCodeHashResponse{
		Hash: strconv.FormatUint(id, 10),
	}
	return &resp, nil
}
