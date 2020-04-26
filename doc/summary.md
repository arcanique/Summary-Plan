# <font color=blue>golang language</font>
## <font color=#881181>go 垃圾回收机制</font>
> golang 垃圾回收主要是三色标记法、写屏障、辅助垃圾回收机制
### 三色标记法
  https://juejin.im/post/5d78b3276fb9a06b1829e691

 https://juejin.im/post/5d56b47a5188250541792ede

## <font color=#881181>go 协程调度</font>
https：//jingwei.link/2019/05/06/golang-routine-scheduler.html
  https://juejin.im/entry/5b3f2f166fb9a04fb900b119
  https://qfgolang.com/?p=2214
  https://zhuanlan.zhizhu.com/p/2698463
## <font color=#881181>go 协程通信</font>
 https://juejin.im/post/5c6a507fe51d45086925e22c

## <font color=#881181>go 反射机制/map</font>
 http://juejin.im/post/5be452d96fb9a049ef261452
 draveness.me/golang/docks/part2-foundation/ch03-datastructure/golang-hashmap/
##<font color=#881181> go 接口机制</font>
 c.biancheng.net/golang/
## <font color=#881181>go cgo原理</font>
## <font color=#881181>go 基本工具</font>
  ### go vet 工具（go tool vet）
  + go vet有价值：它可以在编译阶段和运行阶段发现bug， 能够做静态检查，避免一些bug
  + go vet 可以检查的部分：
  > + Print-format 错误( go 是强类型，但是print format 错误不会在编译时不会报错出来)
  > + vet 可以检查出一些bool的错误，比如一直为false 或 true ，以及一些冗余表达式
  > + 注意，这个代码包含竞态，可能不输出任何东西。事实上，main函数可能在所有的协程执行前已经结束，这导致进程退出。当读取变量的时候，在range块内的go协程可能是有问题的
  > + 检测一些不可达代码
  > + vet 可能会误报或漏报
  ### go 内存逃逸分析
  + go build -gcflags=-m 使用-gcflags可以进行内存逃逸分析

  ### go test UT测试
  ### go tool pprof
  ### go tool trace
  ### go complie -S xx.go (反汇编)

## <font color=#881181>go 内存调优（内存逃逸）
  studygolang.com/articles/19386
  studygolang.com/articles/12008
## <font color=#881181>go 异常</font>

### Go中用 defer panic recover来处理
  + go 语言通过panic函数抛出一个异常，在defer语句中通过recover来捕获这个异常（recover只会在derfer中才会捕获异常）
  + go 语言中也支持自定义异常，通过errors.New来产生一个自定义的异常，panic语句可以接收这个错误作为异常返回
  + 将异常捕获函数放在匿名函数不会影响主函数运行
 
# <font color=blue>net</font>
## ip/tcp
### TCP的拥塞算法
+ 基本概念
  > RTT(Round Trip Time)：一个连接的往返时间，即数据发送时刻到接收到确认的时刻的差值；
  > RTO(Retransmission Time Out)：重传超时时间，即从数据发送时刻算起，超过这个时间便执行重传
  > 流量控制是通过大小可变的滑动窗口实现的。
+ Cubic
  >它们是事件驱动的！无论是丢包，还是RTT增加，都被视为一种严重的事件
+ BBR
  > bbr周期性的试图探测是否有更多的带宽，如果有，那么就利用它，如果没有，就退回到之前的状态。

### 隧道技术
> GRE 需要内核模块 ip_gre.ko ，GRE是最初由CISCO开发出来的隧道协议，能够做一些IP-in-IP隧道做不到的事情。比如，你可以使用GRE隧道传输多播数据包和IPv6数据包。
> ipip 需要内核模块 ipip.ko ，该方式最为简单！但是你不能通过IP-in-IP隧道转发广播或者IPv6数据包。

### overlay（vxLan）技术
https://zhuanlan.zhihu.com/p/84693975

### underlay 技术（ipvlan）

