package geetest

import (
	"testing"

	configmocks "github.com/goravel/framework/mocks/config"
	"github.com/stretchr/testify/suite"
)

type GeetestTestSuite struct {
	suite.Suite
	instance *Geetest
}

func TestGeetestTestSuite(t *testing.T) {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "geetest.captcha_id").Return("12345").Once()
	mockConfig.On("GetString", "geetest.captcha_key").Return("67890").Once()
	mockConfig.On("GetString", "geetest.api_url").Return("https://gcaptcha4.geetest.com").Once()

	suite.Run(t, &GeetestTestSuite{
		instance: NewGeetest(mockConfig),
	})
}

func (s *GeetestTestSuite) TestNewGeetest() {
	s.NotNil(s.instance)
}

func (s *GeetestTestSuite) TestVerify() {
	verify, err := s.instance.Verify(Ticket{
		LotNumber:     "12345",
		CaptchaOutput: "12345",
		PassToken:     "12345",
		GenTime:       "12345",
	})
	s.Error(err)
	s.False(verify)
	s.EqualError(err, "-50102 illegal captcha_id")
}
