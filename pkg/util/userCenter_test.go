package util

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string
}

func TestMethod(t *testing.T) {
	var user = User{}
	if checkTest(&user) && user.Name == "quon" {
		t.Log("hiahiahi")
	}
	t.Logf("%v", user)
}

func checkTest(user *User) bool {
	str := `{"Name":"quon"}`
	json.Unmarshal([]byte(str), user)
	return true
}
