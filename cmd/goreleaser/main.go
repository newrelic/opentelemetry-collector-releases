// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/newrelic/opentelemetry-collector-releases/cmd/goreleaser/internal"
)

var distsFlag = flag.String("d", "", "Collector distributions(s) to build, comma-separated")

func main() {
	flag.Parse()

	if len(*distsFlag) == 0 {
		log.Fatal("no distributions to build")
	}
	dists := strings.Split(*distsFlag, ",")

	project := internal.Generate(internal.ImagePrefixes, dists)

	if err := yaml.NewEncoder(os.Stdout).Encode(&project); err != nil {
		log.Fatal(err)
	}
}
