package viewx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {

	Init("Default title")

	httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Render(w, "home", nil)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestRenderHTML(t *testing.T) {

	httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	RenderHTML(w, "views/index", nil)
	RenderHTML(w, "views.index", nil)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}
}
