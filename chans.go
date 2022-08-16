package util

func SendOrDiscard[T any](ch chan<- T, v T) {
	select {
	case ch <- v:
	default:
	}
}
