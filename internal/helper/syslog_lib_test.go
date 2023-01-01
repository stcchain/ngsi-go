/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package helper

import (
	"errors"
	"io"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestSyslogLibNew(t *testing.T) {
	s := MockSyslogLib{}

	w, err := s.New()

	if assert.NoError(t, err) {
		assert.NotEqual(t, (io.Writer)(nil), w)
	}
}

func TestSyslogLibNewError(t *testing.T) {
	s := MockSyslogLib{Err: errors.New("syslog error")}

	_, err := s.New()

	if assert.Error(t, err) {
		assert.Equal(t, "syslog error", err.Error())
	}
}
