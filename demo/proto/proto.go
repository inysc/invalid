package proto

type NewReq struct {
	ID    int64
	Score *float64 `iv:"!nil; !20; [1,)"`
	Name  string   `iv:"s(,10]"`
	Phone string
	Role  struct{ Type string }
}
