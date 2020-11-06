# 在for循环中goroutine调用循环变量问题

> 今天有一个小伙伴问我一个for循环中调用了goroutine，有些时候打印的是连续的数字，
>有些时候打印的是最终的一个数字

- 上代码

```goregexp
	for i:=1;i<=3;i++{
		go func() {
			fmt.Println(i)
		}()
	}
```
> 输出的是 4,4,4

```goregexp
	for i:=1;i<=3;i++{
		temp := i
		go func() {
			fmt.Println(temp)
		}()
	}
```
> 输出的是 3,2,1 【注：不保证顺序】

- 我一般是这样写的
```goregexp
	for i:=1;i<=4;i++{
		go func(temp int) {
			fmt.Println(temp)
		}(i)
	}
```
> 输出的是 3,2,1 【注：不保证顺序】

- 为什么第一种方式打印的就都是 444 呢？
> 这里就需要考虑到变量的指针问题了
>第一种方式：当把变量**i**传如到goroutine中，这时候goroutine并没有执行呢。
>注意一点这里的**i**在for中算是全局变量。全局变量你在修改以后，这是你给goroutine中的
>变量**i**是不是也发生了改变，所以打印出来的都是4

>第二种方式与第三种其实一样的，都是将变量**i**赋值给临时变量**temp**了。
>第二种方式不好理解的话，你看第三种方式。你将变量**i**传给了func，是不是以后temp的修改
>与变量**i**都没有任何关系了。所以打印出来的就是【3，2，1】


[GitHub地址](https://github.com/wrack0001/note/blob/master/golnag/%E5%9F%BA%E7%A1%80%E7%AF%87/goroutine/%E5%9C%A8for%E5%BE%AA%E7%8E%AF%E4%B8%ADgoroutine%E8%B0%83%E7%94%A8%E5%BE%AA%E7%8E%AF%E5%8F%98%E9%87%8F%E9%97%AE%E9%A2%98.md)
