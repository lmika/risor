package object

import (
	"fmt"
	"strings"

	"github.com/cloudcmds/tamarin/v2/op"
)

// Partial is a partially applied function
type Partial struct {
	fn   Object
	args []Object
}

func (p *Partial) Function() Object {
	return p.fn
}

func (p *Partial) Args() []Object {
	return p.args
}

func (p *Partial) Type() Type {
	return PARTIAL
}

func (p *Partial) Inspect() string {
	var args []string
	for _, arg := range p.args {
		args = append(args, arg.Inspect())
	}
	return fmt.Sprintf("partial(%s, %s)", p.fn.Inspect(), strings.Join(args, ", "))
}

func (p *Partial) Interface() interface{} {
	return p.fn
}

func (p *Partial) Equals(other Object) Object {
	if p == other {
		return True
	}
	return False
}

func (p *Partial) GetAttr(name string) (Object, bool) {
	return nil, false
}

func (p *Partial) IsTruthy() bool {
	return true
}

func (p *Partial) RunOperation(opType op.BinaryOpType, right Object) Object {
	return NewError(fmt.Errorf("eval error: unsupported operation for nil: %v", opType))
}

func (p *Partial) Cost() int {
	return 0
}

func NewPartial(fn Object, args []Object) *Partial {
	return &Partial{
		fn:   fn,
		args: args,
	}
}
