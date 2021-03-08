package mapper

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
	// Method needed to assure compliance with the interface implementation
	GetSkipNumber() int
	//This is to do the assignment without modifying original code
	TransformRune2(pos int)
}

type SkipString struct {
	Value string
	Skip  int
}

func MapString(i Interface) {
	skip := i.GetSkipNumber()
	count := 1
	for pos, crune := range i.GetValueAsRuneSlice() {
		if count%skip == 0 {
			i.TransformRune(pos)
		}

		if (crune >= 'A' && crune <= 'Z') || (crune >= 'a' && crune <= 'z') {
			count++
		}
	}
}

func (ss *SkipString) GetValueAsRuneSlice() []rune {
	return []rune(ss.Value)
}

func (ss *SkipString) TransformRune(pos int) {
	runeArray := []rune(ss.Value)
	if runeArray[pos] >= 'a' && runeArray[pos] <= 'z' {
		runeArray[pos] = runeArray[pos] - 32
	} else if runeArray[pos] >= 'A' && runeArray[pos] <= 'Z' {
		runeArray[pos] = runeArray[pos] + 32
	}
	ss.Value = string(runeArray)
}

func (ss *SkipString) GetSkipNumber() int {
	return ss.Skip
}

// Original code from the assignment
func MapString2(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune2(pos)
	}
}

// Function to do the assignment without modifying the original interface
func (ss *SkipString) TransformRune2(pos int) {
	/*
		I have to reiterate over the array, because I don't have awareness inside the function
		of the count of characters that are not letters and if that position should change.
		Other implementations are keeping a counter inside
		the struct, however I don't have the guarantee it will be reset each time the process will
		be done. Other solution could be to modify the MapString, and add a getter method to the interface
		to get the number of skipping to assure that the MapString function gets the number of skipping from
		any struct that implements the Interface.
		I don't know if I should modify the original code given, however I add the other implementation
	*/

	runeArray := []rune(ss.Value)
	count := 1
	for i := range runeArray {
		if (count)%ss.Skip == 0 {
			if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
				runeArray[i] = runeArray[i] - 32
			}
		} else {
			if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
				runeArray[i] = runeArray[i] + 32
			}
		}

		if (runeArray[i] >= 'A' && runeArray[i] <= 'Z') || (runeArray[i] >= 'a' && runeArray[i] <= 'z') {
			count++
		}
	}

	ss.Value = string(runeArray)
}
