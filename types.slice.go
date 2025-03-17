package core

// AnySlice 转为any分片
func AnySlice[V any](vs []V) []any {
	rt := make([]any, len(vs))
	for i, v := range vs {
		rt[i] = v
	}
	return rt
}

// AnyMap 转为any映射
func AnyMap[V any](vs map[string]V) map[string]any {
	rt := make(map[string]any, len(vs))
	for k, v := range vs {
		rt[k] = v
	}
	return rt
}
