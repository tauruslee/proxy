package service

import(
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"taurus.com/proxy/dbclient"
	"fmt"
	"net"
)

var(
	DBClient dbclient.IBoltClient
)


