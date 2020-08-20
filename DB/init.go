package DB

import (
	"TestGOWeb/config"
	"TestGOWeb/recordlog"
	"database/sql" //GO语言自带的数据库依赖
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" //使用_代表执行这个包的init函数 用于初始化函数  sqlServer数据库驱动包
)

//创建数据库连接信息对象
type Deinfo struct {
	IP     string
	Port   string
	User   string
	Pwd    string
	DbName string
}

type DBLStruct struct {
	DB *sql.DB
}

//声明数据参数载体
var (
	DB     *sql.DB //声明数据连接对象
	DbPort string
	DbHost string
	DbUser string
	DbName string
	DbPwd  string
	DbPool int                   //数据库连接池
	DBLink map[string]*DBLStruct = map[string]*DBLStruct{}
)

//init方法 自动执行
func init() {

	dbinit()
}

func dbinit() error {
	var err error
	recordlog.Debug("============开始初始话数据库=================")
	conf_String := func(section string, option string) (string, error) {
		if err != nil {
			return "", err
		}
		return conf.Conf.String(section, option) //获取配置文件信息
	}
	DbHost, err = conf_String("DB", "DbHost")
	DbPort, err = conf_String("DB", "DbPort")
	DbName, err = conf_String("DB", "DbName")
	DbUser, err = conf_String("DB", "DbUser")
	DbPwd, err = conf_String("DB", "DbPassword")
	recordlog.Debug("数据库连接信息：", DbHost, ":", DbPort)
	DbPool, err = conf.Conf.Int("DB", "DbPool") //获取数据连接池容量

	if err != nil { //错误本身没有意义
		recordlog.Debug("Initialize DB config error:", err)
	}

	dbUrl := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", DbHost, DbPort, DbName, DbUser, DbPwd) //拼接数据库连接字符串

	recordlog.Debug("===================", dbUrl, "======================")
	DB, err = sql.Open("mssql", dbUrl) //驱动名称,数据源连接参数  获取数据库对象
	if err != nil {
		recordlog.Debug("创建数据库连接对象报错")
		panic(err)
	}
	DB.SetMaxOpenConns(DbPool) /*设置连接池容量*/
	err = DB.Ping()
	if err != nil {
		recordlog.Debug("ping DBserver error:", err)
	}
	/**/
	recordlog.Debug("Initialize DB connect OK...")
	return nil
}
