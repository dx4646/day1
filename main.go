package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// post github.com/gin-gonic/gin

func main() {
	// group := sync.WaitGroup{}
	fmt.Println("start main")
	c := colly.NewCollector()
	count := 0
	// url := "https://imgur.com"
	// c.OnHTML("body", func(e *colly.HTMLElement) {
	// 	fmt.Println("目標值 text :", e.Text)
	// 	fmt.Println("目標值 attr :", e.Attr("src"))
	// 	// GetImg(e.Attr("src"))
	// 	count++
	// })
	url := "https://zh.pngtree.com/free-vectors?source_id=310&chnl=ggas&srid=764883278&gpid=57293140505&asid=423823300736&ntwk=g&tgkw=kwd-286971347&mchk=%E5%9C%96%E5%BA%AB&mcht=b&pylc=9040379&dvic=c&gclid=Cj0KCQjwkt6aBhDKARIsAAyeLJ1bU9NabQ7KjghGeeBGvWAZDHUiFyuIp6Jigh174AiOGl9AiEgUeHoaAraCEALw_wcB"
	// 要爬蟲的相關 KEY 值 //此為 LY平台登入頁面 的title
	c.OnHTML(".mb-box li img[src]", func(e *colly.HTMLElement) {
		fmt.Println("目標值 text :", e.Text)
		fmt.Println("目標值 attr :", e.Attr("src"))
		GetImg(e.Attr("src"))
		count++
	})

	c.OnRequest(func(r *colly.Request) {
		// r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	//收到返回資訊
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("收到响应后调用:", response.Request.URL)
	})

	// 此為 爬蟲目標網站 網站
	err := c.Visit(url)
	if err != nil {
		fmt.Println("具体错误:", err)
	}
	fmt.Println("查詢到項目 :", count)
}

// 爬蟲相關 套件 : go get -u github.com/gocolly/colly/v2

// // 下载图片信息
// func downLoad(base string, url string) error {
// 	pic := base
// 	idx := strings.LastIndex(url, "/")
// 	if idx < 0 {
// 		pic += "/" + url
// 	} else {
// 		pic += url[idx+1:]
// 	}
// 	v, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Http get [%v] failed! %v", url, err)
// 		return err
// 	}
// 	defer v.Body.Close()
// 	content, err := ioutil.ReadAll(v.Body)
// 	if err != nil {
// 		fmt.Printf("Read http response failed! %v", err)
// 		return err
// 	}
// 	err = ioutil.WriteFile(pic, content, 0666)
// 	if err != nil {
// 		fmt.Printf("Save to file failed! %v", err)
// 		return err
// 	}
// 	return nil
// }

// type test1 struct {
// 	testq int
// }

// func testg(testtt test1) string {
// 	if testtt.testq != 1 {
// 		return "hoow1"
// 	}
// 	return "ggyy"
// }

// func test001(num int) string {
// 	if num != 0 {
// 		return "num + 6"
// 	}
// 	return "num - 1"
// }

// func main() {
// 	var foo string
// 	foo = "sss"
// 	fmt.Println(foo)
// 	num := 2
// 	fmt.Println(testg(num))
// 	fmt.Println(testg(test1{
// 		testq: 0,
// 	}))
// 	fmt.Println(errors.New("hallow"))
// 	errors.New("hallow")
// }

// func modify(foo []string) {
// 	foo = fmt.Append(foo, "c")
// 	fmt.Println("modify foo:", foo)
// }

// func addvalue1(foo []string) []string {
// 	foo = append(foo, "c")
// 	fmt.Println("addvalue foo :", foo)
// 	return foo
// }

// func addvalue(foo *[]string) {
// 	*foo = append(*foo, "c")
// 	fmt.Println("addvalue foo :", foo)
// }

// func main() {
// 	foo := []string{"a", "b"}
// 	fmt.Println("befor foo :", foo)
// 	// modify(foo)
// 	// foo = addvalue1(foo)
// 	// addvalue(&foo)
// 	fmt.Println("after foo :", foo)
// 	bar := foo[:2]
// 	fmt.Println("bar :", bar)
// 	s1 := append(bar, "c")
// 	s2 := append(bar, "d")
// 	s3 := append(bar, "e", "f")
// 	fmt.Println("foo :", foo)
// 	fmt.Println("s1 :", s1)
// 	fmt.Println("s2 :", s2)
// 	fmt.Println("s3 :", s3)
// }

// func checkvalue(val int) {
// 	switch val {
// 	// case 0:
// 	// 	fallthrough
// 	case 0:
// 		fmt.Println("check value is ", val)
// 	}
// }

// func main() {
// 	checkvalue(0)
// 	checkvalue(1)
// }

// type car struct {
// 	name   string
// 	carnum int
// }

// func main() {
// 	// foo := 1
// 	// foostr := "1"
// 	// fooarr := []string[]
// 	fmt.Println("main : 1")
// }

// func init() {
// 	fmt.Println("init : 1")
// }

// func init() {
// 	fmt.Println("init : 2")
// }

// type car struct {
// 	name   string
// 	carnum int
// }

// type email struct {
// 	from string
// 	to   string
// }

// func (e *email) From(s string) {
// 	// fmt.Println("1 address: %p\n", &c)
// 	e.from = s
// }

// func (e *email) To(s string) {
// 	//fmt.Println("2 address: %p\n", e)
// 	e.to = s
// }

// func (e *email) Send() {
// 	fmt.Printf("From: %s ,To: %s\n", e.from, e.to)
// }

// func main() {
// 	// c := &car{}
// 	// fmt.Println("c.name :", c.name)
// 	// c.SetName01("bar")
// 	// fmt.Println("c.name :", c.name)
// 	// c.SetName02("foo")
// 	// fmt.Println("c.name :", c.name)
// 	// c := &email{}
// 	for i := 0; i < 11; i++ {
// 		go func(i int) {
// 			c := &email{}
// 			c.From(fmt.Sprintf("test%02d@gmail.com", i))
// 			c.To(fmt.Sprintf("test%02d@gmail.com", i+1))
// 			c.Send()
// 		}(i)
// 		time.Sleep(1 * time.Second)
// 		// c.From(fmt.Sprintf("test%02d@gmail.com", i))
// 		// c.To(fmt.Sprintf("test%02d@gmail.com", i+1))
// 		// c.Send()
// 	}
// }
