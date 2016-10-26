package gosession

import (
	"bytes"
	"github.com/suboat/go-session"
	"net/http"
	"testing"
)

type ResponseRecorder struct {
	Code      int
	HeaderMap http.Header
	Body      *bytes.Buffer
	Flushed   bool
}

func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
	}
}

func (rw *ResponseRecorder) Header() http.Header {
	return rw.HeaderMap
}

func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	if rw.Body != nil {
		rw.Body.Write(buf)
	}
	if rw.Code == 0 {
		rw.Code = http.StatusOK
	}
	return len(buf), nil
}

func (rw *ResponseRecorder) WriteHeader(code int) {
	rw.Code = code
}

func (rw *ResponseRecorder) Flush() {
	rw.Flushed = true
}

func TestSession(t *testing.T) {
	store := NewSession()
	if err := store.Store("", []byte("dsadadsdasdadsadsd")); err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	rsp := NewRecorder()

	if session, err := store.Get(req, "session-key"); err != nil {
		t.Fatalf("Error getting session: %v", err)
	} else {
		flashes := session.Flashes()
		if len(flashes) != 0 {
			t.Errorf("Expected empty flashes; Got %v", flashes)
		}
		session.AddFlash("foo")
		session.AddFlash("bar")

		session.AddFlash("baz", "custom_key")

		if err := session.Save(req, rsp); err != nil {
			t.Fatalf("Error saving session: %v", err)
		}

		flashes = session.Flashes()
		if len(flashes) == 0 {
			t.Fatal("Unexpected empty flashes")
		}
		for _, flash := range flashes {
			t.Logf("%#v", flash)
		}

		flashes = session.Flashes("custom_key")
		if len(flashes) != 1 {
			t.Errorf("Expected flashes; Got %v", flashes)
		} else if flashes[0] != "baz" {
			t.Errorf("Expected baz; Got %v", flashes)
		}
	}
	hdr := rsp.Header()
	if cookies, ok := hdr["Set-Cookie"]; !ok || len(cookies) != 1 {
		t.Fatal("No cookies. Header:", hdr)
	}

	if _, err := store.Get(req, "session:key"); err == nil {
		//t.Fatal("Expected error is nil")
	} else if err.Error() != "sessions: invalid character in cookie name: session:key" {
		t.Fatal("Expected error due to invalid cookie name")
	}
}

func TestSessionSimple(t *testing.T) {
	store := NewSession()
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	rsp := NewRecorder()

	if err := store.Session(req, gosession.SessionKeyword); err != nil {
		t.Fatalf("Error getting session: %v", err)
	}

	if tag := store.GetTag(req); len(tag) != 0 {
		t.Errorf("Expected empty tag; Got %v", tag)
	} else if err := store.SetTag(req, rsp, "this is a tag"); err != nil {
		t.Fatalf("Error SetTag: %v", err)
	} else if tag := store.GetTag(req); len(tag) == 0 {
		t.Errorf("Expected empty tag; Got %v", tag)
	} else {
		t.Logf("tag: %v", tag)
	}

	if level := store.GetLevel(req); level != 0 {
		t.Errorf("Expected empty level; Got %v", level)
	} else if err := store.SetLevel(req, rsp, 10); err != nil {
		t.Fatalf("Error SetLevel: %v", err)
	} else if level := store.GetLevel(req); level == 0 {
		t.Errorf("Expected empty level; Got %v", level)
	} else {
		t.Logf("level: %v", level)
	}

	if status := store.GetStatus(req); status != 0 {
		t.Errorf("Expected empty status; Got %v", status)
	} else if err := store.SetStatus(req, rsp, 100); err != nil {
		t.Fatalf("Error SetStatus: %v", err)
	} else if status := store.GetStatus(req); status == 0 {
		t.Errorf("Expected empty status; Got %v", status)
	} else {
		t.Logf("status: %v", status)
	}
}