## epoll 原理（select/poll）
 > 复用的含义：那么这些文件描述符 FD 要复用什么呢？在网络场景中复用的就是任务处理线程，所以简单理解就是多个IO共用1个处理线程
  ### select/poll 技术
  >+ select 使用一个宏定义函数按照 bitmap原理填充 fd，默认大小是 1024个
  ### epoll 技术



  ### select/poll 与 epoll的 技术总结
  segmentfault.com/a/1190000003063859
  jianshu.com/p/dfd940e7fca2
  leokongwq.github.io/2016/10/20/linux-io-epoll.html
  yq.aliyun.com/articles/367908
## iptables/ipvs
![iptables_ipvs](/images/iptables_ipvs.PNG)

### ipvs 三种模式
+ NAT 模式
![ipvs_nat](/images/ipvs_nat.PNG)
> + client 能够和vs 三层互通， rs和vs能够三层互通
> + nat 模式vs需要处理请求报文和响应报文

+ DR 模式
![ipvs_dr](/images/ipvs_dr.PNG)

> + dr 模式 vs 和 rs 机器必须有一个对应网卡在同一二层内（负责无法实现）
> + dr 模式 client 和 rs 必须在同一个二层内
> + dr 模式 rs上必须配置一个虚ip 
> + dr 模式 rs 只处理请求报文
+ TUN 模式
![ipvs_tun](/images/ipvs_tun.PNG)

> + tun模式 即隧道模式 client 和 realserver 三层互通
> + tun模式 rs上的tun设备上必须配置vip，便于与client通信
> + tun模式 rs 只处理请求报文

+ ipvs的arp问题
+ > 因此，必须要保证路由器只保存Director上VIP对应的MAC地址，即只允许Director才对路由器的ARP广播进行回应。也就是说，所有RS上的VIP必须隐藏起来。
一般通过将Real Server上的VIP设置在lo接口的别名接口上(如lo:0)，并设置arp_ignore=1和arp_announce=2的方式来隐藏RS上的VIP。至于VIP为何要设置在lo接口上以及为何要这样设置这两个arp参数，请参看应用负载均衡之LVS(二)：VS_TUN和VS_DR的arp问题，内容非常详细。

 + > 操作如下： 
    ```
    echo 1 >/proc/sys/net/ipv4/conf/all/arp_ignore
    echo 2 >/proc/sys/net/ipv4/conf/all/arp_announce
    或者

    sysctl -w net.ipv4.conf.all.arp_ignore=1
    sysctl -w net.ipv4.conf.all.arp_announce=2
    或者将arp参数设置到内核参数配置文件中以让其永久生效。

    echo "net.ipv4.conf.all.arp_ignore=1" >>/etc/sysctl.conf
    echo "net.ipv4.conf.all.arp_announce=2" >>/etc/sysctl.conf
    sysctl -p
    在网上几乎所有文章还设置了lo接口上的arp参数：

    echo 1 >/proc/sys/net/ipv4/conf/lo/arp_ignore
    echo 2 >/proc/sys/net/ipv4/conf/lo/arp_announce
    但这没有任何意义，因为从lo接口不受arp参数的影响。
    ```

应该在配置VIP之前就设置arp参数，以防止配置VIP后、设置arp抑制之前被外界主机发现。
+ ipvs的调度算法
  +  >静态调度：
  轮叫调度（Round-Robin Scheduling，rr）
  加权轮叫调度（Weighted Round-Robin Scheduling，wrr），按照权重比例作为轮询标准
  目标地址散列调度（Destination Hashing Scheduling，dh），目标地址哈希，对于同一目标IP的请求总是发往同一服务器
  源地址散列调度（Source Hashing Scheduling，sh），源地址哈希，在一定时间内，只要是来自同一个客户端的请求，就发送至同一个realserver
  + > 动态反馈调度：
  最小连接调度（Least-Connection Scheduling，lc），调度器需要记录各个服务器已建立连接的数目，当一个请求被调度到某服务器，其连接数加1；当连接中止或超时，其连接数减1。当各个服务器的处理能力不同时，该算法不理想。
  加权最小连接调度（Weighted Least-Connection Scheduling，wlc）
  基于本地的最少连接（Locality-Based Least Connections Scheduling，lblc），目前该算法主要用于cache集群系统。
  带复制的基于局部性最少连接（Locality-Based Least Connections with Replication Scheduling，lblcr），目前主要用于Cache集群系统。


