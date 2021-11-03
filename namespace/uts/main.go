// Copyright 2021 px
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 指定fork() 出来的进程的初始命令， 使用 sh 来执行
	cmd := exec.Command("sh")
	// 对这个进程的属性进行设置 ， 创建一个 UTS Namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run 函数 封装 好进程的创建设置等
	// start() 函数进行进程创建， 并执行 sh  命令
	// wait() 等待相关进程输入输出
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
