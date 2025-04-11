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

package cache

import (
	_ "embed"
	"errors"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

//go:embed scripts/get_short_code.lua
var GetShortCodeScript string

type CInter interface {
	// Count 短码池中的预生成短码数量
	Count(ctx context.Context) (int64, error)
	// GetShortCode 从短码池中获取一个预生成短码
	GetShortCode(ctx context.Context) (string, error)
	// InsertShortCode 向短码池中新增一条预生成短码
	InsertShortCode(ctx context.Context, code string) error
	// BatchInsertShortCodes 批量向短码池中增加多条预生成短码
	BatchInsertShortCodes(ctx context.Context, codes []string) error
}

type Cache struct {
	client redis.Cmdable
}

func NewCache(client redis.Cmdable) CInter {
	return &Cache{client: client}
}

func (c *Cache) Count(ctx context.Context) (int64, error) {
	return c.client.Get(ctx, "shortCodeCount").Int64()
}

// GetShortCode 查询短码数量、获取一条可用的预生成短码、更新短码数量
func (c *Cache) GetShortCode(ctx context.Context) (string, error) {
	code := c.client.Eval(ctx, GetShortCodeScript, []string{}).String()
	if code == "" {
		return "", errors.New("short code not found")
	}

	return code, nil
}

func (c *Cache) InsertShortCode(ctx context.Context, code string) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) BatchInsertShortCodes(ctx context.Context, codes []string) error {
	//TODO implement me
	panic("implement me")
}
