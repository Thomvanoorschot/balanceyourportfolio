package api

import (
	"log"
	"net"

	"etfinsight/api/handlers"
	"etfinsight/config"
	"etfinsight/generated/proto"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/fund"
	"etfinsight/services/portfolio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ListenAndServe(host string, port string) error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	cfg := config.Load()

	repo := pgrepo.NewRepository(cfg)
	fundHandler := handlers.NewFundHandler(fund.NewService(repo))
	portfolioHandler := handlers.NewPortfolioHandler(portfolio.NewService(repo))

	proto.RegisterFundServiceServer(s, fundHandler)
	proto.RegisterPortfolioServiceServer(s, portfolioHandler)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
	return nil

	//return http.ListenAndServe(fmt.Sprintf("%s:%s", host, port))
}
