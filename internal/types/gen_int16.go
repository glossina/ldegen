/*
 This file was autogenerated via
 -----------------------------------------------------------
 gen-builtin --declarable --native --type Int16 --name int16
 -----------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Int16("")

// Int16 represents field of type int16
type Int16 string

// Name returns field name
func (i Int16) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Int16) TypeName() string {
	return "int16"
}

// Register registers a field
func (i Int16) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddInt16(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Int16) GoName() string {
	return "int16"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["int16"] = func(fieldName string) Field {
		return Int16(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["int16"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["int16"] = struct{}{}

}
