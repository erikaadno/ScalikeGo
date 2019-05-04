package main

type anyList []int
func (list anyList) mapMethod(f func(int) int) anyList {
    var acc anyList
    for _, l := range list {
        acc = append(acc, f(l))
    }
    return acc
}