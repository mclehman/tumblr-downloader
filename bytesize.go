package main

import (
	"fmt"
)

const (
	_         = iota // ignore first value by assigning to blank identifier
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func byteSize(b uint64) string {
	switch {
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
