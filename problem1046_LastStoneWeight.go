func lastStoneWeight(stones []int) int {
    if len(stones) == 0 {
        return 0
    }

    for len(stones)>1 {
            sort.Ints(stones)
            val := getDiff(stones[len(stones)-1], stones[len(stones)-2])
            stones = append(stones[:len(stones)-2], val)
    }
    return stones[len(stones)-1]
}

func getDiff(a, b int) int {
    if a > b {
        return a-b
    }
    return b-a
}

