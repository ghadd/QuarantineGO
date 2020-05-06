package gen

import (
	"math/rand"
	"zoo/animals"

	bhb "zoo/animals/birds/herbivores"
	bpd "zoo/animals/birds/predators"
	fhb "zoo/animals/fish/herbivores"
	fpd "zoo/animals/fish/predators"
	mhb "zoo/animals/mammals/herbivores"
	mpd "zoo/animals/mammals/predators"
)

// animals enum
const (
	eParrot  = iota
	eEagle   = iota
	eWhale   = iota
	eShark   = iota
	eGiraffe = iota
	eLion    = iota

	// goes for total number of species available
	species = iota
)

const (
	percentage = 20
)

// GenerateAnimals ...
func GenerateAnimals(an *[]animals.Animal) {
	for i := 0; i < len(*an); i++ {
		isAnimal := (rand.Intn(100) < percentage)
		if isAnimal {
			sp := rand.Intn(species)
			switch sp {
			case eParrot:
				(*an)[i] = new(bhb.Parrot)
			case eEagle:
				(*an)[i] = new(bpd.Eagle)
			case eWhale:
				(*an)[i] = new(fhb.Whale)
			case eShark:
				(*an)[i] = new(fpd.Shark)
			case eGiraffe:
				(*an)[i] = new(mhb.Giraffe)
			case eLion:
				(*an)[i] = new(mpd.Lion)
			}
		}
	}
}
