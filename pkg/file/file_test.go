package file

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetExecPath(t *testing.T) {
	dirName, err := GetDirName()
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, dirName, "go-gin-example")
}
