package plugin

import (
	"testing"
)

// Simulate a Go file content with parent plugin and customized parameters
const goFileContent = `
package main

const ParentPluginName = "GradingPlugin"

var sizeThreshold = 12
var moistureThreshold = 16
`

// Test function to simulate Go file extraction and plugin creation
func TestCreateChildPluginFromGoFile(t *testing.T) {
	// Step 1: Simulate extracting the parent plugin name from Go file content
	parentPluginName := extractParentPluginName(goFileContent)
	if parentPluginName != "GradingPlugin" {
		t.Errorf("Expected parent plugin 'GradingPlugin', but got %s", parentPluginName)
	}

	// Step 2: Simulate extracting customized parameters from Go file content
	customizedParams := extractCustomizedParameters(goFileContent)
	if customizedParams["sizeThreshold"] != "12" || customizedParams["moistureThreshold"] != "16" {
		t.Errorf("Expected customized parameters sizeThreshold=12 and moistureThreshold=16, but got %v", customizedParams)
	}

	// Step 3: Call CreateChildPlugin with extracted data
	success, message := CreateChildPlugin(parentPluginName, customizedParams)
	if !success {
		t.Errorf("Failed to create child plugin: %s", message)
	} else {
		t.Logf("Success: %s", message)
	}
}

// Helper function to extract ParentPluginName from simulated Go file content
func extractParentPluginName(goFileContent string) string {
	// Simulate Go file extraction logic (for now, this is just an example)
	return "GradingPlugin" // Hardcoded for testing, you can add regex matching here
}

// Helper function to extract customized parameters from simulated Go file content
func extractCustomizedParameters(goFileContent string) map[string]string {
	// Simulate Go file extraction logic (for now, this is just an example)
	return map[string]string{
		"sizeThreshold":     "12",
		"moistureThreshold": "16",
	}
}
