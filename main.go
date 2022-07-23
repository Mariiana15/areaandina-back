package main

func main() {
	server := NewServer(":8080")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/generic", server.AddMiddleware(GetZonesSummary, CheckAuth()))
	server.Handle("GET", "/zone/:id", server.AddMiddleware(HandleRootZone, CheckAuth()))
	server.Handle("POST", "/sensor", server.AddMiddleware(ZoneDataRequest, CheckAuth()))
	server.Listen()
}
