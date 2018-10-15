package confluence

import (
	"testing"
)

func TestContentRequestPayload(t *testing.T){
  spaceId := "test"
  version := 1
  title := "Hello!"
  html := "<p>Hello, World</p>"

  req := ContentRequestPayload(spaceId, version, title, html)

  if 2 != req.Version.Number {
    t.Errorf("Expected to increment the version number")
  }

  if spaceId != req.Space.Key {
    t.Errorf("expected to set Space.Key (%s), got %s", spaceId, req.Space.Key)
  }

  if title != req.Title {
    t.Errorf("expected to set Title (%s), got %s", title, req.Title)
  }

  if html != req.Body.StorageView.Value {
    t.Errorf("expected to set req.Body.StorageView.Value (%s), got %s", html, req.Title)
  }
}
