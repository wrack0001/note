# Linux出现rootfs根目录磁盘站满问题

## 1.清理磁盘中的日志文件

​	查看磁盘的使用情况：

```shell
[root@feng020 ~]# df -l
Filesystem     1K-blocks     Used Available Use% Mounted on
/dev/xvda1      20641404 10565932   9026948  54% /
tmpfs            4029028        0   4029028   0% /dev/shm
/dev/xvdb1     103210940 67011820  30956312  97% /hotdata

```

然后使用找到磁盘中那个文件夹占用空间大。

```shell
du -h --max-depth=1
```

也可以使用查看占用空间大的文件

```shell
ls -lh
```

如果知道哪些日志文件占用空间大的话，可以忽略上面，直接找到目录删除日志。

## 2.如果删除了日志文件发现rootfs空间还是100%

1. 首先获得一个已经被删除但是仍然被应用程序占用的文件列表，如下所示：

```shell
[root@test]# lsof|grep deleted
proftpd  3468   nobody  4r   REG        8,2    1648        667033 /etc/passwd (deleted)
proftpd  3468   nobody  5r   REG        8,2    615        667032 /etc/group (deleted)
syslogd  3854    root  2w   REG        8,2  65521380        164531 /var/log/messages.1 (deleted)
syslogd  3854    root  3w   REG        8,2  22728648        164288 /var/log/secure.1 (deleted)
syslogd  3854    root  5w   REG        8,2  4247977        164316 /var/log/cron.1 (deleted)
```



2. kill 掉上面列出的进程

如果知道是哪些进程，可以直接重启进程。

【总结】出现此问题的原因：系统打开了很多大数据文件，占用了数据库缓存。

解决办法：删除这些大数据文件，然后直接重启打开日志程序。

我们上次出现这个问题的原因是，程序中的日志文件过大。大的日志文件加载到缓存中，导致rootfs占满。