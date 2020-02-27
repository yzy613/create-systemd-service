package main

import (
	"fmt"
)

var (
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
		"BusName":         "", // 在 Type 是 dbus 的时候才需要
		"RemainAfterExit": "no",
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
)

func main() {
	fmt.Println(123)

}
