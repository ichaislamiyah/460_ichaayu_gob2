package main

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Ir. Soekarno"

	randomValues = 10

	randomValues = true

	randomValues = []string{"Mark", "Jaehyun"}

	//Empty Interface (type assertion)
	var a interface{}

	a = 10

	a = a * 9

	//Empty Interface (Map and Slice with Empty Interface)
	rs := []interface{}{1, "Mark", true, 2, "Jaehyun", true}

	rm := map[string]interface{}{
		"Nama":   "Mark",
		"Status": true,
		"Umur":   23,
	}

	_, _ = rs, rm
}
