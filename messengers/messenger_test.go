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
package pipelines

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walter-cd/walter-integration-test/utils"
)

// NOTE: environment variables (SLACK_CHANNEL, SLACK_URL etc) need to be set before running the following tests. 
func TestRunWalterWithSlackReport(t *testing.T) {
	assert.Equal(t, true, utils.RunWalter("pipeline_slack.yml").IsSucceed)
}

func TestRunWalterWithSlackReportWithSuppress(t *testing.T) {
	assert.Equal(t, true, utils.RunWalter("pipeline_slack_suppress.yml").IsSucceed)
}

func TestRunWalterWithHipChatReport(t *testing.T) {
	assert.Equal(t, true, utils.RunWalter("pipeline_hipchat.yml").IsSucceed)
}

func TestRunWalterWithHipChatV2Report(t *testing.T) {
	assert.Equal(t, true, utils.RunWalter("pipeline_hipchat_v2.yml").IsSucceed)
}
