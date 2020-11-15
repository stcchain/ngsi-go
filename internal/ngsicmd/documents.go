/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

package ngsicmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func documents(c *cli.Context) error {
	const funcName = "documents"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	fmt.Fprintln(ngsi.StdWriter, "English:")
	fmt.Fprintln(ngsi.StdWriter, "  https://fiware-orion.readthedocs.io/")
	fmt.Fprintln(ngsi.StdWriter, "  https://telefonicaid.github.io/fiware-orion/api/v2/stable/")
	fmt.Fprintln(ngsi.StdWriter, "  https://ngsi-go.letsfiware.jp/")
	fmt.Fprintln(ngsi.StdWriter, "Japanese:")
	fmt.Fprintln(ngsi.StdWriter, "  https://fiware-orion.letsfiware.jp/")
	fmt.Fprintln(ngsi.StdWriter, "  https://open-apis.letsfiware.jp/fiware-orion/api/v2/stable/")

	return nil
}
