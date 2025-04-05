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

package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debugf(string, ...Field)
	Infof(string, ...Field)
	Warnf(string, ...Field)
	Errorf(string, ...Field)
}

type ZapLogger struct {
	l zap.SugaredLogger
}

func NewZapLogger(l zap.SugaredLogger) Logger {
	return &ZapLogger{l: l}
}

func (z *ZapLogger) Debugf(s string, fields ...Field) {
	z.l.Debugf(s, z.transfer(fields...))
}

func (z *ZapLogger) Infof(s string, fields ...Field) {
	z.l.Infof(s, z.transfer(fields...))
}

func (z *ZapLogger) Warnf(s string, fields ...Field) {
	z.l.Warnf(s, z.transfer(fields...))
}

func (z *ZapLogger) Errorf(s string, fields ...Field) {
	z.l.Errorf(s, z.transfer(fields...))
}

func (z *ZapLogger) transfer(fields ...Field) []zap.Field {
	res := make([]zap.Field, len(fields))
	for _, field := range fields {
		res = append(res, zap.Any(field.Key, field.Val))
	}

	return res
}

type Field struct {
	Key string
	Val any
}

type NopLogger struct{}

func NewNopLogger() Logger {
	return &NopLogger{}
}

func (n *NopLogger) Debugf(s string, field ...Field) {}

func (n NopLogger) Infof(s string, field ...Field) {}

func (n NopLogger) Warnf(s string, field ...Field) {}

func (n NopLogger) Errorf(s string, field ...Field) {}
