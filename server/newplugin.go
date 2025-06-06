package server

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	pbv "github.com/harith2001/Coconut-Peat-Supply-chain_core_system/proto"
)

type NewPlugin struct {
	pbv.UnimplementedNewPluginServiceServer
}

func (s *NewPlugin) NewPluginCreate(ctx context.Context, req *pbv.NewPluginCreateRequest) (*pbv.NewPluginCreateResponse, error) {
	filename := req.FileName
	filedata := req.FileData
	savePath := filepath.Join("plugins", filename)

	// Ensure the directory exists
	dir := filepath.Dir(savePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Printf("Failed to create directory: %v", err)
		return &pbv.NewPluginCreateResponse{
			Success: false,
			Message: "Failed to create directory",
		}, err
	}

	// Save the file
	err = os.WriteFile(savePath, filedata, 0644)
	if err != nil {
		log.Printf("Failed to save the file: %v", err)
		return &pbv.NewPluginCreateResponse{
			Success: false,
			Message: "Failed to save the file",
		}, err
	}

	//run the plugin.sh command file to unzip, install and run docker container
	pluginName := filename[:len(filename)-len(filepath.Ext(filename))] // remove .zip
	cmd := exec.Command("/bin/bash", "plugin.sh", pluginName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Printf("Failed to execute command file: %v", err)
		return &pbv.NewPluginCreateResponse{
			Success: false,
			Message: "Failed to execute command file",
		}, err
	}

	log.Printf("File %s uploaded successfully", filename)
	return &pbv.NewPluginCreateResponse{
		Success: true,
		Message: "File uploaded and command executed successfully",
	}, nil
}
