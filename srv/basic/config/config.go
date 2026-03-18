package config

// AppConfig 应用配置
type AppConfig struct {
	Nacos struct {
		Addr      string
		Port      int
		Namespace string
		DataID    string
		Group     string
	}
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
	}

	Consul struct {
		Host        string
		Port        int
		ServiceName string
		ServicePort int
		TTL         int
	}

	AliPay struct {
		PrivateKey      string
		AlipayPublicKey string
		AppId           string
		NotifyURL       string
		ReturnURL       string
	}
}
