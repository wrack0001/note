#php操作kafka之win下安装php7的kafka扩展

1. 在这个[地址](http://pecl.php.net/package/rdkafka)下载指定版本的扩展文件
2. 我这里下载的是 3.0.5 的版本（这里基本上支持的都是 7.1和7.0版本的php）

![1](https://upload-images.jianshu.io/upload_images/15839628-be16051aa1fe18c1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1000/format/webp)

3. 点击进去后选择你对应的php版本（我下载的是 7.1 Non Thread Safe (NTS) x64 ）

![2](https://upload-images.jianshu.io/upload_images/15839628-9e0c8a4583ff7305.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1000/format/webp)

4. 下载下来后的文件有这些文件

![3](https://upload-images.jianshu.io/upload_images/15839628-2ad9b3bbfcf8f7c9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/615/format/webp)

5. 在php.ini中增加 extension=php_rdkafka.dll



---
[GitHub地址](https://github.com/wrack0001/note/blob/master/php/php%E6%93%8D%E4%BD%9Ckafka%E4%B9%8Bwin%E4%B8%8B%E5%AE%89%E8%A3%85php7%E7%9A%84kafka%E6%89%A9%E5%B1%95.md)

