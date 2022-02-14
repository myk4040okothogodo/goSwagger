package main
import (
  "time"
  "github.com/gorilla/mux"
  "context"
  "os"
  "os/signal"
  "net/http"
  "log"
  "github.com/go-openapi/runtime/middleware"
  "github.com/myk4040okothogodo/GoMicroserve/handlers"
  "github.com/myk4040okothogodo/GoMicroserve/data"
)

func main() {

    
    l := log.New(os.Stdout, "products-api ", log.LstdFlags)
    v := data.NewValidation()
	// create the handlers
	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", ph.ListAll)
  getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareValidateProduct)


  deleteR := sm.Methods(http.MethodDelete).Subrouter()
  deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

  //handler for documentation
  opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
  sh := middleware.Redoc(opts, nil)

  getR.Handle("/docs", sh)
  getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
    // create a new server 
    s := &http.Server{
        Addr: ":9090",
        Handler: sm,
        ErrorLog: l,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    
    }
    go func(){
      l.Println("\n Starting server on port 9090: \n")
        
        err := s.ListenAndServe()
        if err != nil {
            l.Printf("Error starting server: %s\n", err)
            os.Exit(1)
        }
    }()
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)
    
    sig := <- c
    l.Println("Received signal", sig)
    
    //gracefully shutdown the server, wiating max 30 seconds for current operations to finish
    tc, _  := context.WithTimeout(context.Background(), 30*time.Second)
    s.Shutdown(tc)
    

}
