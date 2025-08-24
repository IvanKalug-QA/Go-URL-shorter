package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code int
		// response string
		// contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Positive test #1",
			want: want{
				code: 200,
				// response: forms.UrlForm,
				// contentType: "html",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, `/`, nil)
			w := httptest.NewRecorder()
			MainPage(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, test.want.code)
		})
	}
}
