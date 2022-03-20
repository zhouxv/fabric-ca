module github.com/hyperledger/fabric-ca

go 1.15

replace github.com/hyperledger/fabric v1.4.11 => github.com/zhouxv/fabric v1.4.13-0.20220319163203-19d166de80dc

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/cloudflare/cfssl v1.4.1
	github.com/felixge/httpsnoop v1.0.1
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grantae/certinfo v0.0.0-20170412194111-59d56a35515b
	github.com/hyperledger/fabric v1.4.11
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kisielk/sqlstruct v0.0.0-20201105191214-5f3e10d3ab46
	github.com/lib/pq v1.8.0
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/mitchellh/mapstructure v1.3.3
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.5.0
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.7.1-0.20210116013205-6990a05d54c2
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ldap.v2 v2.5.1
)
