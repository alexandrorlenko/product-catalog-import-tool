package ontology

import (
	"io"
	"log"
	"ts/fileHandler"
	"ts/models"
)

type RulesHandler struct {
	Reader fileHandler.FileHandlerInterface
}

func NewRulesHandler(deps Deps) *RulesHandler {
	return &RulesHandler{
		Reader: deps.FilesHandler,
	}
}

func (h *RulesHandler) UploadRules(path string) *models.Ontology {
	reader := h.Reader
	reader.InitReader(path)
	o := models.Ontology{
		Categories: map[string]*models.Category{},
	}
	pos := 0

	for {
		pos++
		line, err := reader.ReadLine()
		//todo skip 1st line
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		var errors []error
		if pos > 1 {
			errors = ValidateLine(line)
		}
		if len(errors) == 0 {
			lineCategory, lineAttribute := h.parseLine(line)
			err = o.AddCategoryAttribute(lineCategory, lineAttribute)
			if err != nil {
				log.Printf("line %v error: %v", pos, err)
			}
		} else {
			log.Printf("line %v: validation errors: %v", pos, errors)
		}
	}
	log.Printf("Rules upload finished. Proceeded %v lines, uploaded %v categories", pos, len(o.Categories))
	return &o
}

func (h *RulesHandler) parseLine(line []string) (*models.Category, *models.Attribute) {
	var c models.Category

	var attribute models.Attribute
	id := line[2]
	attribute.SetID(id)
	attribute.SetName(line[3])
	attribute.SetDefinition(line[4])
	attribute.SetDataType(line[5])
	attribute.SetMaxCharacterLength(line[6])
	attribute.SetIsRepeatable(line[7])
	attribute.SetUoM(line[8])
	attribute.SetIsMandatory(line[9])
	attribute.SetCodeValue(line[10])

	c.SetUNSPSC(line[0])
	c.SetName(line[1])

	return &c, &attribute
}
