package flags

import (
	"fmt"
	"strings"

	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/utils"
)

var validStrEnums = map[string]map[string]string{
	LogLevelFlag: convertSliceToDummyMap(constants.AvailableLogLevels),
}

func validateStrEnum(flagName string) error {
	defer utils.PanicHandler()
	value := GetStrFlag(flagName)
	caseInsensitiveMap := make(map[string]string)
	for key, value := range validStrEnums[flagName] {
		caseInsensitiveMap[strings.ToLower(key)] = value
	}
	validEnumsValues := utils.SortedKeys(validStrEnums[flagName])
	if _, ok := caseInsensitiveMap[strings.ToLower(value)]; value != "" && !ok {
		return fmt.Errorf(
			"unknown argument for --%s: %s\nvalid arguments:\n  %s",
			flagName,
			value,
			strings.Join(validEnumsValues, "\n  "),
		)
	}
	return nil
}
