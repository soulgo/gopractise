package main

func main() {
	/*
		//	1、字符串转换为int类型
		s1 := "100"
		i1,err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println("can't convert to int")
		} else {
			fmt.Printf("type:%T value:%#v\n", i1, i1)
		}
	*/
	/*
		// 2、int转换为字符串
		i2 := 200
		s2 := strconv.Itoa(i2)
		fmt.Printf("type:%T value:%#v\n", s2, s2)
	*/
	/*
		//	3、转为bool值：接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；
		b1 := "a"
		b2,_ := strconv.ParseBool(b1)
		fmt.Println(b2)
	*/
	/*
		//	4、字符串转为int
		i2 := "200"
		i3,_ := strconv.ParseInt(i2,10,64)
		fmt.Println(i3)
	*/
	/*
		// 4、appendTP函数
		//a1 := []byte("int (base 10):")
		//a1 = strconv.AppendInt(a1,-42,10)
		//fmt.Println(a1)
		s1 := make([]byte,0)
		s2 := strconv.AppendInt(s1,20,10)
		fmt.Println(string(s2))
		s3 := strconv.AppendFloat(s1,3.1415926,'f',4,64)
		fmt.Println(string(s3))
	*/
}
