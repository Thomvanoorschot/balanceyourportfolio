package http

//func ListenAndServe(addr string) error {
//	e := gin.Default()
//	cfg := config.Load()
//
//	repo := pgrepo.NewRepository(cfg)
//	h := handler.NewHandler(fund.NewService(repo), portfolio.NewService(repo), nil)
//
//	router := routes.NewRouter()
//	err := router.SetupRouter(e, cfg, h)
//	if err != nil {
//		return err
//	}
//
//	return e.Run(addr)
//}
