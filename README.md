# go-itermap

Go package with (ordered) interface map structure with iterator. Easy to use

    package main

    import (
        "fmt"

        "github.com/oruelle/go-itermap"
    )

    func main() {
        var v itermap.IterMap

        v.Set("c", "hello")
        v.Set("b", "world")
        v.Set(15, 24)

        fmt.Println(v.Get("a"))

        //Iterator
        for v.StartIter(); v.Next(); {
            k, v := v.Value()
            fmt.Println(k, v)
        }
    }

Output :

    <nil> false
    c hello
    b world
    15 24
