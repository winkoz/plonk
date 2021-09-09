package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	sut Service
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.sut = &service{allowedHosts: []string{"winkoz.com"}}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

func (suite *ServiceTestSuite) TestIsValidUrl_ReturnsTrue_WhenPassedURLIsISOComplaint() {
	assert.True(suite.T(), suite.sut.IsValidUrl("https://winkoz.com/plonk"))
}

func (suite *ServiceTestSuite) TestIsValidUrl_ReturnsFalse_WhenPassedURLIsNotISOComplaint() {
	assert.False(suite.T(), suite.sut.IsValidUrl("not_valid_url"))
}

func (suite *ServiceTestSuite) TestIsValidUrl_ReturnsFalse_WhenURLIsNotWithinAllowList() {
	assert.False(suite.T(), suite.sut.IsValidUrl("https://not-allowed-url.com"))
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
