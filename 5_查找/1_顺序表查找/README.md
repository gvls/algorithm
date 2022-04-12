##  顺序查找
实现简单
使用时间长 : `O(n)` 

###   步骤
从头到尾遍历

```shell
for i = 0 ; i <= n ; i++ {
	if a[i] == key {
		return i
	}
}
return -1
```

###   优化

```shell
a[0] = key // 哨兵
i = n
for a[i] != key { // 减少了 判断是否数组越界
	i--
}
return i
```
