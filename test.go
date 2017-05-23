package main
import (
"github.com/go-redis/redis"
 "fmt"
 "os"  //我们要用到os包中的env
)
func main() {
 //os.Getenv检索环境变量并返回值，如果变量是不存在的，这将是空的。
 HOME:= os.Getenv("DATABASE_URL")
 fmt.Println(HOME)
 fmt.Printf(os.Getenv("JAVA_HOME"))
}
