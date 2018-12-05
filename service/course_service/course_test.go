package course_service

import (
	"github.com/Quons/go-gin-example/models"
	"github.com/Quons/go-gin-example/pkg/setting"
	"testing"
)

func init() {
	setting.Setup("dev")
	models.Setup()
}

func TestAddArticleAndTag(t *testing.T) {

}
