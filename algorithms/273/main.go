func numberToWords(num int) string {
    if num < 100 {
        return translateUpToHundred(num)
    } else if num < 1000 {
        return translateFromRecursive(num, 100, "Hundred")
    } else if num < 1000000 {
        return translateFromRecursive(num, 1000, "Thousand")
    } else if num < 1000000000 {
        return translateFromRecursive(num, 1000000, "Million")
    } else {
        return translateFromRecursive(num, 1000000000, "Billion")
    }
}

func translateFromRecursive(num int, unit int, unitString string) string {
    for i := 1 ; i < 1000 ; i++ {
        if num >= i * unit && num < (i+1) * unit {
            switch num {
            case i * unit:
                return numberToWords(i) + " " + unitString
            default:
                return numberToWords(i) + " " + unitString + " " + numberToWords(num- (i * unit))
            }
        }
    }
    return "Unknown value"    
}

func translateUpToHundred(num int) string {
    if num <= 10 {
        return translateUpToTen(num)
    } else if num < 20 {
        return translateTeen(num)
    } else if num < 100 {
        return translateFromTwentyToHundred(num)
    }
    return "Unkown value"
}

func translateFromTwentyToHundred(num int) string {
    dict := map[int]string{
        2: "Twenty",
        3: "Thirty",
        4: "Forty",
        5: "Fifty",
        6: "Sixty",
        7: "Seventy",
        8: "Eighty",
        9: "Ninety",
    }
    for i := 2 ; i < 10 ; i++ {
        if num >= i * 10 && num < (i+1) * 10 {
            switch num {
            case i * 10:
                return dict[i]
            default:
                return dict[i] + " " + numberToWords(num-(i*10))
            }
        }
    }
    return "Unknown value"
}

func translateTeen(num int) string {
    return map[int]string{
        11: "Eleven",
        12: "Twelve",
        13: "Thirteen",
        14: "Fourteen",
        15: "Fifteen",
        16: "Sixteen",
        17: "Seventeen",
        18: "Eighteen",
        19: "Nineteen",
    }[num]
}

func translateUpToTen(num int) string {
    return map[int]string{
        0: "Zero",
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
        10: "Ten",
    }[num]
}
