package handler_test

import (
	"context"
	"fmt"
	"kolo_marvel_project/config"
	"kolo_marvel_project/internal/server"
	"kolo_marvel_project/internal/server/handler"
	"kolo_marvel_project/pkg/cache"
	"kolo_marvel_project/pkg/marvel"
	"kolo_marvel_project/utils/initialize"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

var router *gin.Engine
var params server.Options

func setupMockServer() (router *gin.Engine) {
	// params = server.Options{}
	gin.SetMode(gin.TestMode)
	app := fx.New(
		fx.Provide(
			//redis server
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		handler.Module,
		server.Module,
		marvel.Module,
		cache.Module,
		// Run app forever

		fx.Populate(&params),
	)
	app.Start(context.TODO())
	defer app.Stop(context.TODO())
	router = server.SetupRouter(&params)
	return
}
func init() {
	router = setupMockServer()
}
func TestHealthz(t *testing.T) {
	// router = setupMockServer()
	assert.NotNil(t, router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"ok":"ok"}`, w.Body.String())
}

func TestFetchCharacterDetails(t *testing.T) {
	// router = setupMockServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/marvel/character/search", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"code":0,"exception":"kolo_marvel_project.UncaughtException","message":"Oops! Something went wrong. Please try later"}`, w.Body.String())
}
func TestFetchCharacterDetailsWithPage(t *testing.T) {
	// router = setupMockServer()
	base, _ := url.Parse("/v1/marvel/character/search")
	params := url.Values{}
	params.Add("page", fmt.Sprint(1))
	base.RawQuery = params.Encode()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", base.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"success":true,"message":"Marvel Character Details Successful","data":{"offset":0,"limit":10,"total":1562,"count":10,"results":[{"id":1011334,"name":"3-D Man","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/c/e0/535fecbbb9784","extension":"jpg"}},{"id":1017100,"name":"A-Bomb (HAS)","description":"Rick Jones has been Hulk's best bud since day one, but now he's more than a friend...he's a teammate! Transformed by a Gamma energy explosion, A-Bomb's thick, armored skin is just as strong and powerful as it is blue. And when he curls into action, he uses it like a giant bowling ball of destruction! ","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/3/20/5232158de5b16","extension":"jpg"}},{"id":1009144,"name":"A.I.M.","description":"AIM is a terrorist organization bent on destroying the world.","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/6/20/52602f21f29ec","extension":"jpg"}},{"id":1010699,"name":"Aaron Stack","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1009146,"name":"Abomination (Emil Blonsky)","description":"Formerly known as Emil Blonsky, a spy of Soviet Yugoslavian origin working for the KGB, the Abomination gained his powers after receiving a dose of gamma radiation similar to that which transformed Bruce Banner into the incredible Hulk.","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/9/50/4ce18691cbf04","extension":"jpg"}},{"id":1016823,"name":"Abomination (Ultimate)","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1009148,"name":"Absorbing Man","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/1/b0/5269678709fb7","extension":"jpg"}},{"id":1009149,"name":"Abyss","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/9/30/535feab462a64","extension":"jpg"}},{"id":1010903,"name":"Abyss (Age of Apocalypse)","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/3/80/4c00358ec7548","extension":"jpg"}},{"id":1011266,"name":"Adam Destine","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}}]},"meta":null}`, w.Body.String())
}
func TestFetchCharacterDetailsWithPageAndCharacterName(t *testing.T) {
	// router = setupMockServer()
	base, _ := url.Parse("/v1/marvel/character/search")
	params := url.Values{}
	params.Add("page", fmt.Sprint(1))
	params.Add("nameStartsWith", "th")
	base.RawQuery = params.Encode()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", base.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"success":true,"message":"Marvel Character Details Successful","data":{"offset":0,"limit":10,"total":49,"count":10,"results":[{"id":1011003,"name":"Thaddeus Ross","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1009652,"name":"Thanos","description":"The Mad Titan Thanos, a melancholy, brooding individual, consumed with the concept of death, sought out personal power and increased strength, endowing himself with cybernetic implants until he became more powerful than any of his brethren.","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/6/40/5274137e3e2cd","extension":"jpg"}},{"id":1015004,"name":"Thanos (Ultimate)","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/5/c0/5317734cbc1d0","extension":"jpg"}},{"id":1011083,"name":"The 198","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1009653,"name":"The Anarchist","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/1/60/4c003aacdeca9","extension":"jpg"}},{"id":1009654,"name":"The Call","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1010714,"name":"The Captain","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1012080,"name":"The Collector (Taneleer Tivan)","description":"The Collector is one of the oldest living beings in the universe, having been among the first of the universe's races to become sentient in the wake of the Big Bang.","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1011294,"name":"The Enforcers","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/b/40/image_not_available","extension":"jpg"}},{"id":1010728,"name":"The Executioner","description":"","thumbnail":{"path":"http://i.annihil.us/u/prod/marvel/i/mg/e/d0/4ce5a2ab860be","extension":"jpg"}}]},"meta":null}`, w.Body.String())
}
func TestFailBindingStruct(t *testing.T) {
	// router = setupMockServer()
	base, _ := url.Parse("/v1/marvel/character/search")
	params := url.Values{}
	params.Add("testing", fmt.Sprint(1))
	params.Add("demo", "th")
	base.RawQuery = params.Encode()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", base.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"code":0,"exception":"kolo_marvel_project.UncaughtException","message":"Oops! Something went wrong. Please try later"}`, w.Body.String())
}
