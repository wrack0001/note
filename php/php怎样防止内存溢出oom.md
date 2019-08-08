#php怎样防止内存溢出oom

> 在编程的时候经常会遇到下载大文件，如果将下载的大文件都读到内存中再存储到硬盘
>这个时候有很大程度会出现内存不够用的情况。这时候就需要我们下载一部分然后存入到硬盘
>然后再下载一部分存入到硬盘。


- 实现方法，直接上代码

```
<?php

    /**
     * @param string $url 要下载的文件的地址
     * @param string $fileName 文件保存位置
     * @param int $step 每次抓取的步长
     * @param int $timeOut 下载超时
     * @return array
     */
    function down_load($url, $fileName, $step = 1024, $timeOut = 3600)
    {
        try {
            $ch = curl_init($url);
            // 关闭安全认证
            curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false); 
            curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false); 

            curl_setopt($ch, CURLOPT_RETURNTRANSFER, true); // 获取数据返回
            curl_setopt($ch, CURLOPT_BINARYTRANSFER, true);
            curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
            curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 60);
            curl_setopt($ch, CURLOPT_TIMEOUT, $timeOut);        // 允许 CURL 函数执行的最长秒数。
            curl_setopt($ch, CURLOPT_BUFFERSIZE, $step);
            // 分段获取远程数据，并存储到本地
            $size = 0;
            curl_setopt($ch, CURLOPT_WRITEFUNCTION, function ($ch, $str) use ($fileName, &$size) {
                $len = strlen($str);
                file_put_contents($fileName, $str, FILE_APPEND);
                $size += $len;
                return $len;
            });

            curl_exec($ch);
        } catch (Exception $e) {
            return [false, $e->getMessage()];
        }

        //返回下载文件的大小
        Return [$size, ""];
    }

```

**这个是自己封装的函数，在发表之前有略微的改动，如果自己不喜欢可以修改。**

[GitHub地址](https://github.com/wrack0001/note/blob/master/php/php%E6%80%8E%E6%A0%B7%E9%98%B2%E6%AD%A2%E5%86%85%E5%AD%98%E6%BA%A2%E5%87%BAoom.md)

---
***参考文献***

[1] [php下载大文件的方法](https://blog.csdn.net/qq_36663951/article/details/82145359) 