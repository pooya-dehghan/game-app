package phoneNumber

func IsValid(phoneNumber string) bool {
	if len(phoneNumber) != 11 {
		return false
	}

	if phoneNumber[0:2] != "09" {
		return false
	}

	if _,err := strconv.Atoi(phonenumber[2:]); err == nil {
		return false
	}

	return true

}