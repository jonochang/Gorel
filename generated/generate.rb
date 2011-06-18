# Arel::Nodes.constants.map{|c| Arel::Nodes.const_get(c).ancestors[1]}.uniq
# Arel::Nodes.constants.map{|c| Arel::Nodes.const_get(c).ancestors }.flatten.uniq
# 
# [Arel::Nodes::Binary, Arel::Nodes::Unary, Arel::Nodes::Equality, Arel::Expression, Arel::Nodes::Function, Arel::Predications, Arel::Nodes::Join].map{|b| Object.constants.select{|c| Object.const_get(c).class == Class && Object.const_get(c).ancestors.include?(b)} }
# 
# Object.constants.select{|c| Object.const_get(c).class == Class && Object.const_get(c).ancestors.include?(b)}
# Object.constants.select{|c| Object.const_get(c).class == Class && Object.const_get(c).ancestors.include?(Arel::Nodes::Node) }
# Object.constants.map{|c| }.flatten.uniq

class Symbol
  include Comparable

  def <=>(other)
    self.to_s <=> other.to_s
  end
end

@all = [
[:Binary, :Equality, :Between, :NotEqual, :Assignment, :Or, :And, :As, :GreaterThan, :GreaterThanOrEqual, :LessThan, :LessThanOrEqual, :Matches, :DoesNotMatch, :In, :NotIn, :Ordering, :Values, :DeleteStatement, :TableAlias, :Except, :Intersect, :Union, :UnionAll],
[:Unary, :Not, :Lock, :Offset, :Limit, :Top, :Having, :UnqualifiedColumn, :Group, :Grouping, :On], 
[:Equality, :In], 
[:Function, :Count, :Sum, :Exists, :Max, :Min, :Avg], 
[:Function, :Count, :Sum, :Exists, :Max, :Min, :Avg], 
[:SqlLiteral], 
[:Join, :InnerJoin, :OuterJoin, :StringJoin]] 

@child_list = {
:Binary => [:Between, :NotEqual, :Assignment, :Or, :As, :GreaterThan, :GreaterThanOrEqual, :LessThan, :LessThanOrEqual, :Matches, :DoesNotMatch, :NotIn, :Ordering, :Values, :DeleteStatement, :TableAlias, :Except, :Intersect, :Union, :UnionAll],
:Unary => [:Not, :Lock, :Offset, :Limit, :Top, :Having, :UnqualifiedColumn, :Group, :Grouping, :On],
:Equality => [:In],
:Function => [:Count, :Sum, :Exists, :Max, :Min, :Avg],
:And => [],
:Join => [:InnerJoin, :OuterJoin, :StringJoin] 
}

@parent_list = {
  :Binary => "
type Binary struct {
  left Node
  right Node
}
",
  :Unary => "
type Unary struct {
  expression Node
}
",
  :Function => "
type Function struct {
  expressions []Node
  alias Literal
  distinct bool
}
",
  :Join => "
type Join struct {Binary}
",
  :And => "
type And struct {
  children []Node
}

func (n And) Visit(v Visitor) (s string) {
  s = v.GetAnd(n)
  return
}
", 
  :Equality => "
type Equality struct {Binary}

func (n Equality) Visit(v Visitor) (s string) {
  s = v.GetEquality(n)
  return
}
"
}

def generate_nodes_visitor
  header = "package gorel

type Node interface {
  Visit(v Visitor) string
}

type Literal struct {
  value interface{}
}

func (n Literal) Visit(v Visitor) (s string) {
  s = v.GetLiteral(n)
  return
}
"

  nodes_file = "generated/nodes.go"
  File.open(nodes_file, 'w') {|f| f.write(header) }
  puts "generating nodes..."

  @parent_list.keys.sort.each do |key|
    puts "generating struct for #{key}"
    s = @parent_list[key]
    File.open(nodes_file, 'a') {|f| f.write(s) }
  end

  @child_list.keys.sort.each do |parent|
    children = @child_list[parent]
    File.open(nodes_file, 'a') {|f| f.write("\n//-----------------#{parent}----------------\n") }
    children.each do |child|
      puts "generating struct for #{child}"
      s = 
