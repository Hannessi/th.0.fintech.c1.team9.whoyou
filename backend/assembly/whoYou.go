package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/config"
	"github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/cors"
	userRpc "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/adaptor/rpc"
	userRecordkeeperMongoImpl "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/impl/recordkeeper/mongo"
	userRegistrarDefaultImpl "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/impl/registrar/default"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Starting WhoYou? v0.1.5")

	configFile := flag.String("config-file", "./config.json", "Config File Path")
	//privateKeyBits := flag.Int("private-key-bits", 4096, "The size in bits of the RSA private key")
	//keyStoreDirectory := flag.String("private-key-location", "./", "The /path/to/privatekey")
	enableTls := flag.String("mongo-tls", "enabled", "Enable/Disable Mongo TLS connection")
	flag.Parse()

	configuration, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("Failed to load config")
	}

	mongoServers := configuration.MongoServer

	log.Info("Connecting to Mongo DB...")

	tlsConfig := &tls.Config{}

	dialInfo := mgo.DialInfo{
		Addrs:    mongoServers,
		Timeout:  15 * time.Second,
		Username: configuration.MongoUsername,
		Password: configuration.MongoPassword,
	}

	if *enableTls == "enabled" {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	mainMongoSession, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		log.Fatal("Could not connect to Mongo DB: ", mongoServers, err)
	}

	log.Info("Connected to Mongo DB")

	jsonCodec := json2.NewCodec()
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(cors.CodecWithCors([]string{"*"}, jsonCodec), "application/json")
	rpcServer.RegisterCodec(jsonCodec, "application/json;charset=UTF-8")

	//
	userRecordkeeper := userRecordkeeperMongoImpl.Recordkeeper{
		MongoSession: mainMongoSession,
		Database:     configuration.DatabaseName,
		Collection:   "user",
	}

	userRegistrar := userRegistrarDefaultImpl.Registrar{
		Recordkeeper: &userRecordkeeper,
	}

	// instantiate all adaptors
	userRpcAdaptor := userRpc.Adaptor{
		Registrar: &userRegistrar,
	}

	// register rpc adaptors
	if err := rpcServer.RegisterService(&userRpcAdaptor, "User"); err != nil {
		log.Fatal(err.Error())
	}

	// create new router
	router := mux.NewRouter()
	router.Methods("OPTIONS").HandlerFunc(preFlightHandler)
	//router.Handle("/api", Authorize(logger(rpcServer)))
	router.Handle("/api", rpcServer)

	go func() {
		if err := http.ListenAndServe(":9002", router); err != nil {
			log.Error("Http server stopped: ", err.Error())
		}
	}()
	log.Info(fmt.Sprintf("HTTP server started and listening on port %s", "9002"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Info("WhoYou? is shutting down..")
			return
		}
	}
}

func preFlightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Origin, Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
