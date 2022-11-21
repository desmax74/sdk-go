// Copyright 2022 The Serverless Workflow Specification Authors
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

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludePaths(t *testing.T) {
	assert.NotNil(t, IncludePaths())
	assert.True(t, len(IncludePaths()) > 0)

	// update include paths
	paths := []string{"/root", "/path"}
	SetIncludePaths(paths)
	assert.Equal(t, IncludePaths(), paths)

	assert.PanicsWithError(t, "1 must be an absolute file path", assert.PanicTestFunc(func() {
		SetIncludePaths([]string{"1"})
	}))
}
