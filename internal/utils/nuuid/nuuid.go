package nuuid

import "github.com/gofrs/uuid/v5"

type NullUUID struct {
	uuid.NullUUID `swaggertype:"primitive,string"`
}

func (n NullUUID) IsZero() bool {
	return n.UUID.IsNil()
}

func FromUUID(id uuid.UUID) NullUUID {
	return NullUUID{
		uuid.NullUUID{UUID: id, Valid: true},
	}
}
