package main

type stack []byte

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(ele byte) {
	(*s) = append((*s), ele)
}

func (s *stack) Pop() byte {
	var ele byte

	if len(*s) > 0 {
		ele = (*s)[len(*s) - 1]
		*s = (*s)[:len(*s) - 1]
	}

	return ele
}

func isValid(s string) bool {
	st := &stack{}

	for idx := 0; idx < len(s); idx++ {
		currCh := s[idx]
		if (currCh == '(') || (currCh == '{') || (currCh == '[') {
			st.Push(currCh)
		} else if (currCh == ')') || (currCh == '}') || (currCh == ']') {
			lastParenthesis := st.Pop()

			switch currCh {
			case '}':
				if lastParenthesis != '{' {
					return false
				}
			case ']':
				if lastParenthesis != '[' {
					return false
				}
			case ')':
				if lastParenthesis != '(' {
					return false
				}
			}
		}
	}

	return st.IsEmpty()
}
