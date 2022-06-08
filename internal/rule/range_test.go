package rule

import (
	"sync"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	r := &rangeRule{
		sName: "Req",
		fName: "ID",
		idx:   1,
		left:  "<",
		lVal:  getr("1"),
		right: ">=",
		rVal:  getr("3"),
	}
	t.Log(r.Check())
}

type Req struct {
	ID    *int
	Name  string
	Roles []string
}

func TestEnum(t *testing.T) {
	r := NewEnumRule("Req", "Name", "string", `{"1", "2"}`)
	t.Log(r.Check())

	r = NewEnumRule("Req", "*ID", "int", `{1, 2}`)
	t.Log(r.Check())
	// err := r.Rule(1)
	// return err
}

func TestDefault(t *testing.T) {
	t.Log(NewDefaultRule("Req", "[]string", "Roles", "make([]string, 0, 100)").Check())
}

type Err struct {
	Status int
	Code   string
	Msg    string
}

func (e *Err) WithError(err error) {

}

func TestWG(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	t.Logf("%#+v, %p", wg, &wg)
	go func(wg sync.WaitGroup) {
		t.Logf("%#+v", wg)
		wg.Done()
		t.Logf("%#+v", wg)
	}(wg)
	time.AfterFunc(time.Second, func() {
		t.Logf("%#+v", wg)
	})
	wg.Wait()
}
