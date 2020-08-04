package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGaugeRun(t *testing.T) {
	cmd := exec.Command("gauge", "run", "-p", "specs")
	o, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}

	expectations := []struct {
		pattern string
		count   int
		name    string
	}{
		{"BeforeSuite", 1, "BeforeSuite=1"},
		{"AfterSuite", 1, "AfterSuite=1"},
		{"BeforeSpec", 4, "BeforeSpec=4"},
		{"AfterSpec", 4, "AfterSpec=4"},
		{"BeforeScenario", 8, "BeforeScenario=8"},
		{"AfterScenario", 8, "AfterScenario=8"},
		{"ThreadName: ", 8, "ThreadCount=8"},
	}

	for _, e := range expectations {
		t.Run(e.name, func(t *testing.T) {
			got := strings.Count(string(o), e.pattern)
			if got != e.count {
				t.Errorf("Expected %s to occur %d, got %d", e.pattern, e.count, got)
			}
		})
	}

	t.Run("streams in different threads", func(t *testing.T) {
		lines := strings.Split(string(o), "\n")
		threads := make(map[string]int)
		for _, line := range lines {
			if strings.HasPrefix(line, "ThreadName: ") {
				threads[strings.ReplaceAll(line, "ThreadName: ", "")]++
			}
		}
		if len(threads) != 4 {
			t.Errorf("Expected 4 threads, got %d", len(threads))
		}
	})
}
