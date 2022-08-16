package testdata

import (
	"context"
	"time"
)

//SomeStruct
//+iface.Iface
type SomeStruct struct {
}

func (s *SomeStruct) PublicMethod(ctx context.Context, id int64) (string, error) {
	return "nil", nil
}

func (s *SomeStruct) privateMethod(ctx context.Context, time time.Time) (int32, error) {
	return 0, nil
}

//NoneMethodStruct
//+iface.Iface
type NoneMethodStruct struct {
}

//NameIfaceStruct
//+iface.Iface Name=CustomIface
type NameIfaceStruct struct {
}

//CommentsIfaceStruct
//+iface.Iface
//+iface.Comment +some.Some one="1"
//+iface.Comment +any.Any two="2"
/*+iface.Comment +more.More three="3"
four="4"
*/
type CommentsIfaceStruct struct {
}
