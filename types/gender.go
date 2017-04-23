package types

type Gender string

const (
	GenderUndefined = ""
	GenderMale      = "male"
	GenderFemale    = "female"
)

func (t Gender) Clean() Gender {
	switch t {
	case GenderMale, GenderFemale:
		return t
	default:
		return GenderUndefined
	}
}

func (t Gender) String() string {
	return string(t.Clean())
}
