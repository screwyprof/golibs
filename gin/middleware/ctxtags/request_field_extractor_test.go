package ctxtags_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/gin-gonic/gin"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/gin/middleware/ctxtags"
)

func TestRequestID(t *testing.T) {
	want, c := givenRequestWithRequestID(t)
	tags := whenFieldsAreParsed(c)
	assertTagsExtracted(t, tags, want)
}

func givenRequestWithRequestID(t *testing.T) (string, *gin.Context) {
	want := gofakeit.UUID()

	c := createTestCtx(t)
	c.Request.Header.Set(ctxtags.XRequestID, want)
	return want, c
}

func whenFieldsAreParsed(c *gin.Context) map[string]interface{} {
	return ctxtags.RequestID(c)
}

func createTestCtx(t *testing.T) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	r, err := http.NewRequest("POST", "/", nil)
	assert.NoError(t, err)

	c.Request = r
	return c
}

func assertTagsExtracted(t *testing.T, tags map[string]interface{}, want string) {
	got, ok := tags[ctxtags.XRequestID]
	assert.True(t, ok)
	assert.Equals(t, got, want)
}
