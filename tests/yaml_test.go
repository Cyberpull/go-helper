package tests

import (
	"testing"

	"cyberpull.com/gotk/v2/yaml"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type YamlDemoData struct {
	Name  string        `yaml:"name"`
	Email string        `yaml:"email"`
	Phone YamlDemoPhone `yaml:"phone"`
}

type YamlDemoPhone struct {
	Home string `yaml:"home"`
	Work string `yaml:"work"`
}

// ========================

type YamlTestSuite struct {
	suite.Suite
}

func (s *YamlTestSuite) TestReadFile() {
	value, err := yaml.ReadFile[*YamlDemoData]("files/demo.yml")
	require.NoError(s.T(), err)
	require.NotNil(s.T(), value)

	assert.IsType(s.T(), &YamlDemoData{}, value)

	assert.Equal(s.T(), "Christian Ezeani", value.Name)
	assert.Equal(s.T(), "christian@example.com", value.Email)

	assert.Equal(s.T(), "070xxxxxxxx", value.Phone.Home)
	assert.Equal(s.T(), "070xxxxxxxx", value.Phone.Work)
}

// ========================

func TestYaml(t *testing.T) {
	suite.Run(t, new(YamlTestSuite))
}
