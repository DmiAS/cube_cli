package iproto

// Абстрагируемся от способа серилизации и десерилизации сущности
type Marshaller interface {
	Marshal() ([]byte, error)
	UnMarshal([]byte) error
}

func Marshal(val Marshaller) ([]byte, error) {
	data, err := val.Marshal()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnMarshal(src []byte, val Marshaller) error {
	return val.UnMarshal(src)
}
