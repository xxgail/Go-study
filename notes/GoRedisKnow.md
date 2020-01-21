# Go连接Redis

~~记录一下redis的类型及具体操作，还有坑~~

#### 0. 连接Redis
- 安装 `go get github.com/go-redis/redis/v7`
- 具体的连接方式可以参考[redislib.go](../sql/redisconn/redislib.go)
- 注意import引入的包`import "github.com/go-redis/redis"`

### redis数据类型

#### 1. String(字符串)
- **SET** `set(key, value)` 设置 key => value，如果已经有key就覆盖
- **SETNX** `setnx(key, value)` 设置 key => value，当且仅当key不存在，若存在则什么也不做
- **SETEX** `setex(key, time, value)` 设置 key => value，如果已经有key就覆盖，并设置过期时间为time
- **SETRANGE** `setrange(key,offset,value)` 用value覆盖给定key所储存的字符串值，从偏移量offset开始
- **MSET** `mset(key, value, key1, value1)` 同时设置一个或多个key => value
- **MSETNX** `msetnx(key, value, key1, value1)` 同时设置一个或多个key => value，当且仅当key不存在
- **APPEND** `append(key,value)` 如果key已经存在并且是一个字符串，APPEND命令将value追加到key原来的值之后。如果key不存在，APPEND就简单地将给定key设为value，就像执行SET key value一样
- **GET** `get(key)` 获取key的值，不存在就返回nil，不是字符串就报错
- **MGET** `mget(key1, key2,…, key N)` 返回所有(一个或多个)给定key的值。如果某个指定key不存在，那么返回特殊值nil。因此，该命令永不失败。
- **GETRANGE** `getrange(key,start,end)` 返回截取的子字符串，从start开始到end
- **STRLEN** `strlen(key)` 返回字符串的长度
- **INCR** `incr(key)` 名称为key的string增1操作
- **INCRBY** `incrby(key, integer)` 名称为key的string增加integer
- **DECR** `decr(key)` 名称为key的string减1操作
- **DECRBY** `decrby(key, integer)` 名称为key的string减少integer

#### 2. SET(集合)
- **SADD** `sadd(key,value)` 将一个或多个value元素加入到集合key当中，已经存在于集合的value元素将被忽略。假如key不存在，则创建一个只包含value元素作成员的集合。
- **SREM** `srem(key,value)` 移除集合key中的一个或多个value元素，不存在的value元素会被忽略。
- **SMEMBERS** `smembers(key)` 返回集合key中的所有成员。
- **SISMEMBER** `sismember(key,value)` 判断value元素是否是集合key的成员。
- **SCARD** `scard(key)` 返回集合中元素的数量
- **SMOVE** `smove(source, destination, value)` 将value元素从source集合移动到destination集合，如果destination不存在则删除source中的value
- **SPOP** `spop(key)` 移除并返回集合中的一个随机元素。
- **SRANDMEMBER** `srandmember(key)` 返回集合中的一个随机元素
- **SINTER** `sinter(key1,key2)` 返回一个集合的全部成员，该集合是所有给定集合的交集。
- **SINTERSTORE** `sinterstore(destination,key1,key2)` 此命令等同于SINTER，但它将结果保存到destination集合，而不是简单地返回结果集。如果destination集合已经存在，则将其覆盖。 destination可以是key本身。
- **SUNION** `sunion(key1,key2)` 返回一个集合的全部成员，该集合是所有给定集合的并集。


#### 3. List(集合)
- **LPUSH** `lpush(key,value)` 将一个或多个值value插入到列表key的表头。
- **LPUSHX** `lpush(key,value)` 将值value插入到列表key的表头，当且仅当key存在并且是一个列表。
- **RPUSH** `rpush(key,value)` 将一个或多个值value插入到列表key的**表尾**。
- **RPUSHX** `rpush(key,value)` 将值value插入到列表key的表尾，当且仅当key存在并且是一个列表。。
- **LPOP** `lpop(key)` 移除并返回列表key的头元素。
- **RPOP** `rpop(key)` 移除并返回列表key的尾元素。
- **LLEN** `llen(key)` 返回列表key的长度，如果key不存在返回0。
- **LRANGE** `lrange(key,start,stop)` 返回列表key中指定区间内的元素，区间以偏移量start和stop指定。
- **LREM** `lrem(key,count,value)` 根据参数count的值，移除列表中与参数value相等的元素。
          count的值可以是以下几种：
          count > 0: 从表头开始向表尾搜索，移除与value相等的元素，数量为count。
          count < 0: 从表尾开始向表头搜索，移除与value相等的元素，数量为count的绝对值。
          count = 0: 移除表中所有与value相等的值。。
- **LSET** `lset(key,index,value)` 将列表key下标为index的元素的值甚至为value。