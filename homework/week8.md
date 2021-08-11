# 第八周作业

## [转换成小写字母](https://leetcode-cn.com/problems/to-lower-case/)

```go
func toLowerCase(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			b[i] = s[i] + 32
		}
	}
	return string(b)
}
```

## [翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

```go
func reverseWords(s string) string {
    if len(s) == 0 {
        return s
    }

    // 获取的字符串放入栈中，从栈中弹出的时候就是翻转的字符串了
    stack := make([]string, 0)
    char := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            char = append(char, s[i])
        } else {
            if len(char) > 0 {
                stack = append(stack, string(char))
            }
            char = []byte{}
        }
    }

    if len(char) > 0 {
        stack = append(stack, string(char))
    }

    res := []byte{}
    for i := len(stack)-1; i >= 0; i-- {
        res = append(res, []byte(stack[i])...)
        res = append(res, ' ')
    }

    // 翻转后在最后会多个空格，因此需要截取
    return string(res[:len(res)-1])
}
```
