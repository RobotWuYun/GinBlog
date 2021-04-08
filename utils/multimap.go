package utils

type Multimap map[string][]string

func (multimap Multimap) Add(key, value string) {
	if len(multimap[key]) == 0 {
		multimap[key] = []string{value}
	} else {
		multimap[key] = append(multimap[key], value)
	}
}

func (multimap Multimap) Get(key string) []string {
	if multimap == nil {
		return nil
	}
	values := multimap[key]
	return values
}
