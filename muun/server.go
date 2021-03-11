package muun

import (
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedStockServiceServer
	// rates map[string]float64
}

func (s *server) GetStockPrice(ctx context.Context, in *GetStockPriceRequest) (*GetStockPriceResponse, error) {
	return &GetStockPriceResponse{Strike: 53.00}, nil
}

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

func StartServer() {

	// Use this for debug related matters
	// http.HandleFunc("/headers", headers)
	// http.ListenAndServe(":8090", nil)

	// gRPC serve
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStockServiceServer(s, &server{ /*rates: make(map[string]float64, 0)*/ })
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
