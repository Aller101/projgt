package singleton

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSingletonGorut(t *testing.T) {
	t.Parallel()
	// s1 := NewSingleton("PETR")
	// fmt.Println("===> ", s1.GetName())
	// s2 := NewSingleton("OLIA")
	// fmt.Println("===> ", s2.GetName())
	// s3 := NewSingleton("SANIYA")
	// fmt.Println("===> ", s3.GetName())

	wg := sync.WaitGroup{}
	// chT := make(chan struct{})

	for i := range 1000 {

		// t.Run(strconv.Itoa(i), func(t *testing.T) {
		// 	t.Parallel()
		// 	s := NewSingleton(fmt.Sprintf("x%d", i))
		// 	fmt.Println("===> ", s.GetName())
		// })

		wg.Add(1)

		go func() {

			defer wg.Done()
			s := NewSingleton(fmt.Sprintf("x%d", i))
			fmt.Println("===> ", s.GetName())
		}()
	}
	wg.Wait()
}

func TestSingletonParall(t *testing.T) {

	for i := range 1000 {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			s := NewSingleton(fmt.Sprintf("x%d", i))
			fmt.Println("===> ", s.GetName())
		})

	}

}
