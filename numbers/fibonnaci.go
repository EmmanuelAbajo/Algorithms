package numbers

// Fibonnaci (n) finds the nth fibonnaci number
func Fibonnaci(n int) int  {
	x,y := 0,1
	for i:=0;i<n;i++ {
		x,y = y,x+y
	}
	return x
}