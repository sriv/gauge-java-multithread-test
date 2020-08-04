package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

const numStreams = 4

func TestGaugeRun(t *testing.T) {
	cmd := exec.Command("gauge", "run", "-p", "-n", fmt.Sprint(numStreams), "specs")
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
		{"BeforeSpec", numStreams, fmt.Sprintf("BeforeSpec=%d", numStreams)},
		{"AfterSpec", numStreams, fmt.Sprintf("AfterSpec=%d", numStreams)},
		{"BeforeScenario", numStreams * 2, fmt.Sprintf("BeforeScenario=%d", numStreams*2)},
		{"AfterScenario", numStreams * 2, fmt.Sprintf("AfterScenario=%d", numStreams*2)},
		{"ThreadName: ", numStreams * 2, fmt.Sprintf("ThreadCount=%d", numStreams)},
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
		if len(threads) != numStreams {
			t.Errorf("Expected %d threads, got %d", numStreams, len(threads))
		}
	})
}
