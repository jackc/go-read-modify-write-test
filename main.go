package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	checkInt32Struct()
	checkInt16Struct()
	checkInt32Slice()
}

func checkInt32Struct() {
	done := make(chan bool)

	var s1 struct {
		a int32
		b int32
	}

	go func() {
		for i := 0; i < 10000000; i++ {
			s1.a = s1.a + 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			s1.b = s1.b + 2
		}
		done <- true
	}()

	<-done
	<-done

	fmt.Println(s1)
	if s1.a == s1.b {
		fmt.Println("OK")
	} else {
		fmt.Println("FAIL")
	}
}

func checkInt16Struct() {
	done := make(chan bool)

	var s2 struct {
		a int16
		b int16
		c int16
		d int16
	}

	go func() {
		for i := 0; i < 10000000; i++ {
			s2.a = s2.a + 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			s2.b = s2.b + 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			s2.c = s2.c + 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			s2.d = s2.d + 2
		}
		done <- true
	}()

	<-done
	<-done
	<-done
	<-done

	fmt.Println(s2)
	if s2.a == s2.b && s2.b == s2.c && s2.c == s2.d {
		fmt.Println("OK")
	} else {
		fmt.Println("FAIL")
	}
}

func checkInt32Slice() {
	done := make(chan bool)

	a := make([]int32, 2)

	go func() {
		for i := 0; i < 10000000; i++ {
			a[0] = a[0] + 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			a[1] = a[1] + 2
		}
		done <- true
	}()

	<-done
	<-done

	fmt.Println(a)
	if a[0] == a[1] {
		fmt.Println("OK")
	} else {
		fmt.Println("FAIL")
	}
}
