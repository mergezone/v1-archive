package pages

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestHomePageHasNoSidebarAndProminentlyPlacesInteractiveSearchAfterCTA(t *testing.T) {
	var output bytes.Buffer
	if err := HomePage("https://example.test", "", "", "abc123").Render(context.Background(), &output); err != nil {
		t.Fatal(err)
	}
	rendered := output.String()
	if strings.Contains(rendered, `class="aside-column`) {
		t.Error("home page should not render the sidebar")
	}
	if strings.Contains(rendered, `id="mobile-dock"`) {
		t.Error("home page should keep the prominent repository search in the main content, not a mobile dock")
	}
	cta := strings.Index(rendered, "and contributor patterns at a glance.")
	search := strings.Index(rendered, `class="home-search`)
	if cta < 0 || search < 0 || search < cta {
		t.Fatal("interactive repository search should follow the call-to-action")
	}
	if !strings.Contains(rendered[search:], `@click="startOwnerEdit()"`) {
		t.Error("home page repository search should remain interactive")
	}
}
