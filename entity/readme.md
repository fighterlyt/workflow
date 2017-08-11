
# 实体描述

## 流程定义
流程包括以下要素

*   名称
*   创建时间
*   步骤定义
*   版本

### 步骤定义

*   步骤名称
*   步骤大类
*   步骤小类
*   下一步以及转换


## 定义语言
使用**graphviz的dot语言**，并且提供了如下扩展


### 属性扩展
由于原始的**dot语言**不满足语义，所以添加一些扩展的属性

#### 图扩展
扩展了如下属性

|   属性名 |   类型  |   合法值     |默认值  | 含义  |
|   ---     |   --- |   --- | --- | --- |
|   version |   int |   正整数 |   1   |   流程版本    |
|   createTime  |   time.Time   |   任意合法时间  |   time.Now()  |   流程定义创建时间  |
|   canWithDraw |   bool    |   true/false    |  true   | 可否从本步骤撤回  |
|   multiSend   |   bool    |   true/false  |   true    | 能否发送多人  |
|   singleForward    |   bool    |   true/false  |   true    |  如果当前处理人为多人，那么单个人处理完是否为处理完成  |

#### 节点扩展

|   属性名 |   类型  |   合法值     |默认值  | 含义  |
|   ---     |   --- |   --- | --- | --- |
|   kind   |   字符串    |   start/finish/processing   | processing  | 是否为开始/结束/处理中节点 |  
|   canWithDraw |   bool    |   true/false    |  true   | 可否从本步骤撤回  |
|   multiSend   |   bool    |   true/false  |   true    | 能否发送多人  |
|   singleForward    |   bool    |   true/false  |   true    |  如果当前处理人为多人，那么单个人处理完是否为处理完成  |

## 流程实例
### 节点
节点表示流程的一个状态，是一个静态信息
### 边
边对应了流程状态的转移，是一个动态信息，包括以下

*   来源
*   去向
*   时间
*   方向
*   附加信息
*   边名称
