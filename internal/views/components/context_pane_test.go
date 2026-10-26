package components

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestContextPaneClosesWhenTappingOutsideOnMobile(t *testing.T) {
	var output bytes.Buffer
	if err := ContextPane().Render(context.Background(), &output); err != nil {
		t.Fatal(err)
	}

	rendered := output.String()
	for _, expected := range []string{
		"@click.away=",
		"window.innerWidth < 1024",
		"$store.contextPane.closePane()",
		"closest('.pull-request')",
		"closest('.filter-panel')",
		"closest('.mobile-dock')",
	} {
		if !strings.Contains(rendered, expected) {
			t.Errorf("context pane output does not contain %q", expected)
		}
	}
}
