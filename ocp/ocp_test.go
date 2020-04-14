package ocp

import "testing"

func TestDemo_noe(t *testing.T) {
	apiStatInfo := &ApiStatInfo{
		Api:               "test/hi",
		RequestCount:      10,
		ErrCount:          20,
		DurationOfSeconds: 25,
	}
	NewApplicationContext().GetAlert().AllCheck(apiStatInfo)
}

func TestDemo_two(t *testing.T) {
	NewDemo(&KafkaMessageQueue{}).Notify(&JsonMessageFormatter{}, &Notification{})
}
