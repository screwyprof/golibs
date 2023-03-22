package ctxtags_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"

	"github.com/screwyprof/golibs/assert"
	"github.com/screwyprof/golibs/gin/middleware/ctxtags"
)

func TestCtxTags(t *testing.T) {
	r := givenMiddlewareIsSetup(t)
	res := whenRequestIsHandledByMiddleware(r)
	thenTagsAreSet(t, res)
}

func givenMiddlewareIsSetup(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := func(c *gin.Context) {
		tags := ctxtags.FromContext(c)
		assert.True(t, tags.Has(ctxtags.XRequestID))
		c.JSON(http.StatusOK, tags)
	}

	r.Use(ctxtags.CtxTags(ctxtags.WithFieldExtractor(ctxtags.RequestID)))
	r.GET("/", h)
	return r
}

func whenRequestIsHandledByMiddleware(r *gin.Engine) *apitest.Request {
	apiTester := apitest.New().Handler(r)
	return apiTester.Get("/")
}

func thenTagsAreSet(t *testing.T, r *apitest.Request) {
	res := r.Expect(t)
	res.Status(http.StatusOK).
		// Body(requestID).
		End()
}
