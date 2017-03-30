package main

type ISettings interface {
	parseSettings()
}

type Settings struct {
	WebPort    string   `json:"web_port"`
	MOTD       string   `json:"motd"`
	AlertEmail string   `json:"alert_email"`
	Carafes    []Carafe `json:"carafes"`
}

func (s *Settings) parseSettings {
		
}


