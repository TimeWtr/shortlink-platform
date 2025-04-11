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
	// InstanceBits 实例ID的位数
	InstanceBits = 5
)

// 全局单例模式保证只有一个node
var (
	instance *SnowflakeNode
	once     sync.Once
)

type SnowflakeNode struct {
	// ID传输通道
	ch chan int64
	// 传输通道的容量
	capacity int64
	// 节点
	node *snowflake.Node
	// 关闭通道
	stop chan struct{}
	// 单例
	once sync.Once
}

func NewGenID(dataCenterID int64, instanceID int64, capacity int64) (IDInter, error) {
	if instanceID < 1 || instanceID > 31 {
		return nil, fmt.Errorf("instanceID out of range: 0-31")
	}

	if dataCenterID < 0 || dataCenterID > 31 {
		return nil, fmt.Errorf("dataCenterID out of range: 0-31")
	}

	nodeID := (dataCenterID << InstanceBits) | instanceID
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize snowflake node, err: %v", err)
	}

	once.Do(func() {
		instance = &SnowflakeNode{
			capacity: capacity,
			once:     sync.Once{},
			node:     node,
			ch:       make(chan int64, capacity),
			stop:     make(chan struct{}),
		}

		go instance.run()
	})

	return instance, nil
}

func (s *SnowflakeNode) GetChannel() (<-chan int64, error) {
	return s.ch, nil
}

func (s *SnowflakeNode) run() {
	for {
		select {
		case <-s.stop:
			close(s.ch)
			return
		case s.ch <- s.node.Generate().Int64():
		default:
		}
	}
}

func (s *SnowflakeNode) Close() {
	s.once.Do(func() {
		close(s.stop)
	})
}
