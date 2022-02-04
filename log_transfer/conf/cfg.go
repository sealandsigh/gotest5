package conf

type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESCfg struct {
	Address string `ini:"address"`
}
