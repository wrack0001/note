# php操作kafka之实例操作(生产者和消费者)

[php kafka手册地址](https://arnaud.le-blanc.net/php-rdkafka/phpdoc/book.rdkafka.html)

## 1.生产者

```

    <?php
    
    $rk = new RdKafka\Producer();
    $rk->setLogLevel(LOG_DEBUG);
    $rk->addBrokers("127.0.0.1");   // kafka服务器地址
    // $rk->addBrokers("10.0.0.1,10.0.0.2");    // 多服务器地址写法
    
    $topic = $rk->newTopic("nginx_log");     // topic 的名称
    
    for ($i = 0; $i < 10; $i++) {
        $topic->produce(RD_KAFKA_PARTITION_UA, 0, "Message $i");
    }


```
[更多例子](https://arnaud.le-blanc.net/php-rdkafka/phpdoc/rdkafka.examples-producer.html)

## 2.消费者

```
    <?php
    
    $rk = new RdKafka\Consumer();
    $rk->setLogLevel(LOG_DEBUG);
    $rk->addBrokers("127.0.0.1");
    
    $topic = $rk->newTopic("nginx_log");
    
    $topic->consumeStart(0, RD_KAFKA_OFFSET_BEGINNING);
    
    while (true) {
        $msg = $topic->consume(0, 1000);
        if ($msg->err) {
            echo $msg->errstr(), "\n";
            break;
        } else {
            echo $msg->payload, "\n";
        }
    }


```

**更多的内容参考上面贴出的手册，里面的内容很全的👉**



[GitHub地址](https://github.com/wrack0001/note/blob/master/php/php%E6%93%8D%E4%BD%9Ckafka%E4%B9%8B%E5%AE%9E%E4%BE%8B%E6%93%8D%E4%BD%9C%E2%80%94%E2%80%94%E7%94%9F%E4%BA%A7%E8%80%85%E5%92%8C%E6%B6%88%E8%B4%B9%E8%80%85.md)

