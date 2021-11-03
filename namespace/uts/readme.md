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

执行程序后：

```
程序进入到了 sh 终端。接收标准输入输出
px@ubuntu:~$ sudo workspace/minidocker/namespace/uts/main 
# 

```
查看进程的信息
```sh
main 为主进程，其中还衍生出 4 个子进程和 一个 sh 进程
4 个子进程[查看源码可知 可能为 标准输入  标准输出  标准错误  进程wait]和 main 进程 在一个命名空间中
proot@ubuntu:/home/px# ps -ef | grep uts
root       67985   59574  0 17:19 pts/3    00:00:00 sudo workspace/minidocker/namespace/uts/main
root       67986   67985  0 17:19 pts/3    00:00:00 workspace/minidocker/namespace/uts/main
root       68424   68391  0 17:29 pts/0    00:00:00 grep --color=auto uts
root@ubuntu:/home/px# pstree -lp | grep main
           |               |-gnome-terminal-(37736)-+-bash(59574)---sudo(67985)---main(67986)-+-sh(67991)
           |               |                        |                                         |-{main}(67987)
           |               |                        |                                         |-{main}(67988)
           |               |                        |                                         |-{main}(67989)
           |               |                        |                                         `-{main}(67990)
root@ubuntu:/home/px# readlink /proc/67986/ns/uts 
uts:[4026531838]
root@ubuntu:/home/px# readlink /proc/67991/ns/uts 
uts:[4026532627]
root@ubuntu:/home/px# readlink /proc/67987/ns/uts 
uts:[4026531838]
root@ubuntu:/home/px# readlink /proc/67990/ns/uts 
uts:[4026531838]

使用 lsns 查看命名空间
root@ubuntu:/home/px# lsns | grep 4026531838
4026531838 uts       344     1 root             /lib/systemd/systemd --system --deserialize 18
root@ubuntu:/home/px# lsns | grep 4026532627
4026532627 uts         1 67991 root             sh
```
修改hostname 不影响主机
```
<!-- 宿主机开启一个进程，并进行uts 隔离，然后修改uts -->
px@ubuntu:~/Desktop$ sudo ~/workspace/minidocker/namespace/uts/main 
[sudo] px 的密码： 
# hostname
ubuntu
# hostname -b testhostname    
# hostname
testhostname
# 

<!-- 另外开启一个shell 终端查看 主机的hostname 并没有受到影响 -->
px@ubuntu:~/Desktop$ hostname
ubuntu
testhostname


```