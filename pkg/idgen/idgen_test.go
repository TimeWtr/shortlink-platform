// Copyright 2025 TimeWtr
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package idgen

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSnowflakeID_GetIDChannel(t *testing.T) {
	idCenter, err := NewGenID(9, 23, 10000)
	assert.NoError(t, err)
	ch, err := idCenter.GetIDChannel()
	assert.NoError(t, err)

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	sig := make(chan struct{})
	go func() {
		for range ticker.C {
			select {
			case <-sig:
				t.Log("received signal")
				return
			case id := <-ch:
				t.Logf("Get id: %d\n", id)
			default:
			}
		}
	}()
	time.Sleep(10 * time.Second)
	close(sig)
	idCenter.Close()
}

func TestSnowflakeID_Close(t *testing.T) {
	idCenter, err := NewGenID(9, 23, 10000)
	assert.NoError(t, err)
	ch, err := idCenter.GetIDChannel()
	assert.NoError(t, err)

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	sig := make(chan struct{})
	go func() {
		for range ticker.C {
			select {
			case <-sig:
				t.Log("received signal")
				return
			case id := <-ch:
				t.Logf("Get id: %d\n", id)
			default:
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(sig)
	idCenter.Close()
	time.Sleep(time.Second)
}

func BenchmarkSnowflakeID_GetIDChannel_Log(b *testing.B) {
	idCenter, err := NewGenID(10, 31, 10000)
	assert.NoError(b, err)
	ch, err := idCenter.GetIDChannel()
	assert.NoError(b, err)
	defer idCenter.Close()

	for i := 0; i < b.N; i++ {
		id := <-ch
		b.Logf("Get id: %d\n", id)
	}
}

func BenchmarkSnowflakeID_GetIDChannel_NoLog(b *testing.B) {
	idCenter, err := NewGenID(0, 31, 10000)
	assert.NoError(b, err)
	ch, err := idCenter.GetIDChannel()
	assert.NoError(b, err)
	defer idCenter.Close()

	for i := 0; i < b.N; i++ {
		<-ch
	}
}
