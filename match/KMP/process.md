##  process
在朴素的模式匹配基础上优化

* 假如
```shell
S:	  abcdefg
	  |||-
T:    abcx 

S:abcdefg
	  -
T:    abcx
```
S的a开头的一轮匹配在x处失败，这个时候不需要用 S 的b和c与 T 匹配: b和c 都不和a匹配
这时跳到用 S 的d和 T 匹配
至于能跳到什么位置取决于 T 的元素的之前元素和 T 自己能匹配到多长

* 假如
```shell
S:	  abcabcabc
	  |||||-
T:	  abcabx 

S: abcabcabc
	  |||||-
T:    abcabx
```
i 一直处于递增 : 判断了 n-1 次成功 第 n 次失败的话，下次只需要从 S 的第 n 次重新开始判断 (好马不吃回头草)
j 变小多少取决于 T 的当前元素的之前字符是否和 T 自己有匹配

