# 二叉树集合

## 定义
### 二叉树几个概念
* 满二叉树：从高到低，除了叶节点外，所有节点左右节点都存在
* 完全二叉树：比满二叉树少几个叶节点，从左向右放子节点。
* 平衡二叉树：空树或者它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树也都是平衡树
* 二叉搜索树：空树或者二叉树的所有节点比他的左子节点大，比他的右子节点小。
* 红黑树：不仅是具有二叉搜索树的属性，还具有平衡树的属性，有序且子树差不超过1，颜色规则：根节点和特殊节点（即叶节点下面两个虚无的节点和未填写的节点）是黑的，红节点的左右子节点是黑的，最重要的是对于每个节点，从该节点到子孙叶节点的所有路径包含相同数目的黑节点。

## 原理
* 左子树和右子树的高度之差绝对值不超过1
* 而且规定，平衡二叉树的每个节点的平衡因子只能是-1 ，1 ，0；

### 普通二叉树
* 普通二叉树在极端情况下，有可能变成单链表。如果是单链表，对于二叉树来说，查找二叉树变成线性的，那就完全没有意义了

### 平衡法
* 要是不平衡的子树维持平衡，需要作出对应的操作，大概分以下几类

#### 右旋
#### 左旋
#### 右左旋
#### 左右旋


## todo list
- [x] 二叉树
- [x] 二叉搜索树
- [ ] 平衡二叉树
- [ ] 红黑树
- [ ] 伸展树
- [ ] AA-树
- [ ] 二叉堆


## 参考资料
* 主要原理和流程
> https://baijiahao.baidu.com/s?id=1577200621749785094&wfr=spider&for=pc
https://blog.csdn.net/followmyinclinations/article/details/50426413

* 平衡因子的计算 https://blog.csdn.net/travelerwz/article/details/52186357 （有点小错误）
* 二叉树的一些概念  https://blog.csdn.net/qq_37887537/article/details/75647670
* 要学习哪些树 [点此前往](https://mp.weixin.qq.com/s?__biz=MzUxMTk0MDI0Mw==&mid=2247483703&idx=1&sn=12b9a7c13036216c25a2c909abad5e19&chksm=f96d42cbce1acbdd0854712086435b910349bec7076da0115b4f720cf91fbc344048d0087c02&mpshare=1&scene=1&srcid=0911JDSXZQDVZG2OFE3Qdq1Q&key=dcbe06cbc6dde9f86ae3bdb0ba4389e0271124d09d13d5e187a935f04a293f3432d72ce5eda0b420e84d732900ac082440746a8d9545ff9d2b88b318fcee868fa2f264cde1a7dce9ef7c5a606605125a&ascene=0&uin=NDcxMTY4Mzc1&devicetype=iMac+MacBookPro12%2C1+OSX+OSX+10.11.5+build(15F34)&version=12020810&nettype=WIFI&lang=zh_CN&fontScale=100&pass_ticket=OTvthPCnI%2BpI58GVRtnjn4%2B5CJ3s4szVZdU16magQoHu3tCWdlzw9jmxnx0JUDOs)