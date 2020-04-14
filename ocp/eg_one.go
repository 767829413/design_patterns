package ocp

/*
//这是种没有遵守ocp的写法(不要死记硬背,灵活运用)
type Config struct {
}

func (this *Config) getMaxTps() int {
	return 0
}
func (this *Config) getMaxErrorCount() int {
	return 0
}

type AlertRule struct {
	config *Config
}

func (this *AlertRule) getMatchedRule(api string) *Config {
	return this.config
}

type Notification struct {
}

func (this *Notification) notify(lever string) {}

type Alert struct {
	rule         *AlertRule
	notification *Notification
}

func NewAlert(rule *AlertRule, notification *Notification) *Alert {
	alert := Alert{
		rule:         rule,
		notification: notification,
	}
	return &alert
}

func (this *Alert) Check(api string, requestCount int, errCount int, durationOfSeconds int) {
	tps := requestCount / durationOfSeconds
	if tps > this.rule.getMatchedRule(api).getMaxTps() {
		this.notification.notify("warning")
	}
	if errCount > this.rule.getMatchedRule(api).getMaxErrorCount() {
		this.notification.notify("error")
	}
}
*/
type Config struct {
}

func (this *Config) getMaxTps() int {
	return 0
}
func (this *Config) getMaxErrorCount() int {
	return 0
}

type AlertRule struct {
	config *Config
}

func (this *AlertRule) getMatchedRule(api string) *Config {
	return this.config
}

func NewAlertRule(config *Config) *AlertRule {
	return &AlertRule{
		config: config,
	}
}

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{
	}
}
func (this *Notification) notify(lever string) {}

type ApiStatInfo struct {
	Api               string
	RequestCount      int
	ErrCount          int
	DurationOfSeconds int
}

type AlertHandler interface {
	Check(apiStatInfo *ApiStatInfo)
}

type Alert struct {
	alertHandlers []AlertHandler
}

func NewAlert() *Alert {
	alert := Alert{}
	return &alert
}

func (this *Alert) addAlertHandler(alertHandler AlertHandler) {
	this.alertHandlers = append(this.alertHandlers, alertHandler)
}

func (this *Alert) AllCheck(apiStatInfo *ApiStatInfo) {
	for _, handler := range this.alertHandlers {
		handler.Check(apiStatInfo)
	}
}

type TpsAlertHandler struct {
	rule         *AlertRule
	notification *Notification
}

func NewTpsAlertHandler(rule *AlertRule, notification *Notification) *TpsAlertHandler {
	return &TpsAlertHandler{
		rule:         rule,
		notification: notification,
	}
}

func (this *TpsAlertHandler) Check(apiStatInfo *ApiStatInfo) {
	tps := apiStatInfo.RequestCount / apiStatInfo.DurationOfSeconds
	if (tps > this.rule.getMatchedRule(apiStatInfo.Api).getMaxTps()) {
		this.notification.notify("warning");
	}
}

type ErrorAlertHandler struct {
	rule         *AlertRule
	notification *Notification
}

func NewErrorAlertHandler(rule *AlertRule, notification *Notification) *ErrorAlertHandler {
	return &ErrorAlertHandler{
		rule:         rule,
		notification: notification,
	}
}

func (this *ErrorAlertHandler) Check(apiStatInfo *ApiStatInfo) {
	if apiStatInfo.ErrCount > this.rule.getMatchedRule(apiStatInfo.Api).getMaxErrorCount() {
		this.notification.notify("error")
	}
}

type ApplicationContext struct {
	AlertRule    *AlertRule
	Notification *Notification
	Alert        *Alert
}

func NewApplicationContext() *ApplicationContext {
	ac := &ApplicationContext{}
	ac.AlertRule = NewAlertRule(&Config{}) //省略一些初始化代码
	ac.Notification = NewNotification()    //省略一些初始化代码
	ac.Alert = NewAlert();
	ac.Alert.addAlertHandler(NewTpsAlertHandler(ac.AlertRule, ac.Notification))
	ac.Alert.addAlertHandler(NewErrorAlertHandler(ac.AlertRule, ac.Notification))
	return ac
}

func (this *ApplicationContext) GetAlert() *Alert {
	return this.Alert
}
