package test

import (
	"CodeLab/designPatterns/Singleton/cache"
	"testing"
	"time"
)

const (
	redisURl          = "redis://localhost:6379"
	numWorkers        = 10                    // 并发工作者的数量
	numOperations     = 100                   // 每个工作者执行的操作次数
	operationInterval = 10 * time.Millisecond // 每次操作之间的间隔
)

func TestCache(t *testing.T) {
	cli1, err := cache.NewCache(redisURl)
	if err != nil {
		t.Fatalf("failed create redis: %v", err)
	}
	cli2, err := cache.NewCache(redisURl)
	if err != nil {
		t.Fatalf("failed create redis: %v", err)
	}

	if cli1 != cli2 {
		t.Fatalf("t1 != t2")
	}

}

//var wg sync.WaitGroup
//for i := 0; i < numWorkers; i++ {
//	wg.Add(1)
//	go func(workerID int) {
//		defer wg.Done()
//
//		// 每个工作者进行一定数量的Get和Set操作
//		for j := 1; j <= numOperations; j++ {
//			key := fmt.Sprintf("key%d-%d", workerID, j)
//			value := fmt.Sprintf("value%d-%d", workerID, j)
//
//			// 执行Set操作
//			err := cli.Set(context.Background(), key, value)
//			if err != nil {
//				t.Fatalf("Error setting key %s: %v\n", key, err)
//				continue
//			}
//
//			// 等待一定时间间隔
//			time.Sleep(operationInterval)
//
//			// 执行Get操作
//			result, err := cli.Get(context.Background(), key)
//			if err != nil {
//				t.Fatalf("Error getting key %s: %v\n", key, err)
//				continue
//			}
//
//			// 打印获取到的结果
//			fmt.Printf("Worker %d: Got key %s: %s\n", workerID, key, result)
//
//			// 等待一定时间间隔
//			time.Sleep(operationInterval)
//		}
//	}(i)
//}
//wg.Wait()
