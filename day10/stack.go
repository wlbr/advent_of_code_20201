package main

type stack struct {
	content string
}

func (s *stack) push(c rune) {
	s.content += string(c)
}

func (s *stack) peek() byte {
	l := len(s.content)
	if l > 0 {
		return s.content[l-1]
	}
	return ' '
}

func (s *stack) pop() byte {
	l := len(s.content)
	if l > 0 {
		last := s.content[l-1]
		s.content = s.content[:len(s.content)-1]
		return last
	}
	return ' '
}

func (s *stack) empty() bool {
	return len(s.content) == 0
}

func (s *stack) len() int {
	return len(s.content)
}

func (s *stack) reset() {
	s.content = ""
}
