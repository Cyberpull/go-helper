package tests

import (
	"testing"

	"cyberpull.com/gotk/v2/http"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite
}

func (s *HttpTestSuite) TestGetString() {
	_, err := http.Get[string]("https://www.cyberpull.com")
	require.NoError(s.T(), err)
}

func (s *HttpTestSuite) TestGetJson() {
	_, err := http.Get[[]any]("https://restcountries.com/v3.1/all")
	require.NoError(s.T(), err)
}

func (s *HttpTestSuite) TestClientGetString() {
	client := http.NewClient[string](&http.Options{
		BaseURL: "https://www.cyberpull.com",
	})

	_, err := client.Get("")

	require.NoError(s.T(), err)
}

func (s *HttpTestSuite) TestClientGetJson() {
	client := http.NewClient[[]any](&http.Options{
		BaseURL: "https://restcountries.com/v3.1",
	})

	_, err := client.Get("all")

	require.NoError(s.T(), err)
}

// ========================

func TestHttp(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
