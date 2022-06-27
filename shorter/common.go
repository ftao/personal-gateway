package shorter


const urlChars string = "abcdefghijklmnopqrstuvwxzyABCDEFGHIJKLMNOPQRSTUVWXZY1234567890"
const shortBase int = 62

func ToShortId(seq int) string {

	parts := make([]byte, 6)

	for i := 5; i >=0 ; i-- {
		parts[i] = urlChars[seq % shortBase];
		seq = seq / shortBase
	}

	return string(parts)
}