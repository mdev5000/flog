package attr

type Attr struct {
	Key   string
	Value interface{}
}

func String(k, v string) Attr {
	return Attr{k, v}
}

func Int(k string, v int) Attr {
	return Attr{k, v}
}

func Int64(k string, v int64) Attr {
	return Attr{k, v}
}

func Int32(k string, v int32) Attr {
	return Attr{k, v}
}

func Uint(k string, v uint) Attr {
	return Attr{k, v}
}

func Uint64(k string, v uint64) Attr {
	return Attr{k, v}
}

func Uint32(k string, v uint32) Attr {
	return Attr{k, v}
}

func Bool(k string, v bool) Attr {
	return Attr{k, v}
}

func Err(err error) Attr {
	return Attr{"error", err}
}

type AnyAttr struct {
	A any
}

func Interface(k string, a any) Attr {
	return Attr{k, AnyAttr{a}}
}
