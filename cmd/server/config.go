// Copyright (c) 2018-2019 Sylabs, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config hold all possible parameters that are used to
// tune Singularity CRI default behaviour.
type Config struct {
	// ListenSocket is a unix socket to serve CRI requests on.
	ListenSocket string `yaml:"listenSocket"`
	// StorageDir is a directory to store all pulled images in.
	StorageDir string `yaml:"storageDir"`
	// StreamingURL is an address to serve streaming requests on (exec, attach, portforward).
	StreamingURL string `yaml:"streamingURL"`
	// CNIBinDir is a directory to look for CNI plugin binaries.
	CNIBinDir string `yaml:"cniBinDir"`
	// CNIConfDir is a directory to look for CNI network configuration files.
	CNIConfDir string `yaml:"cniConfDir"`
	// BaseRunDir is a directory to store currently running pods and containers.
	BaseRunDir string `yaml:"baseRunDir"`
}

var defaultConfig = Config{
	ListenSocket: "/var/run/singularity.sock",
	StorageDir:   "/var/lib/singularity",
	StreamingURL: "127.0.0.1:12345",
	CNIBinDir:    "/usr/local/libexec/singularity/cni",
	CNIConfDir:   "/usr/local/etc/singularity/network",
	BaseRunDir:   "/var/run/singularity",
}

func parseConfig(path string) (Config, error) {
	var config Config

	f, err := os.Open(path)
	if err != nil {
		return config, fmt.Errorf("could not open config file: %v", err)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("could not decode config: %v", err)
	}
	return mergeConfig(config, defaultConfig), nil
}

func mergeConfig(config, defaultConfig Config) Config {
	if config.ListenSocket == "" {
		config.ListenSocket = defaultConfig.ListenSocket
	}
	if config.StorageDir == "" {
		config.StorageDir = defaultConfig.StorageDir
	}
	if config.StreamingURL == "" {
		config.StreamingURL = defaultConfig.StreamingURL
	}
	if config.CNIBinDir == "" {
		config.CNIBinDir = defaultConfig.CNIBinDir
	}
	if config.CNIConfDir == "" {
		config.CNIConfDir = defaultConfig.CNIConfDir
	}
	if config.BaseRunDir == "" {
		config.BaseRunDir = defaultConfig.BaseRunDir
	}
	return config
}
