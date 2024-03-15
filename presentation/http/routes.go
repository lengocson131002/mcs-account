package http

func WithAccountV1Routes() HttpServerStartOption {
	return func(s *HttpServer) error {
		v1 := s.App.Group("/api/v1")
		v1.Get("/accounts/check-balance", s.AccountController.GetAccountBalance)
		return nil
	}
}
