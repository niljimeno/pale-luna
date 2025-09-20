package game

func endSection() {
	message("PALE LUNA SMILES AT YOU")
	sequence := []string{DIGHOLE, DROPGOLD, FILLHOLE}
	step := 0
	maxStep := len(sequence)

	progress := func() {
		step += 1
		switch step {
		case 1:
			message("PALE LUNA SMILES AT YOU")
		case 2:
			message("PALE LUNA SMILES WIDE")
		case 3:
			message("Congratulations.")
		}
	}

	switch dial() {
	case sequence[step]:
		progress()
		if maxStep == 3 {
			return
		}

	case DIGHOLE, DROPGOLD, FILLHOLE:
		message("Not yet.")
	case HELP:
		// re-defining the order to confuse people
		listCommands(DROPGOLD, DIGHOLE, FILLHOLE)
	}
}