### iptables 使用
+ 五个netfilter 链
  LOCAL_IN
  PRE_ROUTING
  FORWARD
  LOCAL_OUT
  POST_ROUTING
+ 5 张表
  nat 表

  filter 表

  manale 表

  ### <font color=red>vpp<font>

+ 基本操作

# <font color=blue>k8s&docker</font>
## <font color=red>docker 原理（ns/cg，基本命令等）</font>
  https://draveness.me/docker/ 

  ### docker三大技术之 namesapce
  + 实现原理
  + 主要

  ### docker 三大技术之 cgroup
  + 实现原理

  ### docker 三大技术之 文件系统
  + 实现原理

  ### docker 镜像&镜像制作

## k8s 基本原理
  kubernetes.io/zh/docs/home/
  infoq.cn/article/KNMAVdo3jXs3qPKqTZBw

## k8s 对象（pod / deployment/rc/rs/service/endpoint/networkpolicy）
  feisky.gitbooks.io/kubernetes/concepts/network-policy.html

  ### schedule 选主过程
  
  ### client-go 
  ### pod 
  + 基本原理
  > + Pod 是一组紧密关联的容器集合，它们共享 IPC、Network 和 UTS namespace，是 Kubernetes 调度的基本单位
  > + pod 分为static pod 和普通pod
  > 静态pod直接由某个节点上的kubelet程序进行管理，不需要api server介入，静态pod也不需要关联任何RC，完全是由kubelet程序来监控，当kubelet发现静态pod停止掉的时候，重新启动静态pod。
  + 创建、删除 过程
 ![pod_create](/images/pod_create_process.png)

  ### rc/rs 概念、区别与联系
  + Replication Controller
  > + 副本控制器，维持pod 副本数量（小于则增加，大于则减少）
  > + 根据副本健康状况，清理掉不正常的副本创建新的副本
  > + 弹性伸缩，可以动态的调整副本数（根据实际的业务需求）
  > + 滚动升级，逐步怎加新版本的副本数，逐步减少老版本的副本数，通过kubectl rolling-update 来进行
  > + RC只支持基于等式的selector（env=dev或environment!=qa），但RS还支持新的，基于集合的selector（version in (v1.0, v2.0)或env notin (dev, qa)），这对复杂的运维管理很方便。

  + RS Replicats Set(副本集合)
  支持基于集合的selector， selector in xxx
  ### deployment
  > + deployment 继承了RC的维持pod副本数量的能力；
  > + 与此同时deployment 支持记录pod升级过程记录（每一次升级的过程），支持回滚的能力
  > + 能够暂停和启动升级（升级的过程中能够暂停和启动）
  > + 多种升级方案：Recreate：删除所有已存在的pod,重新创建新的; RollingUpdate：滚动升级，逐步替换的策略，同时滚动升级时，支持更多的附加参数，例如设置最大不可用pod数量，最小升级间隔时间等等
  > + 常用命令kubectl rollout
    kubectl rollout undo 回滚
    kubectl rollout status 查看状态
    kubectl rollout pause 暂停
    kubectl rollout resume 恢复

  ### statfulset
  > + 稳定的持久化存储，即 Pod 重新调度后还是能访问到相同的持久化数据，基于 PVC 来实现
  > + 稳定的网络标志，即 Pod 重新调度后其 PodName 和 HostName 不变，基于 Headless Service（即没有 Cluster IP 的 Service）来实现
  > + 有序部署，有序扩展，即 Pod 是有顺序的，在部署或者扩展的时候要依据定义的顺序依次依序进行（即从 0 到 N-1，在下一个 Pod 运行之前所有之前的 Pod 必须都是 Running 和 Ready 状态），基于 init containers 来实现,有序收缩，有序删除（即从 N-1 到 0）

  ### service/endpoint/networkpolicy

# <font color=blue>数据结构与算法</font>
## 二叉树
## <font color=red>map</font>
## 二分法
## 排序算法
## 限流算法
## 红黑树