"type #{child} struct {#{parent}}

func (n #{child}) Visit(v Visitor) (s string) {
  s = v.Get#{child}(n)
  return
}\n
"
      File.open(nodes_file, 'a') {|f| f.write(s) }
    end
  end

  visitor_file = "generated/visitor.go"
  visitor_interface = "package gorel

type Visitor interface {
  GetLiteral(n Literal) string
  GetEquality(n Equality) string
  GetAnd(n And) string
"
  File.open(visitor_file, 'w') {|f| f.write(visitor_interface) }
  puts "generating interface visitor..."

  @child_list.keys.sort.each do |parent|
    children = @child_list[parent]
    File.open(visitor_file, 'a') {|f| f.write("\n//-----------------#{parent}----------------\n") }
    children.each do |child|
      puts "generating interface for #{child} in visitor"
      s = "  Get#{child}(n #{child}) string\n"
      File.open(visitor_file, 'a') {|f| f.write(s) }
    end
  end

  File.open(visitor_file, 'a') {|f| f.write("}") }
end

def generate_visitor filename, type
  header = "package gorel

import (
  \"fmt\"
  \"strconv\"
  \"strings\"
)

type #{type} struct {
}

func (c #{type}) GetLiteral(n Literal) (s string) {
  switch val := n.value.(type) {
    case string:
      s = val
    case bool:
      s = strconv.Btoa(val)
    case int:
      s = strconv.Itoa(val)
    case int64:
      s = strconv.Itoa64(val)      
    case uint:
      s = strconv.Uitoa(val)
    case uint64:
      s = strconv.Uitoa64(val)
    case float32:
      s = strconv.Ftoa32(val, 'f', -1)
    case float64:
      s = strconv.Ftoa64(val, 'f', -1)  
  }
  return
}

func (b #{type}) VisitNodes(nodes []Node) (s []string) {
  s = make([]string, 0)
  for i:=0; i < len(nodes); i++ {
    if (nodes[i] != nil) {
      s = append(s, nodes[i].Visit(b))
    }
  }
  return
}

func (b #{type}) VisitNodesString(nodes []Node, sep string) (s string) {
  results := b.VisitNodes(nodes)
  s = strings.Join(results, sep)
  return
}
"
  File.open(filename, 'w') {|f| f.write(header) }
  
  @child_list.keys.sort.each do |parent|
    children = @child_list[parent]
    File.open(filename, 'a') {|f| f.write("\n//-----------------#{parent}----------------\n") }
    case parent
#       when :Binary
#         s = "func (c #{type}) GetBinary(n Binary) (ls string, rs string) {
#   ls = n.left.Visit(c)
#   rs = n.right.Visit(c)
#   return
# }\n\n"
#         File.open(filename, 'a') {|f| f.write(s) }
    when :Equality
      s = "func (c #{type}) GetEquality(n Equality) (s string) {
  ls := n.left.Visit(c)
  rs := n.right.Visit(c)
  s = fmt.Sprintf(\"%v = %v\", ls, rs)
  return
}\n
"
      File.open(filename, 'a') {|f| f.write(s) }
    when :And
      s = "func (c #{type}) GetAnd(n And) (s string) {
  results := c.VisitNodes(n.children)
  s = strings.Join(results, \" AND \")
  return
}\n
"
            File.open(filename, 'a') {|f| f.write(s) }
    end
    
    children.each do |child|
      puts "generating function #{child} to satisfy Visitor interface"
      s = "func (c #{type}) Get#{child}(n #{child}) (s string) {\n"
      File.open(filename, 'a') {|f| f.write(s) }
      case parent
        when :Binary, :Equality, :Join
