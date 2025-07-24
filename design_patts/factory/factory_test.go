package factory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRocketFactory(t *testing.T) {

	usa := NewRocketFactory("usa")
	require.Equal(t, "usa", usa.GetName())

	ru := NewRocketFactory("ru")
	require.Equal(t, "ru", ru.GetName())

	ch := NewRocketFactory("ch")
	require.Equal(t, "ch", ch.GetName())

	unknown := NewRocketFactory("")
	require.Nil(t, unknown)

}
