// Copyright © 2017 The Kubicorn Authors
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

package fileresource

import (
	"fmt"

	"github.com/kubicorn/kubicorn/pkg/version"
)

const (
	githubProtocol = "https://"
	githubRepo     = "raw.githubusercontent.com/kubicorn/bootstrap"
)

var (
	githubBranch   = version.GetVersion().Version
)

// getGitHubUrl will build a query-able URL from a bootstrap script that we can parse in at runtime.
// Example URL:
// https://raw.githubusercontent.com/kubicorn/bootstrap/master/amazon_k8s_centos_7_master.sh
func getGitHubUrl(bootstrapScript string) string {
	// TODO(@xmudrii): this is a bad solution but I want this working.
	gitHubUrl := fmt.Sprintf("%s%s/%s/%s", githubProtocol, githubRepo, githubBranch, bootstrapScript[10:])
	return gitHubUrl
}
