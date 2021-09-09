package imp

import(
    "fmt"
)

func Show(name string, age int) string {
    fmt.Printf("My Name is %s, My age is %d\n", name, age)
    return fmt.Sprintf("My Name is %s, My age is %d\n", name, age)
}
