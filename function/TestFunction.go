package function

import (
	"TestGOWeb/DB"
	"TestGOWeb/function/mapFunc"
	"TestGOWeb/function/reflectFunc"
	"TestGOWeb/function/resDataFunc"
	"TestGOWeb/recordlog"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
)

type TestFunc struct{}

//创建结构体的指针对象 用于调用下面的方法
var TestFuncObj = new(TestFunc)
var lock sync.RWMutex

/*
 功能那个方法具有固定的格式 func (参数名称 结构体的指针类型) 方法名称(w http.ResponseWriter, r *http.Request, _ httprouter.Params)   参数是固定的
*/

func (this *TestFunc) GetMessage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	addMap(w, r, p)
}

func DBinsert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("============打印数据=============")
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)
	id := queryForm.Get("id")
	name := queryForm.Get("name")
	sql := fmt.Sprintf("INSERT INTO class (id,name) VALUES(%s,'%s');", id, name)
	// CREATE_TABLE := "create table class (id char(4) primary key ,name varchar(50)not null)"
	DB.ExeSQLIU(sql)
	recordlog.Debug("数据库操作执行完成")
	resData.Write(w, "数据异常", "parameter is null")
}

func addMap(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lock.RLock()

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println("panic: ", err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	panic("直接报错")
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)
	id := queryForm.Get("id")
	name := queryForm.Get("name")

	mapFunc.DataMap[id] = name
	fmt.Println("map数量：", len(mapFunc.DataMap))
	lock.RUnlock()
}

func TestXieChen(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	consumerAndProducerFunc()
	// cacheChannal()
}

func TestReflectFunc(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reflectFunc.TestReflectFunc()
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func consumerAndProducerFunc() {

	ch := make(chan int)

	go produce(ch)
	go consumer(ch)
}

func cacheChannal() {
	ch := make(chan int)
	recordlog.Debug("==============================管道存值之前======================")
	ch <- 1 //先管道中加入值
	recordlog.Debug("==============================管道存值之后======================")

	go func() {
		// var x int
		recordlog.Debug("==============================管道取值之前======================")
		x := <-ch

		recordlog.Debug("==============================管道取值之后======================")
		fmt.Println("线程内部", x)
	}()
	go func() {
		// var x int
		recordlog.Debug("==============================管道取值之前======================")
		x := <-ch
		recordlog.Debug("==============================管道取值之后======================")
		fmt.Println("线程内部", x)
		fmt.Println("2")
	}()

	time.Sleep(time.Second * 1)
}

func init() {
	// go chackMap()
}

func chackMap() {
	for {
		var num = len(mapFunc.DataMap)
		if num == 5 {
			lock.Lock()
			fmt.Println("======================================================================")
			mapFunc.DataMap = make(map[string]string)
			lock.Unlock()
		}
	}

}
