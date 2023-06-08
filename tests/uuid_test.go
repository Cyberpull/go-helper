package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UUIDTestSuite struct {
	suite.Suite
}

// ========================

func TestUUID(t *testing.T) {
	suite.Run(t, new(UUIDTestSuite))
}
