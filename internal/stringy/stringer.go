package stringy

import "strings"

type Stringer string

func New() *Stringer {
	return new(Stringer)
}

func (s Stringer) String() string {
	return string(s)
}

func (s Stringer) Trim() Stringer {
	return Stringer(strings.TrimSpace(string(s)))
}

func (s Stringer) Lower() Stringer {
	return Stringer(strings.ToLower(string(s)))
}

func (s Stringer) ReplaceAll(old string, new string) Stringer {
	return Stringer(strings.ReplaceAll(string(s), old, new))
}

func (s Stringer) TrimPrefix(str string, prefix string) Stringer {
	return Stringer(strings.TrimPrefix(str, prefix))
}
