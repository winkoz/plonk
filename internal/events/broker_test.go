package events

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BrokerTestSuite struct {
	suite.Suite
	sut     Broker
	sutChan chan Message
}

func (suite *BrokerTestSuite) SetupTest() {
	suite.sutChan = make(chan Message, 1)
	suite.sut = broker{
		bus: suite.sutChan,
	}
}

//-------------------------------------------------
// Tests
//-------------------------------------------------

func (suite *BrokerTestSuite) Test_PostMessage_ShouldSubmitMessageToChannel() {
	msg := NewCommandSelectedMessage(suite.T().Name())
	suite.sut.PostMessage(msg)
	require.Equal(suite.T(), 1, len(suite.sutChan))
	for m := range suite.sutChan {
		require.Equal(suite.T(), msg, m)
		break
	}
}

func (suite *BrokerTestSuite) Test_GetBrokerChannel_ShouldReturnTheBrokerChannel() {
	gotChannel := suite.sut.GetBrokerChannel()
	var expectedChannel <-chan Message = suite.sutChan
	require.Equal(suite.T(), expectedChannel, gotChannel)
}

func TestBrokerTestSuite(t *testing.T) {
	suite.Run(t, new(BrokerTestSuite))
}
