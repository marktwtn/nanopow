package nanopow

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGenerateWork(t *testing.T) {
	file, err := os.Open("hash")
	if err != nil {
		fmt.Println("error")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//hash, difficulty := make([]byte, 32), CalculateDifficulty(8)
		//rand.Read(hash)
		difficulty := CalculateDifficulty(8)
		hash := []byte(scanner.Text())
		fmt.Println("difficulty: ", difficulty)
		fmt.Println("hash: ", string(hash))

		// Start
		start := time.Now()
		w, err := GenerateWork(hash, difficulty)
		if err != nil {
			t.Error(err)
		}

		if IsValid(hash, difficulty, w) == false {
			t.Error("create invalid work")
		}
		// End
		elapsed := time.Since(start)
		log.Printf("Time: %s", elapsed)
	}

	return
}

func TestGenerateWork2(t *testing.T) {
	rand.Seed(42) // To make always test the same

	for n := 0; n < 1; n++ {
		hash, difficulty := make([]byte, 32), V1BaseDifficult
		rand.Read(hash)

		w, err := GenerateWork(hash, difficulty)
		if err != nil {
			t.Error(err)
		}

		if IsValid(hash, difficulty, w) == false {
			t.Error("create invalid work", w[:], hash)
		}
	}

	return
}
