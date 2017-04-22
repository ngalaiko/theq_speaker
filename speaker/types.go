package speaker

type Gender string

const (
	GenderInvalid Gender = ""
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
)

func (t Gender) Clean() Gender {
	switch t {
	case GenderMale, GenderFemale:
		return t
	default:
		return GenderInvalid
	}
}
