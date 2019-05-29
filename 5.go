package main
import "fmt"

func main(){

	//Go 语言的字符有以下两种：
	//一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
	//另一种是 rune 类型，代表一个 UTF-8 字符。当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型实际是一个 int32。

	var a byte = 'a'
        var b uint8 ='b'
   
        var c rune = '中'
        var d int32 = '国'
	
	//使用 fmt.Printf 中的%T动词可以输出变量的实际类型，使用这个方法可以查看 byte 和 rune 的本来类型
	fmt.Printf("%d %T\n",a,a)
	fmt.Printf("%d %T\n",b,b)
	fmt.Printf("%d %T\n",c,c)
	fmt.Printf("%d %T\n",d,d)

	//https://www.cnblogs.com/logo-fox/p/6473125.html
}


