// Copyright 2019 Google LLC
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

package certmanager_test

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/juangodPerlego/webpackager/certchain/certmanager"
	"github.com/juangodPerlego/webpackager/internal/timeutil"
)

func TestFetchAtIntervalsWithEventFactory(t *testing.T) {
	timing := certmanager.FetchAtIntervalsWithEventFactory(
		time.Hour,
		newDummyFutureEvent,
	)

	now := time.Date(2020, time.April, 1, 15, 0, 0, 0, time.UTC)
	want := newDummyFutureEvent(
		time.Date(2020, time.April, 1, 16, 0, 0, 0, time.UTC),
	)
	timeutil.StubNowWithFixedTime(now)
	defer timeutil.ResetNow()
	if diff := cmp.Diff(want, timing.GetNextRun()); diff != "" {
		t.Errorf("GetNextRun() mismatch (-want +got):\n%s", diff)
	}
}

func TestOnSignal(t *testing.T) {
	timing := certmanager.FetchOnSignal(syscall.SIGUSR1)
	nextRun := timing.GetNextRun()
	syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	if r := waitForEvent(nextRun, defaultTimeout); r != waitSuccess {
		t.Errorf("waitForEvent() = %v, want waitSuccess", r)
	}
}
