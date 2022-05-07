package main

import (
	"database/sql"
	"flag"
	"log"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	gen "github.com/wxxhub/gen_sqlpb/internal"
	idb "github.com/wxxhub/gen_sqlpb/internal/db"
)

type TableConfig struct {
	mysqldsn       string
	tableName string
}

type Config struct {
	tableConfigs []TableConfig
	serviceName  string
	savePath     string
}

func parseTableConfig(dsn) {

}

func parseFlag() *Config {
	config := &Config{
		tableConfigs: make([]TableConfig, 0),
	}
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		switch flag.Arg(i) {
		case "dsn":

		}
	}

	return config
}

func main() {
	dsn := flag.String("dsn", "", "the database dsn")
	serviceName := flag.String("srv", "", "the protobuf service name , defaults to the database schema.")
	savePath := flag.String("save_path", "", "the protobuf service name , defaults to the database schema.")
	tableName := flag.String("table_name", "", "the protobuf service name , defaults to the database schema.")

	//packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	//goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	cols, err := idb.GenerateSchema(db, *tableName)
	if err != nil {
		log.Panic(err)
		//panic(err)
	}

	path := *serviceName + ".proto"
	if len(*savePath) > 0 {
		path = filepath.Join(*savePath, path)
	}

	gen.GenProto(cols, *serviceName, *tableName, path)
}