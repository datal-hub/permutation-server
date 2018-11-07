package testing

import (
	"github.com/lib/pq"

	"permutation-server/models"
)

var TestPermutation = models.Permutation{
	Uuid: "d8931bfd-a16b-4a79-a863-f9e51e728042",
	Data: pq.Int64Array{1, 2, 3},
}
