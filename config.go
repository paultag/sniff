/* {{{ Copyright 2017 Paul Tagliamonte
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License. }}} */

package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Bind    Bind
	Servers []Server
	Verbose bool
}

type Bind struct {
	Host string
	Port int
}

type Server struct {
	Default bool
	Regexp  bool
	Host    string
	Names   []string
	Port    int
}

func LoadConfig(path string) (*Config, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	config := Config{}
	return &config, json.NewDecoder(fd).Decode(&config)
}
