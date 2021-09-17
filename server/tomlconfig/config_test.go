// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tomlconfig_test

import (
	"testing"

	"github.com/juangodPerlego/webpackager/server/tomlconfig"
)

func TestReadFromFile(t *testing.T) {
	const filename = "../../cmd/webpkgserver/webpkgserver.example.toml"

	_, err := tomlconfig.ReadFromFile(filename)
	if err != nil {
		t.Errorf("ReadFromFile(%q) = error(%q), want success", filename, err)
	}
}
