package main

import (
	"fmt"
	"runtime"
	"strings"
	"strconv"
	"reflect"
	"log"
	"os"
)
func main(){
	fmt.Println("Hello world !")
	fmt.Println(runtime.Version())

	//变量声明了必须用,否则编译不通过
	var a = true
	fmt.Println(a)
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Println(mean)

	//默认为int类型
	vname1, vname2, vname3 := 1,2,3
	vname3 = 5
	fmt.Println(vname1,vname2,vname3)

	vname1 = int(mean)
	fmt.Println(mean)

	if(7%2 == 0){
		fmt.Println("aaaa")
	}else{
		fmt.Println("bbb")
	}

	if num:= 9 ; num < 0{
		fmt.Println("ccc")
	}else if (num < 10){
		fmt.Println("ddd")
	}else {
		fmt.Println("eee")
	}

	for i := 0 ; i < 10 ; i++{

	}

	arr := [...]int{6,7,8}
	for i,v := range arr{
		fmt.Println(i,v)
	}

	i := 2
	switch i {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println("fallthrough")
			fallthrough //执行这里的时候会继续执行后面的case语句
		case 3:
			fmt.Println(3)
		default:
			fmt.Println("default")
	}

	num := 6
	var grade string = "B"
	var marks int  = 90
	switch {
		case 0 < num && num <= 3:
			fmt.Println("0-3")
		case 4 <= num && num <= 6:
			fmt.Println("4-6")
		case 7 <= num && num <= 9:
			fmt.Println("7-9")
	}
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default : grade = "D"
	}
	switch  grade {
		case  "A":
			fmt.Println("good")
		case  "B":
			fmt.Println("fine")
		case  "C":
			fmt.Println("well")
		case  "D":
			fmt.Println("bad")
	}
	fmt.Println("-------------------")
	var arr1 [3] int
	for i,v := range arr1{
		fmt.Println(i,v)
	}
	fmt.Println("-------------------")
	arr2 := [5]int {1,2,3,4,5}
	for i,v := range arr2{
		fmt.Println(i,v)
	}
	fmt.Println("-------------------")
	arr3 := [5]string {"a","b","c","d","e"}
	for i,v := range arr3{
		fmt.Println(i,v)
	}
	fmt.Println("-------------------")
	var arr4 = [...]int {6,7,8}
	for i,v := range arr4{
		fmt.Println(i,v)
	}
	fmt.Println("-------------------")
	r := [...] int{99:-1}
	for i,v := range r{
		fmt.Println(i,v)
	}
	fmt.Println("-------------------")

	myArray := [3][4]int {{1,2,3,4,},{1,2,3,4},{1,2,3,4}}
	fmt.Println(len(myArray))
	fmt.Println(len(myArray[0]))
	fmt.Println(myArray[0])
	fmt.Println(myArray)

	for i:=0 ; i < len(myArray) ; i++{
		for j:=0 ; j<len(myArray[i]) ; j++{
			fmt.Println(i,j,myArray[i][j])
		}
	}

	fmt.Println("-------------------")
	sum1 := 0.0
	var avg float64
	xs := [6]float64{1,2,3,4,5,6}
	switch len(xs) {
		case 0:
			avg = 0
		default:
			for _,v := range xs{
				sum1 = sum1 + v
			}
			avg = sum1 / float64(len(xs))
	}
	fmt.Println(avg)

	fmt.Println("-------------------")
	str0 := `This is a raw string \n` //原生字符,原生字符原样输出(包括转义字符)
	str1 := "This is a raw string \n" //引用字符串(转义字符不会原样输出)
	fmt.Println("原生字符与引用字符的区别")
	fmt.Println(str0)
	fmt.Println(str1)
	fmt.Println("字符串连接")
	fmt.Println(str0+str1)
	fmt.Println("判断是否有字符串前缀")
	fmt.Println(strings.HasPrefix(str1,"Th"))
	fmt.Println("判断是否包含某字符串")
	fmt.Println(strings.Contains(str1,"string"))
	fmt.Println("字符串中某个字符存在的数量")
	fmt.Println(strings.Count(str1,"a"))
	fmt.Println("判断字符串中是否存在某个字符串,存在返回索引,不存在返回-1")
	fmt.Println(strings.IndexRune(str1,'a'))
	fmt.Println(strings.IndexRune(str1,'f'))
	fmt.Println("将字符串数组通过join连接起来")
	strArr := []string {"111","222","333"}
	fmt.Println(strings.Join(strArr,","))
	fmt.Println("replace的用法")
	fmt.Println(strings.Replace(str1,"a","aaa",1))//最后一个参数表示替换几个,默认替换一个
	fmt.Println("字符串拆分")
	var temp = strings.Split(str0," ")
	fmt.Println(len(temp))
	for i,v := range temp{
		fmt.Println(i,v)
	}
	fmt.Println("字符串全部转换为大小写")
	fmt.Println(strings.ToLower(str1))
	fmt.Println(str1)
	fmt.Println(strings.ToUpper(str1))

	fmt.Println("整型转字符串")
	an := 66
	newS := strconv.Itoa(an)
	fmt.Println(newS)

	fmt.Println("-------------------")
	fmt.Println("更改字符串的内容必须先将字符串转换为数组才行")
	str2 := "hello"
	str3 := []byte(str2)
	fmt.Println(str3)
	str3[0] = 'n'
	str2 = string(str3)
	fmt.Println(str2)

	fmt.Println("-------------------")
	fmt.Println(Add(1,2))
	fmt.Println(Add_Multi_Sub(1,2))
	arr5 :=[]int{1,2,3}
	fmt.Println(sums(arr5...))
	fmt.Println(sums(1,2,3))

	fmt.Println("-------------------")
	//方法调用
	var aa myInt = 4
	bb := myInt(5)
	aa.add(1,2)
	bb.add(1,2)

	fmt.Println("-------------------")
	s1 := []int{1,2,3}//slice切片
	a1 := [3]int {1,2,3}//数组a1
	s2 := a1[0:3]//切片由数组a1得到
	s3 := make([]int,10,20)//切片s3由make函数初始化,长度为10,容量为20,容量必须大于等于长度
	s4 := make([]int,10)//切片s3由make函数初始化,长度为10,默认容量为10,容量必须大于等于长度,默认初始化值为0
	fmt.Println(len(s1),cap(s1))
	fmt.Println(len(s2),cap(s2))
	fmt.Println(len(s3),cap(s3))
	fmt.Println(len(s4),cap(s4))
	for i,v := range s4{
		fmt.Println(i,v)
	}
	//slice追加元素
	fmt.Println("slice追加元素")//slice元素追加有专门的追加规则,1024前追加的话容量 *2 ,  1024之后 *1/4
	s4 = append(s4,1,2,3)
	fmt.Println(len(s4),cap(s4))
	for i,v := range s4{
		fmt.Println(i,v)
	}
	fmt.Println("slice合并")
	s4 = append(s4,s4...)
	fmt.Println(len(s4),cap(s4))
	for i,v := range s4{
		fmt.Println(i,v)
	}

	fmt.Println("-------------------")
	map1 := make(map[string]string , 5)
	map2 := make(map[string]string)//
	map3 := map[string]string{}//创建了初始化了一个空的的map，这个时候empty_map没有任何元素
	map4 := map[string]string{"a":"1","b":"2","c":"3"}
	fmt.Println(len(map1),len(map2),len(map3),len(map4))
	fmt.Println(map1,map2,map3,map4)

	fmt.Println("map元素遍历")
	for key,value := range map4{
		fmt.Println(key,value)
	}
	fmt.Println("map删除元素")
	delete(map4,"a")
	for key,value := range map4{
		fmt.Println(key,value)
	}
	fmt.Println("map中查询某个元素")
	if value,ok := map4["b"]; ok{
		fmt.Println(value+":存在")
	}

	fmt.Println("-------------------")
	alan := new(student)
	alan.name = "Alan"
	alan.age = 26
	alan.weight = 148
	alan.score = []int{10,11,12}
	fmt.Println(alan)
	fmt.Println("结构体中调用自定义类型的方法")
	dog := new(anmial)
	dog.age = 18
	dog.age.add(10,10)
	fmt.Println("结构体指针")
	pp := new(student)
	*pp = student{"qishuangming", 23, 65.0, []int{2, 3, 6}}
	fmt.Println((*pp).score)
	fmt.Println(pp.score)//Go语言自带隐式解引用

	fmt.Println("结构体的多种初始化")
	var stu1 student
	stu1.name = "Alan"
	stu1.age = 26
	stu1.weight = 148
	stu1.score = []int{10,11,12}
	upStudent(&stu1)
	fmt.Println(stu1.name)
	(&stu1).name = "Tom"//不能用(*stu1).name = "Tom"
	fmt.Println(stu1.name)

	//&等价于new
	stu2 := &student{"qishuangming", 23, 65.0, []int{2, 3, 6}}
	upStudent(stu2)
	fmt.Println(stu2.name)

	fmt.Println("-------------------")
	var phone Phone
	phone = new(IPhone)
	phone.call()
	phone = new(NokiaPhone)
	phone.call()

	fmt.Println("接口类型")
	sq1 := new(square)
	sq1.side = 5.0
	fmt.Println(sq1.area())

	//结构体实现了接口的方法则可以直接复制
	areaIntf := sq1
	fmt.Println(areaIntf.area())

	//反射
	fmt.Println("-------------------")
	var circle float32 = 6.28
	var icir interface{}
	icir = circle
	fmt.Println(reflect.ValueOf(icir))
	fmt.Println(reflect.TypeOf(icir))
	fmt.Println(reflect.ValueOf(circle))
	fmt.Println(reflect.TypeOf(circle))
	fmt.Println("反射类型对象转换为接口类型变量")
	valueof := reflect.ValueOf(icir)
	fmt.Println(valueof.Interface())
	fmt.Println(valueof.Interface().(float32))
	fmt.Println("反射赋值")
	value := reflect.ValueOf(circle)
	fmt.Println("Reflect : value = ", value)
	fmt.Println("Settability of value : ", value.CanSet())

	value2 := reflect.ValueOf(&circle)
	fmt.Println("Settability of value : ", value2.CanSet())

	value3 := value2.Elem()
	fmt.Println("Settability of value : ", value3.CanSet())

	value3.SetFloat(3.14)
	fmt.Println("Value of value3: ", value3)
	fmt.Println("value of circle: ", circle)

	//协程
	fmt.Println("-------------------")

	go loop()
	x,ok := <- complete//只有通道接收到消息之后才继续下面的运行,否则阻塞
	fmt.Println(x,ok,"完结")
	y,ok := <- complete
	fmt.Println(y,ok,"完结")

	messages := make(chan string)
	go func() {messages <- "ping"}()
	msg := <- messages
	fmt.Println(msg)

	fmt.Println("-------------------")
	httpAddr := os.Getenv("ADDR")
	fmt.Println("11111"+httpAddr)
	log.Println("Listening on ",os.Getenv("ADDR"))

	fmt.Println("-------------------")
	fmt.Println(os.Getenv("JAVA_HOME"))

}
var complete chan int = make(chan int,10)
func loop(){
	for i := 0 ; i < 10 ; i++{
		fmt.Println(i)
	}
	//运行完上面的代码必须在complete通道中发个消息,只有通道的消息被接收到了,否则这里就会被阻塞
	complete <- 0
	for i := 0 ; i < 1000 ; i++{
		fmt.Println(i)
	}
	complete <- 1
	fmt.Println("协程运行完")
}

//常规函数
func Add(i int, j int)(int){
	return i+j
}
//多值返回函数
func Add_Multi_Sub(i,j int)(int ,int ,int){
	return i+j,i-j,i*j
}
//变参函数
func sums(nums ...int)(int){
	total := 0
	for _,v := range nums{
		total += v
	}
	return total
}

type myInt int
//方法
func (a myInt)add(i int ,j int){
	fmt.Println(i+j+int(a))
}
//参数类型为一个结构体指针
func upStudent(s *student)  {
	s.name = strings.ToUpper(s.name)
}

//结构体
type student struct{
	name 	string
	age  	int
	weight	float32
	score	[]int
}
//结构体
type anmial struct{
	age		myInt
}

//接口定义
type Phone interface {
	call()
}
//两个结构体实现接口的方法
type IPhone struct {

}
func (iPhone IPhone)call(){
	fmt.Println("IPhone")
}
type NokiaPhone struct {

}
func (nokiaPhone NokiaPhone)call(){
	fmt.Println("NokiaPhone")
}

//接口
type shaper interface {
	Area()float32
}
type square struct {
	side float32
}

func (sq *square)area()float32  {
	return sq.side * sq.side
}

