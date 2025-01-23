package main

import "fmt"

func main(){
	// var value any = 5
	// var value any = "Hello"
	var value any = true
	// var value any = 5.5

	printSomething(value)
	printSomething2(value)

	fmt.Println(add(1, 3))
	fmt.Println(add(1.5, 3.5))
	fmt.Println(add("Hello", " World"))
}

// any == interface{}
func printSomething(value any){
	switch value.(type){
	case int:
		fmt.Println("This is an integer, value is", value)
	case string:
		fmt.Println("This is a string, value is", value)
	case bool:
		fmt.Println("This is a boolean,	value is", value)
	case float64:	
		fmt.Println("This is a float64, value is", value)	
	default: 
		fmt.Println("This is an unknown type, value is", value)
	}
}

func printSomething2(value any){
	isDefault := true
	intValue, ok := value.(int)
	isDefault = isDefault && !ok
	if ok {
		fmt.Println("This is an integer", intValue)
	}

	stringValue, ok := value.(string)
	isDefault = isDefault && !ok
	if ok {
		fmt.Println("This is a string", stringValue)
	}
	
	boolValue, ok := value.(bool)
	isDefault = isDefault && !ok
	if ok {
		fmt.Println("This is a boolean", boolValue)
	}

	floatValue, ok := value.(float64)
	isDefault = isDefault && !ok
	if ok {
		fmt.Println("This is a float64", floatValue)
	}

	if(isDefault){
		fmt.Println("This is an unknown type")
	}	
}


// func add (a, b any) any {
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)
	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat64, aIsFloat64 := a.(float64)
	// bFloat64, bIsFloat64 := b.(float64)
	// if aIsFloat64 && bIsFloat64 {
	// 	return aFloat64 + bFloat64
	// }

	// aString, aIsString := a.(string)
	// bString, bIsString := b.(string)
	// if aIsString && bIsString {
	// 	return aString + bString
	// }
	
	// return nil;
// }

// use Generics
func add[T int|float64|string](a,b T) T {
	return a + b;
}