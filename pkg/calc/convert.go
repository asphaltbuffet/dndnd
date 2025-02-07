package calc

// ConvertToCopper converts the given amount of silver, electrum, gold, and platinum to copper.
func ConvertToCopper(silver, electrum, gold, platinum int) int {
	return gold*100 + silver*10 + electrum*5 + platinum*50
}

//
