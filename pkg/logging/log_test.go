package logging

import (
	"testing"
	"os"
	"strings"
)

func TestOS(t *testing.T) {
	execFile := os.Args[0]
	lastIndex := strings.LastIndex(execFile, "/")
	execFile = string([]rune(execFile)[lastIndex+1:])
	t.Log(execFile)
}
