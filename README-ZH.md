## 项目描述

该项目旨在提供一个便捷的方式列出系统中的块设备信息。通过使用 Go 语言编写的 `lsblk.go` 和 `lsblk_linux.go`
文件，该项目实现了对块设备的检索和显示功能。此外，项目根目录包含一个 `README.md` 文件，可能用于提供项目的总体介绍、使用说明和其它相关信息。

## 核心功能

- 列出块设备：主要功能是通过Go语言代码列出系统中的块设备。例如，[main.go](example/main.go) 文件展示了一个示例程序，它调用 `go-lsblk`
  库来获取和打印块设备信息。
