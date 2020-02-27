package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	fileName = "default"

	unitConfig = map[string]string{
		"Description":         "Build by github.com/yzy613/create-systemd-service",
		"Documentation":       "",
		"Requires":            "",
		"RequiresOverridable": "",
		"Requisite":           "",
		"Conflicts":           "",
		"After":               "",
		"Before":              "",
		"OnFailure":           ""}

	serviceConfig = map[string]string{
		"Type":            "simple", // simple, exec, forking, oneshot, dbus, notify, idle
		"GuessMainPID":    "",
		"PIDFile":         "",
		"BusName":         "",   // 在 Type 是 dbus 的时候才需要
		"RemainAfterExit": "no", // no(Default), yes
		"EnvironmentFile": "",
		"ExecStartPre":    "",
		"ExecStart":       "",
		"TimeoutStartSec": "",
		"ExecStartPost":   "",
		"ExecReload":      "",
		"ExecStop":        "",
		"TimeoutStopSec":  "",
		"TimeoutSec":      "",           // 同时设置 TimeoutStartSec TimeoutStopSec
		"Restart":         "on-failure", // no, on-success, on-failure, on-abnormal, on-watchdog, on-abort, always
		"RestartSec":      "",           // Default 100ms
		"KillMode":        "",           // control-group(Default), process, mixed, none
		// not often to use
		"SuccessExitStatus":        "",
		"RuntimeMaxSec":            "",
		"WatchdogSec":              "",
		"RestartPreventExitStatus": "",
		"RestartForceExitStatus":   "",
		"RootDirectoryStartOnly":   "", // Default no (yes|no)
		"NonBlocking":              "", // Default no (yes|no)
		"NotifyAccess":             "", // none(Default), main, exec, all
		"Sockets":                  "",
		"FileDescriptorStoreMax":   "", // Default 0
		"USBFunctionDescriptors":   "",
		"USBFunctionStrings":       "",
		"StandardOutput":           "",
		"StandardError":            ""}

	installConfig = map[string]string{
		"WantedBy": "multi-user.target",
		"Alias":    "",
		"Also":     ""}

	helper = map[string]string{
		"ls":  "Usage: ls [name|unit}",
		"set": "usage: set [name|unit] [target] [content]"}
)

func appendValuableStr(key, value string, src *string) {
	if value != "" {
		*src += key + "=" + value + "\n"
	}
}

func output() {
	outStr := "[Unit]\n"
	for key, value := range unitConfig {
		appendValuableStr(key, value, &outStr)
	}
	outStr += "\n[Service]\n"
	for key, value := range serviceConfig {
		appendValuableStr(key, value, &outStr)
	}
	outStr += "\n[Install]\n"
	for key, value := range installConfig {
		appendValuableStr(key, value, &outStr)
	}
	outputFile := "./" + fileName + ".service"
	ioutil.WriteFile(outputFile, []byte(outStr), 0644)
}

func listContent(cmdArr []string) {
	arrLen := len(cmdArr) // 命令个数
	nowStep := 1          // 记录步骤
	switch cmdArr[0] {
	case "name":
		fmt.Println(fileName)
	case "unit":
		if nowStep >= arrLen {
			for key, value := range unitConfig {
				fmt.Println(key, "=", value)
			}
		} else {
			fmt.Println(cmdArr[nowStep], "=", unitConfig[cmdArr[nowStep]])
		}
	default:
		fmt.Println("未定义的命令")
	}
}

func setContent(cmdArr []string) {
	arrLen := len(cmdArr) // 命令个数
	nowStep := 1          // 记录步骤
	switch cmdArr[0] {
	case "name":
		if nowStep >= arrLen {
			fmt.Println(helper["set"])
			break
		}
		fileName = cmdArr[1]
	case "unit":
		nowStep++
		if nowStep >= arrLen {
			fmt.Println(helper["set"])
			break
		}
		for _, value := range cmdArr[nowStep:arrLen] {
			unitConfig[cmdArr[nowStep-1]] += value
		}
	default:
		fmt.Println("未定义的命令")
	}
}

func main() {
	for {
		// 仿终端
		fmt.Print("[create-systemd-service]$ ")
		// 读取一行
		gets := bufio.NewReader(os.Stdin)
		inputStr, err := gets.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// 对命令分片
		front := 0
		var cmdArr []string
		strLen := len(inputStr)
		if strLen == 0 || strLen == 1 { // 处理没有键入任何命令
			continue
		}
		for i, content := range inputStr {
			if content == ' ' || i == strLen-1 || content == '\r'{
				cmdArr = append(cmdArr, inputStr[front:i])
				if content == ' ' {
					front = i + 1
				}
			}
		}
		arrLen := len(cmdArr)
		signExit := false // 退出信号
		nowStep := 1      // 记录步骤
		// 处理命令
		switch cmdArr[0] {
		case "exit":
			signExit = true
		case "help":
			for key, value := range helper {
				fmt.Println(key, value)
			}
		case "set":
			if nowStep >= arrLen {
				fmt.Println(helper["set"])
				break
			}
			setContent(cmdArr[1:arrLen])
		case "ls":
			if nowStep >= arrLen {
				fmt.Println(helper["ls"])
				break
			}
			listContent(cmdArr[1:arrLen])
		case "build":
			if fileName == "default" {
				fmt.Println("你还没为该service文件命名， 是否返回命名？(y|n)")
				var yesOrNo string
				fmt.Scan(&yesOrNo)
				if yesOrNo != "n" {
					continue
				}
				output()
			}
		default:
			fmt.Println("未定义的命令")
		}
		// 退出程序
		if signExit {
			break
		}
	}
}
