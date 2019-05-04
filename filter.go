package main

type anyList []int
func (list anyList) filter(f func(int) bool) anyList {
    var acc anyList
    for _, l := range list {
        if f(l) {
            acc = append(acc, l)
        }
    }
    return acc
}