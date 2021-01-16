/*
MIT License

Copyright (c) 2020-2021 Kazuhito Suda

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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type cmdDef func(int, []string) error

var cmdTable map[string]cmdDef

func initCmdTable() {
	cmdTable = map[string]cmdDef{
		"```":   compareCmd,
		"halt":  haltCmd,
		"http":  httpCmd,
		"ngsi":  ngsiCmd,
		"print": printCmd,
		"sleep": sleepCmd,
		"wait":  waitCmd,
	}
}

func compareCmd(line int, args []string) error {
	const funcName = "compareCmd"

	var err error

	line = line - len(args) + 1

	if len(args[0]) < 3 {
		return &ngsiCmdError{funcName, 1, "expected code error", nil}
	}
	expectedCode := args[0][3:]

	v, ok := val["?"]
	if !ok {
		return &ngsiCmdError{funcName, 2, "acttual code error", nil}
	}
	actualCode := v[0]

	if expectedCode != actualCode {
		fmt.Printf("Exit code error, expected:%s, actual:%s\n", expectedCode, actualCode)
		err = &ngsiCmdError{funcName, 3, fmt.Sprintf("Exit code error, expected:%s, actual:%s", expectedCode, actualCode), nil}
	}

	expected := args[1 : len(args)-1]
	actual := val["$"]
	if actual == nil {
		actual = []string{}
	}

	if err := diffLines(line, expected, actual); err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	return err
}

func haltCmd(line int, args []string) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	fmt.Println("halt")

	<-sig

	fmt.Println("resume")

	return nil
}

func httpCmd(line int, args []string) error {
	const funcName = "httpCmd"

	if len(args) < 2 {
		return &ngsiCmdError{funcName, 1, "http error", nil}
	}

	url := args[2]
	if !isHTTP(url) {
		return &ngsiCmdError{funcName, 2, "url error: " + url, nil}
	}

	switch args[1] {
	default:
		return &ngsiCmdError{funcName, 3, "http verb error", nil}
	case "get":
		return httpRequest(http.MethodGet, nil, args)
	case "post":
		if len(args) < 4 {
			return &ngsiCmdError{funcName, 4, "http post url --data \"{\"data\":\"post data\"}", nil}
		}
		if args[3] != "--data" {
			return &ngsiCmdError{funcName, 5, "http post url --data \"{\"data\":\"post data\"}", nil}
		}
		header := map[string]string{"Content-Type": "application/json"}
		return httpRequest(http.MethodPost, header, args)
	case "delete":
		return httpRequest(http.MethodDelete, nil, args)
	}

}

func ngsiCmd(line int, args []string) error {
	const funcName = "ngsiCmd"

	if *gArgs {
		for i, s := range args {
			fmt.Printf("%s001 %d: %s\n", funcName, i, s)
		}
	}
	param := []string{}
	if *gNgsiConfig != "" {
		param = append(param, "--config", *gNgsiConfig)
	}
	if *gNgsiCache != "" {
		param = append(param, "--cache", *gNgsiCache)
	}
	param = append(param, args[1:]...)

	cmd := exec.Command(args[0], param...)
	cmd.Stderr = nil
	rc := "0"

	result, err := cmd.Output()

	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			result = e.Stderr
		}
		rc = strconv.Itoa(cmd.ProcessState.ExitCode())
	}
	val["?"] = []string{rc}

	if len(result) > 0 {
		s := strings.TrimRight(string(result), "\n")
		val["$"] = strings.Split(s, "\n")
	} else {
		val["$"] = []string{}
	}

	return nil
}

func printCmd(line int, args []string) error {
	if len(args) == 2 {
		fmt.Println(args[1])
	}
	return nil
}

func sleepCmd(line int, args []string) error {
	const funcName = "sleepCmd"

	if len(args) == 2 {
		v := strings.Split(args[1], ".")
		if len(v) > 2 {
			return &ngsiCmdError{funcName, 1, "value error: " + args[1], nil}
		}
		i, err := strconv.Atoi(v[0])
		if err != nil {
			return &ngsiCmdError{funcName, 2, "value error: " + v[0], nil}
		}
		t := time.Second * time.Duration(i)
		if len(v) == 2 && len(v[1]) == 1 {
			i, err = strconv.Atoi(v[1])
			if err != nil {
				return &ngsiCmdError{funcName, 3, "value error: " + v[1], nil}
			}
			t += time.Millisecond * time.Duration(i*100)
		}
		time.Sleep(t)
		return nil
	}

	return &ngsiCmdError{funcName, 4, "param error" + args[1], nil}
}

func waitCmd(line int, args []string) (err error) {
	const funcName = "waitCmd"

	retry := 600
	if len(args) == 2 {
		if !isHTTP(args[1]) {
			return &ngsiCmdError{funcName, 1, "url error: " + args[1], nil}
		}
		fmt.Printf("Waiting for response from %s\n", args[1])
		for {
			var res *http.Response
			res, err = http.Get(args[1])
			if err != nil {
				retry--
				if retry == 0 {
					return &ngsiCmdError{funcName, 2, "no response from " + args[1], nil}
				}
				time.Sleep(time.Second * time.Duration(1))
				continue
			}
			defer func() { setNewError(funcName, 3, res.Body.Close(), &err) }()

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
			if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
				return nil
			}
		}
	}
	return &ngsiCmdError{funcName, 4, "param error" + args[1], nil}
}

func httpRequest(method string, header map[string]string, args []string) (err error) {
	const funcName = "httpRequest"

	b := []byte(nil)
	if method == http.MethodPost {
		b = []byte(args[4])
	}
	var req *http.Request
	req, err = http.NewRequest(method, args[2], bytes.NewBuffer(b))
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	defer func() { setNewError(funcName, 3, res.Body.Close(), &err) }()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	status := "0"
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		status = "1"
	}
	val["?"] = []string{status}

	if len(b) > 0 {
		s := strings.TrimRight(string(b), "\n")
		val["$"] = strings.Split(s, "\n")
	} else {
		val["$"] = []string{}
	}

	return nil
}

func isHTTP(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
