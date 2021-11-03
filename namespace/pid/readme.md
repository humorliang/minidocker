<!--
 Copyright 2021 px
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

PID 命名空间为进程提供了一组独立于其他命名空间的进程 ID (PID)。 
PID 命名空间是嵌套的，这意味着当创建一个新进程时，它将为每个命名空间拥有一个 PID，从其当前命名空间到初始 PID 命名空间。
因此，初始 PID 命名空间能够看到所有进程，尽管其 PID 与其他命名空间将看到的进程不同。

在 PID 命名空间中创建的第一个进程被分配了进程 ID 号 1，并接受与普通 init 进程相同的大部分特殊处理，最值得注意的是命名空间内的孤立进程附加到它。
这也意味着该 PID 1 进程的终止将立即终止其 PID 命名空间中的所有进程及其任何后代。


```
<!-- 查看到main 进程的 id 为 20131 -->
# pstree -lp 
├─gnome-terminal-(19983)─┬─bash(19991)───sudo(20130)───main(20131)─┬─sh(20136)

<!-- 进入当前pid 的命名空间 打印当前进程的 id 为1  -->
px@ubuntu:~/Desktop$ sudo ~/workspace/minidocker/namespace/pid/main
# echo $$
1
# 

<!-- 在该进程内不能使用 ps  和 pstree 命令，因为 这个命令还是 查找的 /proc 文件夹下面的，需要 设置 mount namespace 才可以 -->
```