package main

func main() {
	SETTINGS.load()
	go startCarafeMonitoring()
	go startWebApp()
	WG.Add(2)
	WG.Wait()
}
