package PeakIndexInAMountainArray

func peakIndexInMountainArray(arr []int) int {
	start, end := 1, len(arr)-2

	for start <= end {
		mid := (start + end) / 2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
