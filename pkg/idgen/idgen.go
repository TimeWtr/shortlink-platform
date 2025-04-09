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
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

const (
	// Capacity channel容量
	Capacity = 10000
	// InstanceBits 实例ID的位数
	InstanceBits = 5
)

type SnowflakeID struct {
	// 当前的机房编号
	DataCenterID int64
	// 当前实例ID
	InstanceID int64
	// ID传输通道
	ch chan any
	// 传输通道的容量
	capacity int64
	// 关闭通道
	stop chan struct{}
	// 单例
	once sync.Once
}

func NewGenID(dataCenterID int64, instanceID int64, capacity int64) (IDInter, error) {
	if instanceID < 1 || instanceID > 32 {
		return nil, fmt.Errorf("instanceID out of range: 0-31")
	}

	if dataCenterID < 0 || dataCenterID > 31 {
		return nil, fmt.Errorf("dataCenterID out of range: 0-31")
	}

	return &SnowflakeID{
		DataCenterID: dataCenterID,
		InstanceID:   instanceID,
		capacity:     capacity,
		once:         sync.Once{},
	}, nil
}

func (s *SnowflakeID) GetIDChannel() (<-chan any, error) {
	nodeID := (s.DataCenterID << InstanceBits) | s.InstanceID
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	// 初始化通道
	s.ch = make(chan any, s.capacity)
	s.stop = make(chan struct{})

	go func() {
		for {
			select {
			case <-s.stop:
				fmt.Println("stopped!")
				return
			default:
				// 这里需要阻塞，如果ID已经满了，就不需要生成
				s.ch <- node.Generate().Int64()
			}
		}
	}()

	return s.ch, nil
}

func (s *SnowflakeID) Close() {
	s.once.Do(func() {
		if s.stop != nil {
			close(s.stop)
		}
	})
}
