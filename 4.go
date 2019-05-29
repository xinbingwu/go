package main
import "fmt"

//const的值只能是数字，字符串或者布尔值。
const str = `第一行
第二行
第三行
\r\n
`
func main(){
	fmt.Println(str)
}

