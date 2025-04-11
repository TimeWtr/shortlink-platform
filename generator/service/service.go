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

package service

import (
	"errors"

	intrv1 "github.com/TimeWtr/shortlink-platform/generator/api/proto/gen/intr.v1"
	"github.com/TimeWtr/shortlink-platform/generator/domain"
	"golang.org/x/net/context"
)

type URLServiceInter interface {
	// GenerateURL 生成单条URL
	GenerateURL(ctx context.Context, req *intrv1.URLRequest) (domain.URLResponse, error)
	// BatchGenerateURL 批量生成URL
	BatchGenerateURL(ctx context.Context, req *intrv1.URLRequest) ([]domain.URLResponse, error)
}

const RetryCounts = 5

type Service struct {
	// ID获取的通道
	idCh <-chan int64
}

func NewService(idCh <-chan int64) URLServiceInter {
	return &Service{
		idCh: idCh,
	}
}

func (s *Service) GenerateURL(ctx context.Context, req *intrv1.URLRequest) (domain.URLResponse, error) {
	// 获取分布式ID
	var id int64
	counter := 0
	for counter < RetryCounts {
		select {
		case <-ctx.Done():
			return domain.URLResponse{}, ctx.Err()
		case id = <-s.idCh:
			break
		default:
			counter++
		}
	}
	if id == 0 {
		return domain.URLResponse{}, errors.New("failed to get id")
	}

	return domain.URLResponse{}, nil
}

func (s *Service) BatchGenerateURL(ctx context.Context, req *intrv1.URLRequest) ([]domain.URLResponse, error) {
	//TODO implement me
	panic("implement me")
}
