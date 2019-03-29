package brand_service

import (
	"github.com/Quons/go-gin-example/models"
	"github.com/Quons/go-gin-example/pkg/setting"
	"testing"
)

func init() {
	setting.Setup("dev")
	models.Setup()
}
func TestGetBrandList(t *testing.T) {
	b := &Brand{}
	list, err := b.GetBrandList()
	if err != nil {
		t.Error(err)
		return

	}
	t.Logf("%v", list)
}
