## 常见的二进制操作
### 位运算
符号|描述|运算规则
---|---|---
&|与|两个位都为1时，结果才为1
||或|两个位都为0时，结果才为0
^|异或|两个位相同为0，相异为1
~|取反|0变1，1变0
<<|左移|各二进位全部左移若干位，高位丢弃，低位补0
>>|右移|各二进位全部右移若干位，对无符号数，高位补0，有符号数，各编译器处理方法不一样，有的补符号位（算术右移），有的补0（逻辑右移）
### 基本操作
* a=a^0=0^a
* 0=a^a
* a=a^b^b
### 交换两个数
* a=a^b
* b=a^b
* a=a^b
### 移除最后一个 1
* a=n&(n-1)
### 获取最后一个 1
* diff=(n&(n-1))^n