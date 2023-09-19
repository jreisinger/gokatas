package promyaml_test

import (
	"testing"
	"time"

	"github.com/jreisinger/gokatas/promyaml"

	"github.com/google/go-cmp/cmp"
)

func TestConfigFrom_CorrectlyParsesYAMLData(t *testing.T) {
	t.Parallel()
	want := promyaml.Config{
		Global: promyaml.GlobalConfig{
			ScrapeInterval:     15 * time.Second,
			EvaluationInterval: 30 * time.Second,
			ScrapeTimeout:      10 * time.Second,
			ExternalLabels: map[string]string{
				"monitor": "codelab",
				"foo":     "bar",
			},
		},
	}
	got, err := promyaml.ConfigFrom("testdata/prom.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
