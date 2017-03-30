package main

type ICarafe interface {
	Monitor()
}

type Carafe struct {
	ID            int     `json:"id"`
	EmptyWeight   float32 `json:"empty_weight"`
	FullWeight    float32 `json:"full_weight"`
	CheckInterval int     `json:check_interval` // time in seconds between polls of the associated sensor
	GPIO          int     `json:"gpio"`
}

func startCarafeMonitoring() {

}
