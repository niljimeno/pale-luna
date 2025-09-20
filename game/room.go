package game

func roomSection() {
	var (
		taken    = 0
		maxTaken = 3

		rope   = true
		shovel = true
		gold   = true
	)

	showItems := func() {
		if taken == maxTaken {
			return
		}

		m := "There is "
		if gold {
			m += "GOLD in the corner"
			if taken < (maxTaken - 1) {
				m += ", along with "
			}
		}

		if shovel {
			m += "a SHOVEL"
		}

		if shovel && rope {
			m += " and "
		}

		if rope {
			m += "a ROPE"
		}

		m += "."
		message(m)
	}

	pick := func(item *bool) {
		if *item {
			taken += 1
			*item = false
			messageItemPick(taken, maxTaken)
		} else {
			warning("You've already done that.")
		}
	}

	message("You are in a dark room. Moonlight shines through the window.")
	for {
		showItems()
		message("There is a door to the EAST.")

		switch dial() {
		default:
			continue
		case PICKUPGOLD:
			pick(&gold)
		case PICKUPSHOVEL:
			pick(&shovel)
		case PICKUPROPE:
			pick(&rope)
		case GOEAST:
			if taken == maxTaken {
				return
			}
			warning("Not yet.")
		case HELP:
			listCommands(PICKUPGOLD, PICKUPSHOVEL, PICKUPROPE, GOEAST)
		}
	}
}
