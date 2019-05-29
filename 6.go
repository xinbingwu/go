package main
import (
	"fmt"
	"math"
)


func main(){
	var c float32 = math.Pi //将常量保存为float32类型
	fmt.Println(c)
	fmt.Println(int(c))  //转换为int类型，浮点发生精度丢失
	fmt.Println(math.Pi)

	//注：布尔型值不能强制转换
}

