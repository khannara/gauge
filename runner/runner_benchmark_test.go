// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package runner

import (
	"github.com/getgauge/common"
	"github.com/getgauge/gauge/config"
	"github.com/getgauge/gauge/conn"
	"github.com/getgauge/gauge/gauge_messages"
	"github.com/getgauge/gauge/manifest"
	"os"
	"testing"
)

func BenchmarkLanguageRunner(b *testing.B) {
	m := &manifest.Manifest{Language: "js"}
	c := make(chan bool)
	os.Setenv(common.GaugePortEnvName, "1234")
	os.Setenv(common.GaugeProjectRootEnv, "C:\\Temp\\js\\")
	r, _ := Start(m, os.Stdout, c, false)
	for n := 0; n < b.N; n++ {
		msg := &gauge_messages.Message{
			MessageType: gauge_messages.Message_StepPositionsRequest,
			MessageId:   int64(n),
			StepPositionsRequest: &gauge_messages.StepPositionsRequest{
				FilePath: "specs\\example.spec",
			},
		}
		res, _ := conn.GetResponseForMessageWithTimeout(msg, r.Connection(), config.RunnerRequestTimeout())
		b.Log(res.MessageType)
	}
}
