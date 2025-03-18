package idgenerator

import (
	idgenerator "rania-eskristal/src/applications/id_generator"

	"github.com/google/uuid"
)

type idGeneratorImpl struct {
}

func New() idgenerator.IDGenerator {
	return &idGeneratorImpl{}
}

func (i *idGeneratorImpl) Generate() string {
	return uuid.NewString()
}
