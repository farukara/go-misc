package main

import ("fmt"
        "os/exec"
)

func main() {
    out, err := exec.Command("zsh", "-c", "brew list").Output()
    if err != nil {
        fmt.Println("error: ", err)
    }
    output := string(out)
    fmt.Println(output)
    fmt.Println(out)
    fmt.Printf("out format: %T", out)
}
