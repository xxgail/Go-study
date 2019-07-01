package main

import "fmt"

// Go 的错误处理

type DivideError struct {
	devidee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.devidee)
}

func Divide(varDividee int, varDivider int) (res int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			devidee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if res, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", res)
	}

	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is :", errorMsg)
	}
}
