package types

type ID int64
type SessionKey string

func (k SessionKey) ToString() string  {
	return string(k)
}
