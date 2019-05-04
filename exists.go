package main

type anyList []int
func (list anyList) exists(f func(int) bool) bool {
    for _, l := range list {
        if f(l) {
            return true
        }
    }
    return false
}