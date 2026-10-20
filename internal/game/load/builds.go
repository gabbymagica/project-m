package load

import (
	"project_m/internal/constants"
	"project_m/internal/engine"
)

var BUILD_TYPES = map[constants.BuildID]*engine.BuildType {
	constants.BuildTest1:
		engine.NewBuildType(constants.BuildTest1, true, 2, 2,
			BUILD_SPRITES[constants.BuildTest1], 1),
	constants.BuildTest2:
		engine.NewBuildType(constants.BuildTest2, true, 1, 1,
			BUILD_SPRITES[constants.BuildTest2], 1),
}
