# gomod

### `gomod tidy`

内置了 `GOPROXY="https://athens.azurefd.net"`，可以默认在 `gomod tidy` 时使用代理（没有设置 `GOPROXY` 的情况下），
代理相关见 `https://github.com/gomods/athens`
 
 > 推荐的方式还是直接设置 `export GOPROXY=https://athens.azurefd.net`，使用 `go mod tidy`
 
 ### `gomod graph`
 
 原版的 `go mod graph` 看起来如下
 
 ```bash
github.com/go-xorm/xorm cloud.google.com/go@v0.34.0
github.com/go-xorm/xorm github.com/cockroachdb/apd@v1.1.0
github.com/go-xorm/xorm github.com/davecgh/go-spew@v1.1.1
github.com/go-xorm/xorm github.com/denisenkom/go-mssqldb@v0.0.0-20190121005146-b04fd42d9952
github.com/go-xorm/xorm github.com/go-sql-driver/mysql@v1.4.1
github.com/go-xorm/xorm github.com/google/go-cmp@v0.2.0
github.com/go-xorm/xorm github.com/jackc/fake@v0.0.0-20150926172116-812a484cc733
github.com/go-xorm/xorm github.com/jackc/pgx@v3.3.0+incompatible
github.com/go-xorm/xorm github.com/kr/pretty@v0.1.0
github.com/go-xorm/xorm github.com/lib/pq@v1.0.0
github.com/go-xorm/xorm github.com/mattn/go-sqlite3@v1.10.0
github.com/go-xorm/xorm github.com/pkg/errors@v0.8.1
github.com/go-xorm/xorm github.com/satori/go.uuid@v1.2.0
github.com/go-xorm/xorm github.com/shopspring/decimal@v0.0.0-20180709203117-cd690d0c9e24
github.com/go-xorm/xorm github.com/stretchr/testify@v1.3.0
github.com/go-xorm/xorm github.com/ziutek/mymysql@v1.5.4
github.com/go-xorm/xorm golang.org/x/crypto@v0.0.0-20190122013713-64072686203f
github.com/go-xorm/xorm gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
github.com/go-xorm/xorm xorm.io/builder@v0.3.5
github.com/go-xorm/xorm xorm.io/core@v0.6.3
google.golang.org/appengine@v1.4.0 github.com/golang/protobuf@v1.2.0
github.com/stretchr/testify@v1.3.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.3.0 github.com/davecgh/go-spew@v1.1.0
github.com/kr/pretty@v0.1.0 github.com/kr/text@v0.1.0
google.golang.org/appengine@v1.4.0 golang.org/x/text@v0.3.0
google.golang.org/appengine@v1.4.0 golang.org/x/net@v0.0.0-20180724234803-3673e40ba225
github.com/kr/text@v0.1.0 github.com/kr/pty@v1.1.1
github.com/stretchr/testify@v1.3.0 github.com/stretchr/objx@v0.1.0
xorm.io/builder@v0.3.5 github.com/stretchr/testify@v1.3.0
xorm.io/builder@v0.3.5 github.com/go-xorm/sqlfiddle@v0.0.0-20180821085327-62ce714f951a
xorm.io/core@v0.6.3 google.golang.org/appengine@v1.4.0
xorm.io/core@v0.6.3 github.com/mattn/go-sqlite3@v1.10.0
xorm.io/core@v0.6.3 github.com/go-sql-driver/mysql@v1.4.1
```

使用 `gomod graph` 则显示为类似于 `tree` 的层级结构
```bash
github.com/go-xorm/xorm
├─ cloud.google.com/go@v0.34.0
├─ github.com/cockroachdb/apd@v1.1.0
├─ github.com/davecgh/go-spew@v1.1.1
├─ github.com/denisenkom/go-mssqldb@v0.0.0-20190121005146-b04fd42d9952
├─ github.com/go-sql-driver/mysql@v1.4.1
├─ github.com/google/go-cmp@v0.2.0
├─ github.com/jackc/fake@v0.0.0-20150926172116-812a484cc733
├─ github.com/jackc/pgx@v3.3.0+incompatible
├─ github.com/kr/pretty@v0.1.0
│  └─ github.com/kr/text@v0.1.0
│     └─ github.com/kr/pty@v1.1.1
├─ github.com/lib/pq@v1.0.0
├─ github.com/mattn/go-sqlite3@v1.10.0
├─ github.com/pkg/errors@v0.8.1
├─ github.com/satori/go.uuid@v1.2.0
├─ github.com/shopspring/decimal@v0.0.0-20180709203117-cd690d0c9e24
├─ github.com/stretchr/testify@v1.3.0
├─ github.com/ziutek/mymysql@v1.5.4
├─ golang.org/x/crypto@v0.0.0-20190122013713-64072686203f
├─ gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
├─ xorm.io/builder@v0.3.5
│  ├─ github.com/go-xorm/sqlfiddle@v0.0.0-20180821085327-62ce714f951a
│  └─ github.com/stretchr/testify@v1.3.0
│     ├─ github.com/davecgh/go-spew@v1.1.0
│     ├─ github.com/pmezard/go-difflib@v1.0.0
│     └─ github.com/stretchr/objx@v0.1.0
└─ xorm.io/core@v0.6.3
   ├─ github.com/go-sql-driver/mysql@v1.4.1
   ├─ github.com/mattn/go-sqlite3@v1.10.0
   └─ google.golang.org/appengine@v1.4.0
      ├─ github.com/golang/protobuf@v1.2.0
      ├─ golang.org/x/net@v0.0.0-20180724234803-3673e40ba225
      └─ golang.org/x/text@v0.3.0
```

也可以设置 `deep` 深度，`gomod graph 1`
```bash
github.com/go-xorm/xorm
├─ cloud.google.com/go@v0.34.0
├─ github.com/cockroachdb/apd@v1.1.0
├─ github.com/davecgh/go-spew@v1.1.1
├─ github.com/denisenkom/go-mssqldb@v0.0.0-20190121005146-b04fd42d9952
├─ github.com/go-sql-driver/mysql@v1.4.1
├─ github.com/google/go-cmp@v0.2.0
├─ github.com/jackc/fake@v0.0.0-20150926172116-812a484cc733
├─ github.com/jackc/pgx@v3.3.0+incompatible
├─ github.com/kr/pretty@v0.1.0
├─ github.com/lib/pq@v1.0.0
├─ github.com/mattn/go-sqlite3@v1.10.0
├─ github.com/pkg/errors@v0.8.1
├─ github.com/satori/go.uuid@v1.2.0
├─ github.com/shopspring/decimal@v0.0.0-20180709203117-cd690d0c9e24
├─ github.com/stretchr/testify@v1.3.0
├─ github.com/ziutek/mymysql@v1.5.4
├─ golang.org/x/crypto@v0.0.0-20190122013713-64072686203f
├─ gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
├─ xorm.io/builder@v0.3.5
└─ xorm.io/core@v0.6.3
```