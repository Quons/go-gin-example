package models

import "github.com/Quons/go-gin-example/pkg/setting"

func init() {
	setting.Setup("dev")
	Setup()
}
