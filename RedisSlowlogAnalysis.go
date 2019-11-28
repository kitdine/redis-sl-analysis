package redis_sl_analysis

type RedisSlowlog struct {
	Index       int64
	Cmd         string
	Key         []string
	Args        []string
	ExecuteTime int64
	ExecuteCost int64
}

type CommandInfo struct {
	Cmd string

}

type RedisSlowlogAnalyst interface {
	FilterKeys(log *RedisSlowlog)

}

type redisSlowlogAnalysis struct {
	event RedisSlowlogAnalyst
	logs  <-chan *RedisSlowlog
}

func Analysis(log <-chan *RedisSlowlog, analyst RedisSlowlogAnalyst) error {
	a :=&redisSlowlogAnalysis{analyst, log}
	return a.analysis()
}

func (a *redisSlowlogAnalysis) analysis() error {
	for log := range a.logs {
		a.event.FilterKeys(log)
	}

	return nil
}