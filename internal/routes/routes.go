package routes

//import()

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	// &Router memanggil struct yang telah di buat
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter,
	}
}

//(r *Router) mengambil struct diatas
func (r *Router) SetRouters() *chi.mux {

}

func (r *Router) setConfigRouters() {

}

func RouterHealth() {}

func RouterProduct() {}

func EnableTimeout() {}

func EnableCORS() {}

func EnableRecover() {}

func EnableRequestId() {}

func EnableRealIp() {}
