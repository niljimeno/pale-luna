package game

var combination = []string{
	GONORTH,
	GOSOUTH,
	GOSOUTH,
	GOEAST,
	GONORTH,
	GONORTH,
	GOSOUTH,
	GOWEST,
	GOSOUTH,
	GONORTH,
	GOSOUTH,
	GOEAST,
	GOEAST,
	// allow me to spare you some time
	// GOSOUTH,
	// GOEAST,
	// GOSOUTH,
	// GONORTH,
	// GONORTH,
	// GOWEST,
	// GONORTH,
	// GONORTH,
	// GOWEST,
	// GOWEST,
	// GOEAST,
	// GOWEST,
	// GOSOUTH,
	// GOSOUTH,
	// GOWEST,
	// GOWEST,
	// GOSOUTH,
	// GONORTH,
	// GONORTH,
	// GOEAST,
	// GOEAST,
	// GOSOUTH,
}

func forestSection() {
	step := 0
	maxStep := len(combination)

	for {
		message("You are in a forest. There are paths to the NORTH, WEST, EAST and SOUTH.")

		switch dial() {
		case combination[step]:
			step += 1
			if step == maxStep {
				return
			}
			message("PALE LUNA SMILES AT YOU.")

		case GONORTH, GOWEST, GOSOUTH, GOEAST:
			message("Are you lost?")
			step = 0
		case "cheat":
			warning(combination[step])

		case USEGOLD:
			message("Not here.")
		case USESHOVEL:
			message("Not now.")
		case USEROPE:
			message("You've already used this.")
		case HELP:
			listCommands(GONORTH, GOWEST, GOSOUTH, GOEAST, USEGOLD, USESHOVEL, USEROPE)
		}
	}
}
