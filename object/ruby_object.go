package object

import (
	"bytes"
	"strings"

	"github.com/goruby/goruby/ast"
)

// Type represents a type of an object
type Type string

const (
	EIGENCLASS_OBJ         Type = "EIGENCLASS"
	FUNCTION_OBJ           Type = "FUNCTION"
	RETURN_VALUE_OBJ       Type = "RETURN_VALUE"
	BASIC_OBJECT_OBJ       Type = "BASIC_OBJECT"
	BASIC_OBJECT_CLASS_OBJ Type = "BASIC_OBJECT_CLASS"
	OBJECT_OBJ             Type = "OBJECT"
	OBJECT_CLASS_OBJ       Type = "OBJECT_CLASS"
	CLASS_OBJ              Type = "CLASS"
	CLASS_CLASS_OBJ        Type = "CLASS_CLASS"
	ARRAY_OBJ              Type = "ARRAY"
	ARRAY_CLASS_OBJ        Type = "ARRAY_CLASS"
	INTEGER_OBJ            Type = "INTEGER"
	INTEGER_CLASS_OBJ      Type = "INTEGER_CLASS"
	STRING_OBJ             Type = "STRING"
	STRING_CLASS_OBJ       Type = "STRING_CLASS"
	SYMBOL_OBJ             Type = "SYMBOL"
	BOOLEAN_OBJ            Type = "BOOLEAN"
	BOOLEAN_CLASS_OBJ      Type = "BOOLEAN_CLASS"
	NIL_OBJ                Type = "NIL"
	NIL_CLASS_OBJ          Type = "NIL_CLASS"
	ERROR_OBJ              Type = "ERROR"
	EXCEPTION_OBJ          Type = "EXCEPTION"
	EXCEPTION_CLASS_OBJ    Type = "EXCEPTION_CLASS"
	MODULE_OBJ             Type = "MODULE"
	MODULE_CLASS_OBJ       Type = "MODULE_CLASS"
	BUILTIN_OBJ            Type = "BUILTIN"
)

type inspectable interface {
	Inspect() string
}

// RubyObject represents an object in Ruby
type RubyObject interface {
	inspectable
	Type() Type
	Class() RubyClass
}

// RubyClass represents a class in Ruby
type RubyClass interface {
	Methods() map[string]RubyMethod
	SuperClass() RubyClass
}

// RubyClassObject represents a class object in Ruby
type RubyClassObject interface {
	RubyObject
	RubyClass
}

type BuiltinFunction func(args ...RubyObject) RubyObject

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type       { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }
func (b *Builtin) Class() RubyClass { return nil }

type ReturnValue struct {
	Value RubyObject
}

func (rv *ReturnValue) Type() Type       { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Class() RubyClass { return rv.Value.Class() }

type Error struct {
	Message string
}

func (e *Error) Type() Type       { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
func (e *Error) Class() RubyClass { return nil }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        Environment
}

func (f *Function) Type() Type { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
func (f *Function) Class() RubyClass { return nil }