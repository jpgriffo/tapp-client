package data

type Policy struct {
	Rules []Rule 
}

type Rule struct {
	Protocol string `json:"ip_protocol"`
	Cidr string `json:"cidr_ip"`
	MaxPort int `json:"max_port"`
	MinPort int  `json:"min_port"`
}
