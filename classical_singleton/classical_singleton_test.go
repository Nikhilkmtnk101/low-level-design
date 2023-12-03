package classical_singleton

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingletonInstance(t *testing.T) {
	counter1 := GetSingletonInstance()
	assert.NotNil(t, counter1, fmt.Sprintf("Got counter after calling GetSingletonInstance is nil"))

	expectedCounter := counter1

	counter2 := GetSingletonInstance()
	assert.Equal(t, expectedCounter, counter2)

	currentCount1 := counter1.AddOne()

	assert.Equal(t, 1, currentCount1,
		fmt.Sprintf("After calling AddOne for first time it should be 1 but it is %d", currentCount1))

	currentCount2 := counter2.AddOne()

	assert.Equal(t, 2, currentCount2,
		fmt.Sprintf("After calling AddOne for second time it should be 2 but it is %d", currentCount2))

}
