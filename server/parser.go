package server

func Parser(payload []byte) (string, string) {
	pointer, par, counter := 0, make([]string, 2), 0

	for i := 0; i < len(payload); i++ {
		if payload[i] == ' ' || payload[i] == '\r' {
			par[counter] = string(payload)[pointer:i]
			pointer, counter = i+1, counter+1

			if counter == 2 {
				break
			}
		}

	}

	return par[0], par[1]
}
