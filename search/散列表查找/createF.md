##  createF
创造一个简单、均匀、存储利用率高的散列函数
it need calculate easy : if complex, do search will use more time
散列地址分布均匀 : improve 利用率 of space. reduce 冲突

* 直接定址
use 线型 函数
need 实现知道关键字分布情况
suit 查找表小且连续
`f(key) = key` 
`f(key) = a x key + b ` 

* 数字分析法
使用关键字的一部分来计算存储位置
suit 关键字位数比较大

* 平方取中
对关键字平方后，取中间位置
suit 不知到关键字分布，位数不是很大

* 折叠法
把关键字分割位数相等的几个部分后，相加求和，按照散列表表长取后几位作为散列地址
suit 关键字位数多

* 除留余数法
若散列表长m : `f(key) = key mod p (p<=m)` 
p 最好是接近m的最小质数/不包含小于20质因子的合数

* 随机数
suit 关键字不定长

