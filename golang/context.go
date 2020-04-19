// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// 	// "os"
// 	// "net/http"
// )

// var key string = "name"

// // func main() {
// // 	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		ctx := r.Context()
// // 		fmt.Fprint(os.Stdout, "processing request\n")

// // 		select {
// // 			case <-time.After(2 * time.Second):
// // 				w.Write([]byte("request processed"))
// // 			case <-ctx.Done():
// // 				fmt.Fprint(os.Stderr, "request cancelled\n")
// // 		}
// // 	}))
// // }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	//附加值
// 	valueCtx := context.WithValue(ctx, key, "【监控1】")
// 	go watch(valueCtx)
// 	// go children(ctx)
// 	// go watchChildren(ctx)
// 	// time.Sleep(10 * time.Second)
// 	// fmt.Println("可以了，通知监控停止")
// 	// for {
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 		time.Sleep(1 * time.Second)
// 	}
// 	// }
// 	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
// 	cancel()

// }

// func watch(ctx context.Context) {
// 	go children(ctx)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			//取出值
// 			fmt.Println(ctx.Value(key), "监控退出，停止了...")
// 			return
// 		default:
// 			//取出值
// 			fmt.Println(ctx.Value(key), "goroutine监控中...")
// 			time.Sleep(2 * time.Second)
// 		}
// 	}
// }

// func children(ctx context.Context) {
// 	ctxChildren, cancel := context.WithTimeout(ctx, 10*time.Second)

// 	childrenValue := context.WithValue(ctxChildren, key, "【监控2】")
// 	go watchChildren(childrenValue)

// 	select {
// 	case <-ctxChildren.Done():
// 		fmt.Println("children context finish 2")
// 	}
// 	defer cancel()

// }

// func watchChildren(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Value(key), "children context finish 1")
// 			return
// 		case <-time.After(2 * time.Second):
// 			fmt.Println(ctx.Value(key), "children 監控中...")
// 		}
// 	}
// }

// package main

// import (
//     "context"
//     "crypto/md5"
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "sync"
//     "time"
// )

// type favContextKey string

// func main() {
//     wg := &sync.WaitGroup{}
//     values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
//     ctx, cancel := context.WithCancel(context.Background())

//     for _, url := range values {
//         wg.Add(1)
//         subCtx := context.WithValue(ctx, favContextKey("url"), url)
//         go reqURL(subCtx, wg)
//     }

//     go func() {
//         time.Sleep(time.Second * 3)
//         cancel()
//     }()

//     wg.Wait()
//     fmt.Println("exit main goroutine")
// }

// func reqURL(ctx context.Context, wg *sync.WaitGroup) {
//     defer wg.Done()
//     url, _ := ctx.Value(favContextKey("url")).(string)
//     for {
//         select {
//         case <-ctx.Done():
//             fmt.Printf("stop getting url:%s\n", url)
//             return
//         default:
//             r, err := http.Get(url)
//             if r.StatusCode == http.StatusOK && err == nil {
//                 body, _ := ioutil.ReadAll(r.Body)
//                 subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
//                 wg.Add(1)
//                 go showResp(subCtx, wg)
//             }
//             r.Body.Close()
//             //启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
//             // doSometing()
//             time.Sleep(time.Second * 1)
//         }
//     }
// }

// func showResp(ctx context.Context, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for {
//         select {
//         case <-ctx.Done():
//             fmt.Println("stop showing resp")
//             return
//         default:
//             //子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
//             fmt.Println("printing ", ctx.Value(favContextKey("resp")))
//             time.Sleep(time.Second * 1)
//         }
//     }
// }

package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		time.Sleep(1 * time.Second)
	}
	
	cancel()

}

func watch(ctx context.Context) {
	go children(ctx)
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			// time.Sleep(2 * time.Second)
		}
	}
}

func children(ctx context.Context) {
	ctxChildren, cancel := context.WithTimeout(ctx, 1*time.Second)

	childrenValue := context.WithValue(ctxChildren, key, "【监控2】")
	go watchChildren(childrenValue)

	select {
	case <-ctxChildren.Done():
		fmt.Println("children context finish 2")
	}
	defer cancel()

}

func watchChildren(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "children context finish 1")
			return
		default:
			fmt.Println(ctx.Value(key), "children 監控中...")
		}
	}
}
