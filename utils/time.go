package utils

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) {
	fmt.Println("TimeTrack", time.Since(start))
}
