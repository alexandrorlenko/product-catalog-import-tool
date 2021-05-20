package models

import (
	"fmt"
	"strconv"
)

type DataType string
type IsMandatory string

const (
	Mandatory IsMandatory = "Mandatory"
	Optional  IsMandatory = "Optional"

	// Data Types
	CodedType  DataType = "Coded"
	FloatType  DataType = "Float"
	NumberType DataType = "Number"
	StringType DataType = "String"
	TextType   DataType = "Text"
)

//ontology the index
type AttributeConfig struct {
	ID                 string
	Name               string
	Definition         string
	DataType           DataType
	MaxCharacterLength int
	IsRepeatable       bool
	MeasurementUoM     string
	IsMandatory        bool
	CodedValue         string
}

type CategoryConfig struct {
	UNSPSC     string
	Name       string
	Attributes map[string]*AttributeConfig
}

type OntologyConfig struct {
	Categories map[string]*CategoryConfig
}

// Raw object
type Ontology struct {
	Categories map[string]*Category
}

func (o *Ontology) AddCategoryAttribute(c *Category, a *Attribute) error {
	_, ok := o.Categories[c.UNSPSC]
	if !ok {
		o.Categories[c.UNSPSC] = c
	} else {
		if o.Categories[c.UNSPSC].Attributes == nil {
			o.Categories[c.UNSPSC].Attributes = map[string]*Attribute{}
		}
		_, ok = o.Categories[c.UNSPSC].Attributes[a.ID]
		if !ok {
			o.Categories[c.UNSPSC].Attributes[a.ID] = a
		} else {
			return fmt.Errorf("attribute %v is allready exists in category %v", a.ID, c.UNSPSC)
		}
	}
	return nil
}

func (o *Ontology) ToConfig() *OntologyConfig {
	configs := make(map[string]*CategoryConfig, len(o.Categories))
	for i, v := range o.Categories {
		configs[i] = v.ToConfig()
	}
	return &OntologyConfig{
		Categories: configs,
	}
}

type Category struct {
	UNSPSC     string
	Name       string
	Attributes map[string]*Attribute
}

func (c *Category) SetUNSPSC(input string) {
	c.UNSPSC = input
}

func (c *Category) SetName(input string) {
	c.Name = input
}

func (c *Category) AddAttribute(a *Attribute) {
	c.Attributes[a.ID] = a
}

func (c *Category) ToConfig() *CategoryConfig {
	configs := make(map[string]*AttributeConfig, len(c.Attributes))
	for i, v := range c.Attributes {
		configs[i] = v.ToConfig()
	}
	return &CategoryConfig{
		UNSPSC:     c.UNSPSC,
		Name:       c.Name,
		Attributes: configs,
	}
}

//--------------//
type Attribute struct {
	ID                 string
	Name               string
	Definition         string
	DataType           DataType
	MaxCharacterLength int
	IsRepeatable       bool
	MeasurementUoM     string
	IsMandatory        bool
	CodedValue         string
}

func (a *Attribute) SetID(input string) {
	a.ID = input
}

func (a *Attribute) SetName(input string) {
	a.Name = input
}

func (a *Attribute) SetDefinition(input string) {
	if input != "" {
		a.Definition = input
	}
}

func (a *Attribute) SetDataType(input string) {
	a.DataType = DataType(input)
}

func (a *Attribute) SetMaxCharacterLength(input string) {
	if input != "" {
		res, _ := strconv.Atoi(input)
		a.MaxCharacterLength = res
	}
}

func (a *Attribute) SetIsRepeatable(input string) {
	if input == "No" {
		a.IsRepeatable = false
	} else {
		a.IsRepeatable = true
	}
}

func (a *Attribute) SetUoM(input string) {
	if input != "" {
		a.MeasurementUoM = input
	}
}

func (a *Attribute) SetIsMandatory(input string) {
	if input == string(Mandatory) {
		a.IsMandatory = true
	} else {
		a.IsMandatory = false
	}
}

func (a *Attribute) SetCodeValue(input string) {
	if input != "" {
		a.CodedValue = input
	}
}

func (a *Attribute) ToConfig() *AttributeConfig {
	return &AttributeConfig{
		ID:                 a.ID,
		Name:               a.Name,
		Definition:         a.Definition,
		DataType:           a.DataType,
		MaxCharacterLength: a.MaxCharacterLength,
		IsRepeatable:       a.IsRepeatable,
		MeasurementUoM:     a.MeasurementUoM,
		IsMandatory:        a.IsMandatory,
		CodedValue:         a.CodedValue,
	}
}
