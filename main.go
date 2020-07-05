package vrsgim

import (
	"os"

	c3mcommon "github.com/tidusant/chadmin-common"

	"github.com/tidusant/chadmin-log"

	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init repo vrsgim")
	strErr := ""
	db, strErr = c3mcommon.ConnectDB("vrsgim")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
}
