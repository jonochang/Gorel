package gorel

import (
  "fmt"
  "reflect"
  "strconv"
)

type Value interface {}

type Node interface {
  Visit(v IVisitor) string
}

type Literal struct {
  value interface{}
}

type Binary struct {
  left Node
  right Node
}

type Equality struct { Binary }
type Unary struct {
  expression Node
}

type Join struct { Binary }

type ToSql struct {

}

type IVisitor interface {
  Accept(n interface{})
  GetLiteral(l Literal) string
  GetBinary(b Binary) string
  GetUnary(u Unary) string
  GetEquality(e Equality) string
}

type MySQL struct {
  ToSql
}


func main() {
}

func (l Literal) Visit(v IVisitor) (s string) {
  s = v.GetLiteral(l)
  return
}

func (b Binary) Visit(v IVisitor) (s string) {
  s = v.GetBinary(b)
  return
}

func (u Unary) Visit(v IVisitor) (s string) {
  s = v.GetUnary(u)
  return
}

func (e Equality) Visit(v IVisitor) (s string){
  s = v.GetEquality(e)
  return
}

func (v ToSql) Accept(any interface{}) {
  switch n := any.(type) {
    case Node:
      fmt.Println("node")
      val := reflect.ValueOf(any)
      fmt.Println(val.Type())
      n.Visit(v)
  }
}

func (v ToSql) GetLiteral(l Literal) (s string) {
  switch val := l.value.(type) {
    case string:
      s = val
    case int:
      s = strconv.Itoa(val)
  }
  fmt.Println(s)
  return
}

func (v ToSql) GetBinary(b Binary) (s string) {
  ls := b.left.Visit(v)
  rs := b.right.Visit(v)

  s = fmt.Sprintf("%v, %v", ls, rs)
  return
}

func (v ToSql) GetUnary(l Unary) (s string) {
  s = l.expression.Visit(v)
  return 
}

func (v ToSql) GetEquality(e Equality) (s string) {
  ls := e.left.Visit(v)
  rs := e.right.Visit(v)

  s = fmt.Sprintf("%v = %v", ls, rs)
  
  return
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

