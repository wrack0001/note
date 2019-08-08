# php导出百万数据到csv

> 起因：有一次业务需求，需要导出数据库中的所有数据给接口下游，以便下游比对所有数据是否一致。

- 这时候数据库中已经有将近100+万数据。下游希望要的数据并不是数据库的格式，格式特定样式的。这时候就需要关联上好多个表进行同时查询。
- 我最开始我是使用orm模式进行处理导对应字段，导出的数据（我做了redis缓存），整个导出过程用了半天。
- 我用phpAdmin联合查询导出数据的时候用了不到3分钟；

**思考**
1. jion与orm用哪个
2. 用orm为什么慢
3. 用orm是否也可以很快

**先说今天的主题**
- 假设现在有一个天津市学生表（有100+万数据）

student table

| 字段        | 类型    |备注     |
| :---       |:-------|:----    |
| id        | int    |     |
| student_no        | string    |     |
| name        | string    |     |
| areaId        | int    |所属区域的id |
| gradeId        | int    |年级ID  |
| SchoolId        | int    |学校ID  |

area table  （15条数据）

| 字段        | 类型    |备注     |
| :---       |:-------|:----    |
| id        | int    |     |
| name        | string    |   区域名称  |

greate table  （12条数据）

| 字段        | 类型    |备注     |
| :---       |:-------|:----    |
| id        | int    |     |
| name        | string    |   区域名称  |

School table  （100+条数据）

| 字段        | 类型    |备注     |
| :---       |:-------|:----    |
| id        | int    |     |
| name        | string    |   校园名称  |


要求导出所有内容到csv文件中
格式要求：学生编号,学生姓名，所属区域，所属年级，所属学校

**下面开始写代码**

```sql
    
    set_time_limit(0);      // 设置超时
    ini_set('memory_limit', '100M');       // 设置最大使用的内存
    
    header("Content-type:text/csv");
    header("Content-Disposition:attachment;filename=" . date('Ymd'). '.csv');
    header('Cache-Control:must-revalidate,post-check=0,pre-check=0');
    header('Expires:0');
    header('Pragma:public');
    $out = fopen('php://output', 'w');

    $bom = chr(0xEF).chr(0xBB).chr(0xBF);     // 防止乱码


    $func = function ($list){
        $arr = [];
        foreach ($list as $v){
            $arr[$v['id']] = $v['name']
        }
        return $arr;
    }

     // todo 注意我这里都是用sql语句直接代替查询
    $list = select * from area;
    $area = $func($list);
    $list = select * from greate;
    $greate = $func($list);
    $list = select * from School;
    $School = $func($list);
    ob_end_clean();
    ob_implicit_flush(5);
    
    fputcsv($out, [$bom . '学生编号','学生姓名','所属区域','所属年级','所属学校']);

    // 上面整理好了对应关系
    do{
        $i = 0;
        $list = select * from student where id > $i order by id asc limit 10000
        if(!$list) break;
        foreach($list as $v){
            if($v['id'] > $i) $i = $v['id'];    // 这里可以不这样写，这些就自己优化吧
            fputcsv($out, [
                        $v['student_no'],$v['name'],
                        $area[$v['areaId']],
                        $greate[$v['gradeId']],
                        $School[$v['SchoolId']],
                    ]);
        }
    }while(true)
    
    fclose($out);
    exit();

```


**上面的问题**
1.
2.
3.


[GitHub地址]()