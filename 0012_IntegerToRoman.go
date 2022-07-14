func intToRoman(num int) string {
    a := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    b := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    i := -1
    r := ""
    for num > 0{
        i += 1
        if num/a[i] > 0 {
            x := num/a[i]
            for j:=0;j<x;j++ {
                r = r + b[i]
            }
            num = num % a[i]
        }
    }
    
    return r
}
