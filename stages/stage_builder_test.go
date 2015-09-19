/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//Package stages contains functionality for managing stage lifecycle
package stages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCommandStage(t *testing.T) {
	builder := new(StageBuilder)
	stage := builder.NewStage("command").SetName("show dir").SetTarget("ls -la").Build()
	assert.Equal(t, "show dir", stage.GetStageName())
}