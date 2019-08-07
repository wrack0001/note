# 简单实现hashSet -- 应用篇.md

> 写这篇文章的起因：
>> - 在使用golang和php的时候我们经常会判断一个变量是否在map或者数组中；
>> - 在php中的时候我们经常使用in_array()来判断，这里是我们经常用的。
>> 其实用in_array()这个函数就是对数组进行遍历判断是否存在时间复杂度是log(n)
>> 一般数组不是太大也就这样了；
>> - 但是在golang中并没有封装这个函数，这个需要我们自己去封装。
>> 如果也是用循环的方式去判断变量是否存在，那是不是感觉很low，也就是这个原因引出了这篇文章

- 此篇文章GitHub地址：https://github.com/wrack0001/note/blob/master/skill/%E7%AE%80%E5%8D%95%E5%AE%9E%E7%8E%B0hashSet%20--%20%E5%BA%94%E7%94%A8%E7%AF%87.md

1. 在php和golang中都适用，主要原理是判断key是否存在

```cassandraql
    // 在golang中可以这样

    type HashSet struct {
        set map[interface{}｝]bool
    }
    
    func NewHashSet() *HashSet {
        return &HashSet{make(map[interface{}]bool)}
    }
    
    func (set *HashSet) Add(k interface{}) bool {
        _, ok := set.set[k]
        set.set[i] = true
        return !ok //False if it existed already
    }
    
    func (set *HashSet) Get(k interface{}) bool {
        _, found := set.set[k]
        return found //true if it existed already
    }
    
    func (set *HashSet) Remove(k interface{}) {
        delete(set.set, k)
    }


```

```cassandraql
    // InArray 函数就简单的模拟了in_array 函数
    // 需要注意的是，在设置 $HashSet 的时候需要设置在key 中，
    // 例 ： 
    
    $HashSet = [
        "a" => true,
        "b" => true,
    ];

    function InArray(&$HashSet,$key) {
        return isset($HashSet[$key]);
    }
   
    var_dump(InArray($HashSet,"b"))

    // 当然这里也可以不用InArray 封装，直接用 isset,这里只是为了举例子

    // 如果是正常数组
    
    $HashSet2 = [
        "a" ,"b"
    ];

    // 可以使用 array_flip() 函数（交换数组中的键和值。）
    // 再使用isset进行判断


```

**总结** 
1. 就是数的key与value调换，利用map和数组中的key的唯一性来判断 
2. 判断key是否存在的时间复杂度是 O(1);



###### *参考文献*
[1] [Go 语言简单实现HashSet](https://blog.cyeam.com/golang/2014/07/15/go_hashset)
