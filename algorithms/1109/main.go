package _1109

// 差分法
func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n)
	for _, booking := range bookings {
		arr[booking[0]-1] += booking[2]
		if booking[1] < n {
			arr[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		arr[i] += arr[i-1]
	}
	return arr
}

func corpFlightBookingsSolution2(bookings [][]int, n int) []int {
	arr := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0] - 1; j < bookings[i][1]; j++ {
			arr[j] += bookings[i][2]
		}
	}
	return arr
}
