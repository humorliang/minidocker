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


IPC 命名空间将进程与 SysV 风格的进程间通信隔离开来。 
这可以防止不同 IPC 命名空间中的进程使用例如 SHM 函数系列在两个进程之间建立一定范围的共享内存。
相反，每个进程将能够对共享内存区域使用相同的标识符并生成两个这样的不同区域。

```
<!-- 创建ipc 进程，并且分配一个ipc 命名空间 -->
px@ubuntu:~/Desktop$ sudo ~/workspace/minidocker/namespace/ipc/main 
# ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息      

# ipcmk -Q
消息队列 id：0
# ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息      
0x0170a15a 0          root       644        0            0           

# 

<!-- 另外开启一个 shell 终端查看 ipc 信息 -->
px@ubuntu:~/Desktop$ sudo ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息

<!-- 宿主机看不到 ipc 进程创建的 message queues ,说明 ipc 的namepsace 隔离成功了-->
```