package ontology

import (
	"fmt"
	"regexp"
	"ts/models"
)

const ValidRowsCount = 11

func ValidateLine(line []string) []error {

	var errors []error
	if len(line) != ValidRowsCount {
		return []error{fmt.Errorf("invalid count")}
	}

	err := validateUNSPSC(line[0])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateUNSPSCName(line[1])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateAttributeID(line[2])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateAttributeName(line[3])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateAttributeDefinition(line[4])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateDataType(line[5])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateMaxCharacterLength(line[6])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateIsRepeatable(line[7])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateMeasurementUoM(line[8])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateIsMandatory(line[9])
	if err != nil {
		errors = append(errors, err)
	}
	err = validateCodedValue(line[10])
	if err != nil {
		errors = append(errors, err)
	}
	return errors
}

func validateUNSPSC(input string) error {
	if input == "" {
		return fmt.Errorf("'UNSPSC' value is required")
	}
	if !isNumber(input) {
		return fmt.Errorf("'UNSPSC' contains not only digits: %v", input)
	}
	return nil
}

func validateUNSPSCName(input string) error {
	if input == "" {
		return fmt.Errorf("'UNSPSC Name' value is required")
	}
	return nil
}

func validateAttributeID(input string) error {
	if input == "" {
		return fmt.Errorf("'Attribute ID' value is required")
	}
	if !isNumber(input) {
		return fmt.Errorf("'Attribute ID' contains not only digits: %v", input)
	}
	return nil
}

func validateAttributeName(input string) error {
	if input == "" {
		return fmt.Errorf("'Attribute Name' value is required")
	}
	return nil
}

func validateAttributeDefinition(input string) error {
	return nil
}

func validateDataType(input string) error {
	if input == "" {
		return fmt.Errorf("'Data Type' value is required")
	}
	if input != string(models.CodedType) &&
		input != string(models.FloatType) &&
		input != string(models.NumberType) &&
		input != string(models.StringType) &&
		input != string(models.TextType) {
		return fmt.Errorf("'Data Type' field has invalid value %v", input)
	}
	return nil
}

func validateMaxCharacterLength(input string) error {
	if len(input) > 0 && !isNumber(input) {
		return fmt.Errorf("'MaxCharacterLength' contains not only digits: %v", input)
	}
	return nil
}

func validateIsRepeatable(input string) error {
	if input == "" {
		return fmt.Errorf("'Is Repeatable' value is required")
	}
	return nil
}

func validateMeasurementUoM(input string) error {
	//todo when mapping will be supported
	return nil
}

func validateIsMandatory(input string) error {
	if input == "" {
		return fmt.Errorf("'Is Mandatory' field is required")
	}
	if input != string(models.Mandatory) && input != string(models.Optional) {
		return fmt.Errorf("'Is Mandatory' field has invalid value: %v, expected: %v", input, string(models.Mandatory))
	}
	return nil
}

func validateCodedValue(input string) error {
	return nil
}

func isNumber(input string) bool {
	re := regexp.MustCompile(`[0-9]+`)
	return re.MatchString(input)
}
