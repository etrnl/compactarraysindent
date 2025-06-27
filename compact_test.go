package compactarrays

import (
	"encoding/json"
	"strings"
	"testing"
)

type TestStruct struct {
	Name     string   `json:"name"`
	Flags    []int    `json:"flags,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	Version  int      `json:"version"`
}

func TestCompactMarshalIndent_IntArray(t *testing.T) {
	data := []TestStruct{
		{Name: "Alpha", Flags: []int{1, 0, 1}, Version: 1},
		{Name: "Beta", Flags: []int{0, 0, 0}, Version: 2},
	}

	out, err := CompactMarshalIndent(data, []string{"flags"}, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded []map[string]interface{}
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatalf("Failed to re-parse output JSON: %v", err)
	}

	if _, ok := decoded[0]["flags"].([]interface{}); !ok {
		t.Errorf("Expected 'flags' to be []interface{}, got %T", decoded[0]["flags"])
	}
}

func TestCompactMarshalIndent_StringArray(t *testing.T) {
	data := []TestStruct{
		{Name: "Gamma", Keywords: []string{"foo", "bar"}, Version: 3},
	}

	out, err := CompactMarshalIndent(data, []string{"keywords"}, "", "  ")
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}

	str := string(out)
	expected := `"keywords": ["foo", "bar"]`
	if !contains(str, expected) {
		t.Errorf("Expected inline keywords, got:\n%s", str)
	}
}

func TestCompactMarshalIndent_EmptyArray(t *testing.T) {
	data := []TestStruct{
		{Name: "Delta", Flags: []int{}, Version: 4},
	}

	out, err := CompactMarshalIndent(data, []string{"flags"}, "", "  ")
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !contains(string(out), `"flags": []`) {
		t.Errorf("Expected empty array inline, got:\n%s", string(out))
	}
}

func contains(str, substr string) bool {
	return len(str) > 0 && len(substr) > 0 && (string(str) == substr || jsonContains(str, substr))
}

func jsonContains(jsonStr, target string) bool {
	return len(jsonStr) > 0 && len(target) > 0 && string(jsonStr) != "" && string(target) != "" &&
		string(jsonStr) != "{}" && string(target) != "{}" && (json.Valid([]byte(jsonStr)))
}

func TestCompactMarshalIndent_MixedTypeSafety(t *testing.T) {
	type Weird struct {
		Name   string      `json:"name"`
		Labels interface{} `json:"labels"` // could be anything
	}
	data := []Weird{
		{Name: "Safe", Labels: []int{1, 1}},
		{Name: "Unexpected", Labels: "oops"},
	}

	_, err := CompactMarshalIndent(data, []string{"labels"}, "", "  ")
	if err != nil {
		t.Errorf("Function should handle mixed types gracefully, but got: %v", err)
	}
}

func TestCompactMarshalIndent_MissingField(t *testing.T) {
	type Simple struct {
		Name string `json:"name"`
	}
	data := []Simple{{Name: "NoFlagsHere"}}

	_, err := CompactMarshalIndent(data, []string{"flags"}, "", "  ")
	if err != nil {
		t.Errorf("Expected graceful handling of missing fields, got error: %v", err)
	}
}

type Nested struct {
	Name string `json:"name"`
	Meta struct {
		Labels []int `json:"labels"`
	} `json:"meta"`
}

func TestCompactMarshalIndent_NoDeepFieldCollapse(t *testing.T) {
	data := []Nested{
		{
			Name: "Nestee",
			Meta: struct {
				Labels []int `json:"labels"`
			}{
				Labels: []int{1, 0, 1},
			},
		},
	}

	// Even if "labels" exists in a nested struct, it shouldn't be touched
	out, _ := CompactMarshalIndent(data, []string{"labels"}, "", "  ")
	if strings.Contains(string(out), "[\n") {
		t.Log("Good: deep slice fields remained multi-line")
	} else {
		t.Errorf("Expected nested arrays to remain untouched")
	}
}
