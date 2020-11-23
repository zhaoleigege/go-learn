package fsm

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func Action(ctx context.Context, run func(args ...interface{}) error) error {
	return run(ctx, 1, "name")
}

func TestAction(t *testing.T) {
	f := func(args ...interface{}) error {
		if len(args) < 3 {
			return errors.New("参数错误")
		}
		age := args[1].(int)
		name := args[2].(string)
		fmt.Printf("age: %d, name: %s", age, name)
		return nil
	}

	if err := Action(context.Background(), f); err != nil {
		t.Error(err)
	}
}

func TestStateMachine(t *testing.T) {
	sm := NewStateMachine([]*Transition{
		{"off", "open", "start", func(args ...interface{}) error {
			if len(args) < 1 {
				return errors.New("参数错误")
			}

			name := args[0].(string)
			fmt.Printf("%s开启\n", name)
			return nil
		}},
		{"open", "first", "start", func(args ...interface{}) error {
			fmt.Println("设置为1挡")
			return nil
		}},
		{"first", "off", "end", func(args ...interface{}) error {
			fmt.Println("关闭`")
			return nil
		}},
	})

	if err := sm.Process("off", "start", "test"); err != nil {
		t.Error(err)
		return
	}
}
