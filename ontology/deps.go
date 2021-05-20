package ontology

import (
	"go.uber.org/dig"
	"ts/config"
	"ts/fileHandler"
)

type Deps struct {
	dig.In
	Config       *config.Config
	FilesHandler fileHandler.FileHandlerInterface
}
