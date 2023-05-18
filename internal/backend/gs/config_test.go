package gs

import (
	"testing"

	"github.com/restic/restic/internal/backend/test"
)

var configTests = []test.ConfigTestData[Config]{
	{"gs:bucketname:/", Config{
		Bucket:      "bucketname",
		Prefix:      "",
		Connections: 5,
	}},
	{"gs:bucketname:/prefix/directory", Config{
		Bucket:      "bucketname",
		Prefix:      "prefix/directory",
		Connections: 5,
	}},
	{"gs:bucketname:/prefix/directory/", Config{
		Bucket:      "bucketname",
		Prefix:      "prefix/directory",
		Connections: 5,
	}},
}

func TestParseConfig(t *testing.T) {
	test.ParseConfigTester(t, ParseConfig, configTests)
}
