# Go连接MySQL

~~最近公司的项目里用到了go去查询同步MySQL，所以记录一下，其实大部分是原生SQL，但是强类型肯定有坑~~

#### 0. 连接MySQL
- 具体的连接方式可以参考[mysqllib.go](../sql/mysqlconn/mysqllib.go)
- 注意import引入的包`import _ "github.com/go-sql-driver/mysql"`

#### 1. 增加(Insert)
- 直接上代码
```go
package main
import (
       	"fmt"
       	"github.com/xxgail/sql/mysqlconn"
)

func main(){
    var name string = "gai"
    var age int = 22
    tx, err := mysqlconn.GetMysqlConn().Begin()
    if err != nil {
    	fmt.Println("tx fail")
    }
    query := "INSERT INTO students (`name`,`age`) VALUES (?,?)"
    stem, err := tx.Prepare(query)
    if err != nil {
    	fmt.Println("Prepare fail")
    }
    res, err := stem.Exec(name, age)
    if err != nil {
    	fmt.Println("Exec fail")
    }
    tx.Commit()
    fmt.Println(res.LastInsertId()) // 返回插入的id
}
```

#### 2. 更新(Update)
- 直接上代码
```go
package main
import (
"fmt"
"github.com/xxgail/sql/mysqlconn"
)

func main()  {
    var studentId = 1
    var name = "gaigai"
    //开启事务 
    tx, err := mysqlconn.GetMysqlConn().Begin()
    if err != nil {
    	fmt.Println("tx fail")
    }
    //准备sql语句
    stmt, err := tx.Prepare("UPDATE students SET name = ? WHERE id = ?")
    if err != nil {
    	fmt.Println("Prepare fail")
    }
    //设置参数以及执行sql语句
    res, err := stmt.Exec(name,studentId)
    if err != nil {
    	fmt.Println("Exec fail")
    }
    //提交事务
    tx.Commit()
    fmt.Println(res.RowsAffected()) // 返回影响函数
}
```

#### 3. 查询单条记录(QueryRow)
- 先上代码
```go
package main
import (
"database/sql"
"fmt"
"github.com/xxgail/sql/mysqlconn"
"strconv"
)

type Student struct {
    name    string
    age     sql.NullInt64
}

func main()  {
    var student Student
    mysqlClient := mysqlconn.GetMysqlConn()
   
    var studentId int = 1
    query := "SELECT name,age FROM students WHERE id = " + strconv.Itoa(studentId)
    fmt.Println(query)
    err := mysqlClient.QueryRow(query).Scan(&student.name, &student.age)
    if err != nil {
    	fmt.Println("initUser 查询数据库单条用户信息 出错：", err)
    }
    var age int64
    if student.age.Valid {
        age = student.age.Int64
    }
    var studentInfo map[string]interface{}
    studentInfo["name"] = student.name
    studentInfo["age"] = age
    fmt.Println(studentInfo)
}
```
- 为了防止数据的中设置的字段为空，所以定义结构体接收类型可以设置为sql.Null*，即有值就存，null就存该类型的初始值。


#### 4. 查询(Query)
- 代码
```go
package main
import (
"fmt"
"github.com/xxgail/sql/mysqlconn"
)

type Student struct {
    id      int64
    name    string
    age     int64
}

func main()  {
    username := "gai"
    mysqlClient := mysqlconn.GetMysqlConn()
    query := "SELECT id,name,age FROM students WHERE name = " + username + "ORDER BY RAND() LIMIT 100"
    rows, err := mysqlClient.Query(query)   
    if err != nil { 
        fmt.Println(err)
    }
    defer rows.Close()
   
    var students []Student  
    for rows.Next() {
    	var student Student
    	if err := rows.Scan(&student.id, &student.name,&student.age); err != nil {
            fmt.Println(err)
    	}
    	students = append(students,student)
    }
    fmt.Println(students)
}
```

#### 5. 删除(Delete)
- 直接上代码
```go
package main
import (
"fmt"
"github.com/xxgail/sql/mysqlconn"
)

func main()  {
    var studentId = 1
    //开启事务 
    tx, err := mysqlconn.GetMysqlConn().Begin()
    if err != nil {
    	fmt.Println("tx fail")
    }
    //准备sql语句
    stmt, err := tx.Prepare("DELETE FROM students WHERE id = ?")
    if err != nil{
    	fmt.Println("Prepare fail")
    }
    //设置参数以及执行sql语句
    res, err := stmt.Exec(studentId)
    if err != nil{
    	fmt.Println("Exec fail")
    }
    //提交事务
    tx.Commit()
    fmt.Println(res.RowsAffected()) // 返回影响函数
}
```