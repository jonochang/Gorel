package gorel

import (
  //"fmt"
  //"reflect"
  //"strconv"
)

type MySQL struct {
  ToSql
}


func main() {
}

//func (b Binary) ToString()

//func (visitor Visitor) Visit(any interface{}) {
//    fmt.Println("test")
//  if v, ok := any.(Node); ok {
//    v.String()
//    fmt.Println("test")
//  }
//
//  switch v := any.(type) {
//    case Node:
//      fmt.Println("Node")
//    case Binary:
//      fmt.Println("Binary")
//    case Join:
//      fmt.Println("Join")
//    }
// 
//  val := reflect.ValueOf(any)
//  if val.Kind() == reflect.Struct {
//    fmt.Println("struct")
//    fmt.Println(val.NumField())
//    fmt.Println(val.Type())
//    tipe := val.Type()
//
//    fmt.Println(tipe.NumMethod())
//    fmt.Println(tipe.Method(0).Name)
//    sf := val.Field(0)
//    fmt.Println(sf)
//  }
//}

