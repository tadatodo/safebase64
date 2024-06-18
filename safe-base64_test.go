package safebase64

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateBadWords(t *testing.T) {
	producesBadWord := NewWithRand([]string{}, rand.New(rand.NewSource(1)))
	badWord := "FTZidia3NjuA"

	found := false
	for i := 0; i < 10; i++ {
		got := producesBadWord.Generate(12)
		if got == badWord {
			found = true
			break
		}
	}
	require.True(t, found)

	doesntProduceBadWord := NewWithRand([]string{"diaen"}, rand.New(rand.NewSource(1)))

	for i := 0; i < 10; i++ {
		got := doesntProduceBadWord.Generate(12)
		if got == badWord {
			t.Errorf("Generate() = %v, never want %v", got, badWord)
		}
	}
}

func TestGenerate(t *testing.T) {
	b := New([]string{})
	require.Equal(t, 0, len(b.Generate(0)))
	require.Equal(t, 0, len(b.Generate(-1)))
	require.Equal(t, 12, len(b.Generate(12)))
	require.Equal(t, 130, len(b.Generate(130)))
}

func TestGeneratEdges(t *testing.T) {
	b := NewWithRand(nil, rand.New(rand.NewSource(1)))
	require.Equal(t, 10, len(b.Generate(10)))
}

func TestPanics(t *testing.T) {
	require.Panics(t, func() {
		NewWithRand(nil, nil)
	})

	// split alphabet into array
	alphabet := make([]string, len(charset))
	for i, c := range charset {
		alphabet[i] = string(c)
	}

	require.Panics(t, func() {
		New(alphabet).Generate(10)
	})
}
