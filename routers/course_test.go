package routers

import (
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	"net/url"
	"github.com/Quons/go-gin-example/pkg/setting"
	"github.com/Quons/go-gin-example/models"
	"github.com/Quons/go-gin-example/pkg/logging"
)

func init() {
	setting.Setup("dev")
	models.Setup()
	logging.Setup()
}

func TestPingRoute(t *testing.T) {
	router := InitRouter()

	w := httptest.NewRecorder()
	postData := url.Values{}
	postData.Set("token", "1")
	postData.Set("courseId", "1")

	t.Log(postData.Encode())
	req, err := http.NewRequest("POST", "/api/v1/getCourse?token=feahvJWLZP88FddSbhv_1NMdddPJMEvXvNiUDCKxLIpMVgOpXvrqGjDhgs1mxKFP&courseId=1", nil)
	if err != nil {
		t.Error(err)
		return
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
