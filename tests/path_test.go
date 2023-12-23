package tests

import (
	"testing"

	"cyberpull.com/gotk/v2"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PathTestSuite struct {
	suite.Suite
}

func (s *PathTestSuite) TestPathValue() {
	value := gotk.Path("files", "demo.yml")
	require.NotEmpty(s.T(), value)
}

func (s *PathTestSuite) TestExecutablePathValue() {
	value, err := gotk.PathFromExecutable("files", "demo.yml")
	require.NoError(s.T(), err)
	require.FileExists(s.T(), value)
}

// ========================

func TestPath(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}
