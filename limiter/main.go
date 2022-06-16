package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// allow()
	// wait()

	reserve()
}

/*
	AllowN方法表示，截止在某一时刻，目前桶中数目是否至少为n个。如果条件满足，则从桶中消费n个token，同时返回true。反之不消费Token，返回false。
	使用场景：一般用在如果请求速率过快，直接拒绝请求的情况
*/
func allow() {
	// 初始化一个限速器，每秒产生10个令牌，桶的大小为100个
	// 初始化状态桶是满的
	var limiter = rate.NewLimiter(10, 100)
	for i := 0; i < 20; i++ {
		if limiter.AllowN(time.Now(), 25) {
			fmt.Printf("%03d Ok  %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			fmt.Printf("%03d Err %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

/*
	当使用Wait方法消费Token时，如果此时桶内Token数量不足(小于N)，那么Wait方法将会阻塞一段时间，直至Token满足条件。否则直接返回。
	可以看到Wait方法有一个context参数。我们可以设置context的Deadline或者Timeout，来决定此次Wait的最长时间
*/
func wait() {
	// 指定令牌桶大小为5，每秒补充3个令牌
	limiter := rate.NewLimiter(3, 5)

	// 指定超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))

		// 每次消费2个令牌
		err := limiter.WaitN(ctx, 2)
		if err != nil {
			fmt.Printf("timeout: %s \n", err.Error())
			return
		}
	}

	fmt.Println("main")
}

/*
  此方法有一点复杂，它返回的是一个*Reservation类型，后续操作主要针对的全是这个类型
  判断限制器是否能够在指定时间提供指定N个请求令牌。
  如果Reservation.OK()为true，则表示需要等待一段时间才可以提供，其中Reservation.Delay()返回需要的延时时间。
  如果Reservation.OK()为false,则Delay返回InfDuration, 此时不想等待的话，可以调用 Cancel()取消此次操作并归还使用的token
*/

func reserve() {
	// 指定令牌桶大小为5，每秒补充3个令牌
	limiter := rate.NewLimiter(3, 5)

	// 指定超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		reserve := limiter.Reserve()
		if !reserve.OK() {
			// 返回是异常的，不能正常使用
			fmt.Println("Not allowed to act! Did you remember to set lim.burst to be > 0 ?")
			return
		}
		delayD := reserve.Delay()
		fmt.Println("sleep delay ", delayD)
		time.Sleep(delayD)
		select {
		case <-ctx.Done():
			fmt.Println("timeout, quit")
			return
		default:
		}
		//TODO 业务逻辑
	}

	fmt.Println("main")
}
