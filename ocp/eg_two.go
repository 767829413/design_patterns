package ocp

type MessageFormatter interface {
	Format()
}

type JsonMessageFormatter struct {
}

func (this *JsonMessageFormatter) Format() {

}

type ProtobufFormatter struct {
}

func (this *ProtobufFormatter) Format() {

}

type MessageQueue interface {
	Send()
}

type KafkaMessageQueue struct {
}

func (this *KafkaMessageQueue) Send() {

}

type RocketMQMessageQueue struct {
}

func (this *RocketMQMessageQueue) Send() {

}

type Demo struct {
	queue MessageQueue
}

func (this *Demo) Notify(mf MessageFormatter, notification *Notification) {

}

func NewDemo(queue MessageQueue) *Demo {
	return &Demo{
		queue: queue,
	}
}
