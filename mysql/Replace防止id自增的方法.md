在使用 replace 的时候总会遇到 id 自动增长的苦恼，如果是一个频繁的任务，而且数据量很大，那么就会出现id不够用，导致数据库不能写入数据。

自然有人说 最好不用 replace ，和replace 的不好；
但是就是要使用 replace 而写还想id 自增应该怎样做。

	1. 首先床架一个表
	```
	-- auto-generated definition
	create table t1
	(
	id int unsigned auto_increment comment 'ID，自增'
	primary key,
	uid bigint unsigned default 0 not null comment '用户uid',
	name varchar(20) default '' not null comment '用户昵称',
	constraint u_idx_uid
	unique (uid)
	)
	comment '测试replace into' engine = InnoDB;
	
	```
	
	
	2. 插入一条数据

	replace into t1 (id, uid, name) value (
		(select id from t1 b where b.uid = 100)
		,100,"test100");
	
	执行这条语句两次，id都是 1
		
	
	这里说明，id 并不会出现自增。那么 auto increment 有没有自增呢？
	3. 这里我们执行下面这条语句
	
	insert into t1(id, uid, name) value (null,103,"test103");
	
	
	
	结果插入的id = 2 说明 auto increment 并没有自增
	
	这就解决了replace 自增的问题。
	
	其实，有些时候我们需要对其中一张表中的数据处理后，导入
	到另外一张表中，而且 数据可能会重复。
	当然是可以通过 查询出来然后通过代码判断是否存在然后决定是
	Update 还是 跳过这条重复的数据。
	
	但是我们想的是直接通过 一条sql 直接将数据 从A 表导入B表，
	而且有重复的就直接替换掉，不想再通过 其他语言的代码进行
	处理判断后再决定是inssert 还是 update或者跳过。
	
	这样的前提下这个方法就需要replace这个方法了。这个方法的弊端
	大家也都清楚，重复的会增加 id 的自增。所以我才写了这篇文章。
	
	
	
