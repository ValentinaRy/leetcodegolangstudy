package main

func sockMerchant(n int32, ar []int32) int32 {
	countMap := make(map[int32]int, n)
	for i := 0; i < int(n); i++ {
		countMap[ar[i]] = countMap[ar[i]] + 1
	}
	var matchedPairs int32 = 0
	for _, v := range countMap {
		matchedPairs += int32(v / 2)
	}
	return matchedPairs
}
