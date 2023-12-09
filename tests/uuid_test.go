package tests

import (
	"testing"

	"cyberpull.com/gotk/v2"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UUIDTestSuite struct {
	suite.Suite
}

func (s *UUIDTestSuite) TestValue() {
	_, err := gotk.UUID()
	require.NoError(s.T(), err)
}

// ========================

func TestUUID(t *testing.T) {
	suite.Run(t, new(UUIDTestSuite))
}
