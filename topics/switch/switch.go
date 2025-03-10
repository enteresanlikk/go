package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	weekday := time.Now().Weekday()
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Morning")
	default:
		fmt.Println("Afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hello")
}
