package utils

func UInt64ToFloat64Ptr(i uint64) *float64 {
	f := float64(i)
	return &f
}

func AddFloat64Ptr(a, b *int64) *int64 {
	if a == nil || b == nil {
		return nil
	}
	sum := *a + *b
	return &sum
}
