// Copyright © 2017 Casey Marshall
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

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tempFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("tempfile: %v", err)
	}
	defer f.Close()
	return f.Name()
}

func TestConfigDefaults(t *testing.T) {
	fpath := tempFile(t)
	defer os.Remove(fpath)
	cfg, err := ReadFile(fpath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !strings.HasPrefix(cfg.Node.Agent.SocksAddr, "127.0.0.1:9") {
		t.Errorf("default Node.Agent.SocksAddr %s", cfg.Node.Agent.SocksAddr)
	}
	if !strings.HasPrefix(cfg.Node.Agent.ControlAddr, "127.0.0.1:9") {
		t.Errorf("default Node.Agent.ControlAddr %s", cfg.Node.Agent.ControlAddr)
	}
}

func TestRoundTrip(t *testing.T) {
	fpath := tempFile(t)
	defer os.Remove(fpath)
	config := Config{
		Dir:  filepath.Dir(fpath),
		Path: fpath,
		Node: Node{
			Agent: Agent{
				TorBinaryPath:  "/usr/bin/tor",
				SocksAddr:      "127.0.0.1:9050",
				ControlAddr:    "127.0.0.1:9051",
				TorrcPath:      "/path/to/torrc",
				TorDataDir:     "/path/to/tor/data",
				TorServicesDir: "/path/to/tor/services",
				ControlCookie:  "yum",
			},
			Service: Service{
				Exports: []Export{{
					LocalAddr: "127.0.0.1:80",
					Port:      80,
				}},
				Clients: []Client{{
					Name:    "bob",
					Address: "qwertyuiop.onion",
				}},
			},
		},
	}
	err := WriteFile(&config, fpath)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}
	config2, err := ReadFile(fpath)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	assert.Equal(t, &config, config2)
}
