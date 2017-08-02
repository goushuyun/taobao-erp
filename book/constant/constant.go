package constant

//book number
var BOOK_NO = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

//give a no and find next no
func FindNextNo(fromNo string) (result string) {
	for i, no := range BOOK_NO {
		if no == fromNo {
			if i < len(BOOK_NO)-1 {
				return BOOK_NO[i+1]
			}
		}
	}
	return
}
