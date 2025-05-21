package main

import "fmt"

func main() {
    // Basic types
    var age int = 25
    var name string = "Tilak"
    var isStudent bool = true
    var grade float64 = 89.5

    // Byte and rune
    var b byte = 'A'
    var r rune = '✓'

    // Array
    var scores [3]int = [3]int{85, 90, 95}

    // Slice
    subjects := []string{"Math", "Science", "History"}

    // Map
    student := map[string]string{
        "name":  "Tilak",
        "major": "CS",
    }

    // Struct
    type User struct {
        username string
        age      int
    }
    user1 := User{"tilak123", 22}

    // Pointer
    var ptr *int = &age

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("Grade:", grade)
    fmt.Println("Byte:", b)
    fmt.Println("Rune:", r)
    fmt.Println("Scores:", scores)
    fmt.Println("Subjects:", subjects)
    fmt.Println("Student Info:", student)
    fmt.Println("User Struct:", user1)
    fmt.Println("Pointer to Age:", ptr, "Value:", *ptr)
}