#           s = '  ls, rs := c.GetBinary(n)
#   s = fmt.Sprintf("%v * %v", ls, rs)
#   return
# '
          s = "  ls := \"\"
  if n.left != nil {
    ls = n.left.Visit(c)
  }
  
  rs := \"\"
  if n.right != nil {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf(\"%v * %v\", ls, rs)
  return
"
          File.open(filename, 'a') {|f| f.write(s) }
          
        when :Unary
          s = "  expr := n.expression.Visit(c)
  s = expr
  return s\n\n"
          File.open(filename, 'a') {|f| f.write(s) }

        when :Function
          s = "  expressions := c.VisitNodesString(n.expressions, \", \")
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf(\"%v * %v * %v\", expressions, alias, distinct)
  return s\n\n"
          File.open(filename, 'a') {|f| f.write(s) }
          
      end
      s = "}\n\n"
      File.open(filename, 'a') {|f| f.write(s) }
    end
  end
end

def generate_visitor_test filename, type
    header = "package gorel

  import (
    \"testing\"
  )
  
  func TestGetLiteral(t *testing.T) {
    v := new(#{type})
    l := new(Literal)
          
    t.Log(\"Test string\")
    l.value = \"test\"
    s := v.GetLiteral(*l)
    if s != \"test\" {
      t.Errorf(\"failed test string\")
    }
    
    t.Log(\"Test boolean\")
    l.value = false
    s = v.GetLiteral(*l)
    if s != \"false\" {
      t.Errorf(\"failed test boolean\")
    }
    
    t.Log(\"Test int\")
    l.value = 123
    s = v.GetLiteral(*l)
    if s != \"123\" {
      t.Errorf(\"failed test int\")
    }
    
    t.Log(\"Test float\")
    l.value = 123.3
    s = v.GetLiteral(*l)
    if s != \"123.3\" {
      t.Errorf(\"failed test float\")
    }
  }
    
  func TestVisitNilNodes(t *testing.T) {
    v := new(#{type})
  	s := v.VisitNodes(nil)

  	if len(s) != 0 {
  		t.Errorf(\"failed to visit nil nodes\")
  	}

  	nodes := make([]Node, 3)
  	s = v.VisitNodes(nodes)
  	if len(s) != 0 {
  		t.Errorf(\"failed to skip nil nodes\")
  	}
  }
"  
  File.open(filename, 'w') {|f| f.write(header) }

    @child_list.keys.sort.each do |parent|
      children = @child_list[parent]
      File.open(filename, 'a') {|f| f.write("\n//-----------------#{parent}----------------\n") }
      case parent
  #       when :Binary
  #         s = "func (c #{type}) GetBinary(n Binary) (ls string, rs string) {
  #   ls = n.left.Visit(c)
  #   rs = n.right.Visit(c)
  #   return
  # }\n\n"
  #         File.open(filename, 'a') {|f| f.write(s) }
      when :Equality
        s = "func TestGetEquality(t *testing.T) {
    v := new(#{type})
    n := new(Equality)
    
    ll := new(Literal)
    ll.value = 1
    n.left = ll
    
    rl := new(Literal)
    rl.value = 2
    n.right = rl
    
    s := v.GetEquality(*n)
    if s != \"1 = 2\" {
      t.Errorf(\"failed to get Equality\")
    }
}\n
"
        File.open(filename, 'a') {|f| f.write(s) }
      when :And
        s = "func TestGetAnd(t *testing.T) {
    v := new(#{type})
    n := new(And)
    
    s := v.GetAnd(*n)
    if s != \"1 = 2\" {
      t.Errorf(\"failed to get And\")
    }
}\n"
        File.open(filename, 'a') {|f| f.write(s) }
      end

      children.each do |child|
        puts "generating function #{child} to satisfy Visitor interface"
        s = "func TestGet#{child}(t *testing.T) {
  v := new(#{type})
  n := new(#{child})
  
  s := v.Get#{child}(*n)
  if s != \"\" {
    t.Errorf(\"failed to get #{child} \")
  }
}
"
        File.open(filename, 'a') {|f| f.write(s) }
      end
    end
end

case ARGV[0] 
  when "visitor"
    generate_visitor ARGV[1], ARGV[2]
  when "visitor_test"
    generate_visitor_test ARGV[1], ARGV[2]
  when "-h"
    puts "use generate.rb visitor [filename] [type] to create a new visitor template
use generate.rb visitor_test [filename] [type] to create a new visitor test template
use generate.rb to generate nodes and visitor interface
    "
  else
    generate_nodes_visitor
end