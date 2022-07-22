package main

func main() {
	server := NewServer(":8081")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/generic", server.AddMiddleware(GetZonesSummary, CheckAuth()))
	server.Handle("GET", "/zone/:id", server.AddMiddleware(HandleRootZone, CheckAuth()))
	server.Handle("POST", "/sensor", server.AddMiddleware(ZoneDataRequest, CheckAuth()))
	server.Handle("GET", "/cars/:id", server.AddMiddleware(CarGetRequest, CheckAuth(), Loggin()))
	server.Handle("DELETE", "/cars/:id", server.AddMiddleware(CarDeleteRequest, CheckAuth(), Loggin()))
	server.Listen()
}
