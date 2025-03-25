package fp_samples

import (
	OL "github.com/IBM/fp-go/optics/lens"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	sampleStreet  = Street{num: 220, name: "Schönaicherstr"}
	sampleAddress = Address{city: "Böblingen", street: &sampleStreet}
)

func TestLens(t *testing.T) {
	// read the value
	assert.Equal(t, sampleStreet.name, StreetLens.Get(&sampleStreet))
	// new street
	newName := "Böblingerstr"
	// update
	old := sampleStreet
	updated := StreetLens.Set(newName)(&sampleStreet)
	assert.Equal(t, old, sampleStreet)
	// validate the new name
	assert.Equal(t, newName, StreetLens.Get(updated))
}

func TestAddressCompose(t *testing.T) {
	// compose
	streetName := OL.Compose[*Address](StreetLens)(AddrLens)
	assert.Equal(t, sampleStreet.name, streetName.Get(&sampleAddress))
	// new street
	newName := "Böblingerstr"
	updated := streetName.Set(newName)(&sampleAddress)
	// check that we have not modified the original
	assert.Equal(t, sampleStreet.name, streetName.Get(&sampleAddress))
	assert.Equal(t, newName, streetName.Get(updated))
}
