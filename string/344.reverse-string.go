package string


//递归
func reverseString(s []byte)  {
    helper(s, 0, len(s) - 1)
}
func helper(*s []byte, left int, right int) {
    if left < right {
        tmp := s[left]
        left ++
        s[left] = s[right]
        right--
        s[right] = tmp
        helper(s, left, right)
    }
    
}
//双指针
func reverseString2(s []byte)  {
    left:=0
    right:=len(s)-1
    for left<right {
        s[left],s[right] = s[right],s[left]
        left++
        right--
    }
}
