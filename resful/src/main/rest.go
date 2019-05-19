package main

import "encoding/json"
import "net/http"
import "fmt"

type student struct {
ID string
Name string
Grade int
}

var data = []student{
student{"E001", "ethan", 21},
student{"W001", "wick", 22},
student{"B001", "bourne", 23},
student{"B002", "bond", 23},
}