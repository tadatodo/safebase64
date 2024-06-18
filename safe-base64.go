package safebase64

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
)

type Base64 struct {
	BlockList []*regexp.Regexp
	Rand      *rand.Rand
}

// New creates a new Base64 instance with the given blockList.
// The blockList is a case-insensitive list of words that should not be generated and
// will be converted to regexes with common vowel<>number swaps, e.g. a<>4 e<>3.
func New(blockList []string) *Base64 {
	return NewWithRand(blockList, rand.New(rand.NewSource(time.Now().UnixNano())))
}

// New creates a new Base64 instance with the given blockList.
// The blockList is a case-insensitive list of words that should not be generated and
// will be converted to regexes with common vowel<>number swaps, e.g. a<>4 e<>3.
// The random number generator r is used to generate the random strings.
func NewWithRand(blockList []string, r *rand.Rand) *Base64 {
	if r == nil {
		panic("rand is nil")
	}
	return &Base64{
		BlockList: generateRegex(blockList),
		Rand:      r,
	}
}

// take the list of words and generate regexes where numbers could be inserted, ensure case insensitivity, 1=i, 3=e, 4=a, 5=s, 0=o anywhere in the string
func generateRegex(blockList []string) []*regexp.Regexp {
	var regexList []*regexp.Regexp
	i1Regex := regexp.MustCompile("1|i|l")
	e3Regex := regexp.MustCompile("3|e")
	a4Regex := regexp.MustCompile("4|a")
	s5Regex := regexp.MustCompile("5|s")
	o0Regex := regexp.MustCompile("0|o")

	for _, word := range blockList {
		// replace 1 with i, 3 with e, 4 with a, 5 with s, 0 with o
		word = i1Regex.ReplaceAllString(word, "[1il]")
		word = e3Regex.ReplaceAllString(word, "[3e]")
		word = a4Regex.ReplaceAllString(word, "[4a]")
		word = s5Regex.ReplaceAllString(word, "[5s]")
		word = o0Regex.ReplaceAllString(word, "[0o]")
		// case insensitive
		word = "(?i)" + word
		fmt.Println("generated regex: ", word)
		regexList = append(regexList, regexp.MustCompile(word))
	}
	return regexList
}

// Generate generates a random string of length n using the charset.
// It avoids generating strings containing words in the block list and will try 100 new strings before panicing.
func (ba *Base64) Generate(n int) string {
	if n <= 0 {
		return ""
	}
	for i := 0; i < 100; i++ {
		// first character should start with a letter
		// this is to avoid issues with URL encoding
		firstCharIndex := ba.Rand.Intn(len(letters))
		str := string(letters[firstCharIndex])
		for i := 1; i < n; i++ {
			str += string(charset[ba.Rand.Intn(len(charset))])
		}
		if !ba.ContainsSwearWord(str) {
			return str
		}
	}
	panic("could not generate a safe string")
}

// ContainsSwearWord checks if the generated string contains any swear words.
func (ba *Base64) ContainsSwearWord(str string) bool {
	for _, regex := range ba.BlockList {
		if regex.MatchString(str) {
			return true
		}
	}
	return false
}
