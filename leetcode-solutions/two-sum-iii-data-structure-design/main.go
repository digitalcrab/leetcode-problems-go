package main

type TwoSum struct {
	data map[int]int
}

func Constructor() TwoSum {
	return TwoSum{data: make(map[int]int)}
}

func (this *TwoSum) Add(number int) {
	this.data[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num, cnt := range this.data {
		reminder := value - num
		if reminder == num {
			if cnt > 1 {
				return true
			}
		} else if this.data[reminder] > 0 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

func main() {

}
