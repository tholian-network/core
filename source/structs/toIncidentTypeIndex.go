package structs

func toIncidentTypeIndex(typ string) int {

	result := int(0)

	switch typ {
	case "Recon":   result = 1
	case "Intel":   result = 2
	case "Conquer": result = 3
	case "Persist": result = 4
	case "Exfil":   result = 5
	case "Destroy": result = 6
	}

	return result

}
