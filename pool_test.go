package workerpool_test

import (
	"fmt"
	"testing"

	"github.com/dmora/workerpool"
	"github.com/stretchr/testify/assert"
    )

func TestNewPool(t *testing.T) {
	pool := workerpool.NewPool(2)
	assert.False(t, pool.IsCompleted())
}
func ResourceProcessor(resource interface{}) error {
	fmt.Printf("Resource processor got: %s", resource)
	fmt.Println()
	return nil
}

func ResultProcessor(result workerpool.Result) error {
	fmt.Printf("Result processor got: %s", result.Err)
	fmt.Println()
	return nil
}

func TestPool_Start(t *testing.T) {
	strings := []string{"first", "second"}
	resources := make([]interface{}, len(strings))
	for i, s := range strings {
		resources[i] = s
	}

	pool := workerpool.NewPool(3)
	pool.Start(resources, ResourceProcessor, ResultProcessor)
}
