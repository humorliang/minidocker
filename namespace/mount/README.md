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

Mount Namespace 是用来隔离各个进程看到的挂载点视图，在不同的namespace中看到的文件层次是不同的。
在当前的Namespace 进行mount 和 unmount 操作，只会影响当前的namespace ,不会影响全局的 Namespace,
新创建的Namespace 会继承，父进程的Namespace,但是两者并没有关联，（除共享子树）


系统调用参数  NEWNS 也是第一个Namespace 。

首先， /proc 是一个虚拟的文件系统，是存储在内存当中的文件系统，proc 文件系统提供了，而外的机制，可以通过内核
或者内核模块将信息发送给相关进程。 可以通过命令：
```sh
#  找到相关挂载点信息
mount | grep proc
# 
# proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
# systemd-1 on /proc/sys/fs/binfmt_misc type autofs (rw,relatime,fd=28,pgrp=1,timeout=0,minproto=5,maxproto=5,direct,pipe_ino=27360)
```

### 子命名空间会影响父命名空间
```
https://bbs.huaweicloud.com/blogs/detail/223530
https://unix.stackexchange.com/questions/281844/why-does-child-with-mount-namespace-affect-parent-mounts
说明： 需要有一套独立的 rootfs 文件系统，否则你挂载的文件夹跟父 namespace 的其他文件视图是同一个。导致出错
```

