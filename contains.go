package main

type anyList []int
func (list anyList) contains(target int) bool {
    for _, n := range list {
        if n == target {
            return true
        }
    }
    return false
}