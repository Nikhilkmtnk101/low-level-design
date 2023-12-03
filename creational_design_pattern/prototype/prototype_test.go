package prototype

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetShirtCloner(t *testing.T) {
	assert.NotNil(t, GetShirtCloner())
}

func TestShirtCacheCloner_GetClone(t *testing.T) {
	shirtCloner := GetShirtCloner()
	item1, err := shirtCloner.GetClone(White)

	if item1 == whiteShirtCache {
		t.Errorf("shirt object should be different from whiteShirtCache object")
	}
	assert.Nil(t, err, fmt.Sprintf("error should be nil"))

	shirt1, ok := item1.(*Shirt)
	assert.Equal(t, ok, true, fmt.Sprintf("Type assertion for item1 canot be done successfully"))
	assert.Equal(t, int(shirt1.Color), White)
	assert.Equal(t, shirt1.Price, float64(15))
	assert.Equal(t, shirt1.SKU, "")
	shirt1.SKU = "xyz"

	item2, err := shirtCloner.GetClone(shirt1.Color)
	if item1 == item2 {
		t.Errorf("shirt object should be different from whiteShirtCache object")
	}
	shirt2, ok := item2.(*Shirt)
	assert.Equal(t, ok, true, fmt.Sprintf("Type assertion for item2 canot be done successfully"))
	assert.NotEqual(t, shirt1.SKU, shirt2.SKU)
}
