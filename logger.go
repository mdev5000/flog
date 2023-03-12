package flog

import (
	"github.com/mdev5000/flog/attr"
)

type Logger interface {
	PrefixAttr(attrs *attr.Chain) Logger
	Trace(msg string, attr ...attr.Attr)
	Debug(msg string, attr ...attr.Attr)
	Info(msg string, attr ...attr.Attr)
	Warn(msg string, attr ...attr.Attr)
	Error(msg string, attr ...attr.Attr)
	Fatal(msg string, attr ...attr.Attr)
}
