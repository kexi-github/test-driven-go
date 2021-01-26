package Goroutines

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(checker WebsiteChecker,websites []string) map[string]bool {

	websiteschecked := make(map[string]bool)
	resultChannel := make(chan result)
	for _,value := range websites{
		go func(u string) {
			resultChannel <- result{u,checker(u)}
		}(value)
	}
	for i := 0;i < len(websites); i++{
		result := <- resultChannel
		websiteschecked[result.string] = result.bool
	}
	return websiteschecked
}
