package ctxtags_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/gin/middleware/ctxtags"
)

func TestFromContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	tags := ctxtags.FromContext(c)

	assert.NotNil(t, tags)
}
