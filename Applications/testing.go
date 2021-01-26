package Applications

type StubPlayserStore struct {
	Scores map[string]int
}

type Player struct {
	Name string
	Wins int
}

func (s *StubPlayserStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayserStore) RecordWin(name string) {
	value,ok := s.Scores[name]
	if !ok {
		s.Scores[name] = 1
		return
	}
	s.Scores[name] = value + 1
}

func (s *StubPlayserStore) GetLeague() []Player{
	var league []Player
	for name,wins := range s.Scores{
		league = append(league,Player{name,wins})
	}
	return league
}
