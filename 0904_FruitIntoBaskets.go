//неоптимальное, прямое в лоб решение
func totalFruit(fruits []int) int {
    max := 0
    left:=0
    for left<len(fruits) {
        b1 := 1
        t1 := fruits[left]
        t2 := -1
        b2 := 0
        for {
            if left+1 >= len(fruits) {
                break
            }
	        if fruits[left] != fruits[left+1] || left > len(fruits)-1{
	        	break
	        }

        	b1 ++
         	left++
        }

        right := left+1

        for right < len(fruits){
            if fruits[right] == t1 {
                b1 += 1
            } else if fruits[right] == t2 {
                b2 += 1
            } else if t2 == -1 {
                t2 = fruits[right]
                b2 += 1
            } else {
                break
            }
            right += 1
        }

        if max < b1 + b2 {
            max = b1 + b2
        }
        left ++
    }

    return max
}
