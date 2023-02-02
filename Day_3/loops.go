package loops

func RepeatChar(chr string, times int) (result string) {

	for i := 0; i < times; i++ {
		result = result + chr
	}

	return
}
