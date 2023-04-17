package acstringmatch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SimpleMatch(t *testing.T) {
	acSimpleMatch := NewAC([]string{"she", "his", "hers", "he"})

	//result := acSimpleMatch.Match("1237777")

	//require.Equal(t, true, result)

	result := acSimpleMatch.Match("ashersa")

	require.Equal(t, 2, result)

	result = acSimpleMatch.Match("ashers")
	require.Equal(t, 2, result)

	result = acSimpleMatch.Match("aaaaaahersbbbbbbb")
	require.Equal(t, 1, result)
}
