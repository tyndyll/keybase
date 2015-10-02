package client

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestBaseUrlTemplate(t *testing.T) {
	Convey(`Then the clients baseUrl will be set correctly`, t, func() {
		So(BASE_URL_TEMPLATE, ShouldEqual, "https://keybase.io/_/api/1.0/<command>.json")
	})
}

func TestNew(t *testing.T) {
	Convey(`When I call New`, t, func() {
		c := New()

		Convey(`Then I will get a client the same type as ApiClient`, func() {
			So(c, ShouldHaveSameTypeAs, &ApiClient{})
		})

	})
}

func TestApiClient(t *testing.T) {
	Convey(`Given I have an ApiClient`, t, func() {
		client := &ApiClient{}

		calledMethod := ""
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calledMethod = r.Method
		}))
		defer ts.Close()

		client.baseUrl = ts.URL

		Convey(`When I call the Get method`, func() {
			_, _ = client.Get("", nil)

			Convey(`Then I will have made GET call to the remote server `, func() {
				So(calledMethod, ShouldEqual, "GET")
			})
		})

		Convey(`When I call the Post method`, func() {
			_, _ = client.Post("", nil)

			Convey(`Then I will have made a POST call to the remote server`, func() {
				So(calledMethod, ShouldEqual, "POST")
			})
		})
	})
}
