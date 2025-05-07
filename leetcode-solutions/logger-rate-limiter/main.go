package main

type Logger struct {
	buffer map[string]int
}

func Constructor() Logger {
	return Logger{
		buffer: make(map[string]int),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	lastPrinted, exist := this.buffer[message]

	// was not printed yet
	if exist && timestamp-lastPrinted < 10 {
		return false
	}

	this.buffer[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {

}
