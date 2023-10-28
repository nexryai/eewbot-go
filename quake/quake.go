package quake

func IsEmergency(DispIntensity string) bool {
	// 5+とかになる可能性があるのでintにできない
	switch DispIntensity {
	case "1":
		return false
	case "2":
		return false
	case "3":
		return false
	case "4":
		return false
	default:
		// 5- ~ 7 は非常事態
		return true
	}
}
