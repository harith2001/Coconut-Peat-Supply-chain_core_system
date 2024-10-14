// package plugin

// import (
// 	"fmt"
// )

// // In-memory store of parent and child plugins (in production, you'd store this in MongoDB)
// var ParentPlugins = map[string]map[string]string{
// 	"GradingPlugin": {
// 		"sizeThreshold":     "10",
// 		"moistureThreshold": "15",
// 	},
// }

// var ChildPlugins = []map[string]interface{}{}

// // CreateChildPlugin creates a child plugin with customized parameters
// func CreateChildPlugin(parentPluginName string, customizedParams map[string]string) (bool, string) {
// 	// Find the parent plugin
// 	parentPlugin, exists := ParentPlugins[parentPluginName]
// 	if !exists {
// 		return false, fmt.Sprintf("Parent plugin %s not found", parentPluginName)
// 	}

// 	// Create a new child plugin by inheriting from the parent plugin
// 	childPlugin := make(map[string]interface{})
// 	childPlugin["parent"] = parentPluginName

// 	// Apply the parent plugin's parameters
// 	for key, value := range parentPlugin {
// 		childPlugin[key] = value
// 	}

// 	// Override with customized parameters
// 	for key, value := range customizedParams {
// 		if _, customizable := parentPlugin[key]; customizable {
// 			childPlugin[key] = value
// 		}
// 	}

// 	// Store the child plugin (could store in MongoDB)
// 	ChildPlugins = append(ChildPlugins, childPlugin)

//		return true, "Child plugin created successfully"
//	}
package plugin

import (
	"fmt"
)

// In-memory store of parent and child plugins (in production, you'd store this in MongoDB)
var ParentPlugins = map[string]map[string]string{
	"GradingPlugin": {
		"sizeThreshold":     "10",
		"moistureThreshold": "15",
	},
}

var ChildPlugins = []map[string]interface{}{}

// CreateChildPlugin creates a child plugin with customized parameters
func CreateChildPlugin(parentPluginName string, customizedParams map[string]string) (bool, string) {
	// Find the parent plugin
	parentPlugin, exists := ParentPlugins[parentPluginName]
	if !exists {
		return false, fmt.Sprintf("Parent plugin %s not found", parentPluginName)
	}

	// Create a new child plugin by inheriting from the parent plugin
	childPlugin := make(map[string]interface{})
	childPlugin["parent"] = parentPluginName

	// Apply the parent plugin's parameters
	for key, value := range parentPlugin {
		childPlugin[key] = value
	}

	// Override with customized parameters
	for key, value := range customizedParams {
		if _, customizable := parentPlugin[key]; customizable {
			childPlugin[key] = value
		}
	}

	// Store the child plugin (could store in MongoDB)
	ChildPlugins = append(ChildPlugins, childPlugin)

	return true, "Child plugin created successfully"
}
