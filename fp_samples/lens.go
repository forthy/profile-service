package fp_samples

import OL "github.com/IBM/fp-go/optics/lens"

type Street struct {
	num  int
	name string
}

type Address struct {
	city   string
	street *Street
}

func (street *Street) GetName() string {
	return street.name
}

func (street *Street) SetName(name string) *Street {
	street.name = name

	return street
}

func (addr *Address) GetStreet() *Street {
	return addr.street
}

func (addr *Address) SetStreet(s *Street) *Address {
	addr.street = s

	return addr
}

var (
	StreetLens = OL.MakeLensRef((*Street).GetName, (*Street).SetName)
	AddrLens   = OL.MakeLensRef((*Address).GetStreet, (*Address).SetStreet)
)
