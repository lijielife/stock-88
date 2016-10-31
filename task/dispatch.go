package task

import "time"

type (
	//Task 任务执行接口
	Task interface {
		Run()
		Interval() time.Duration
	}
)

var (
	//任务列表
	tasks = []Task{
		NewBasicsTask(),
	}
)

//Start 启动任务
func Start() {
	for _, task := range tasks {
		go func(task Task) {
			t := time.NewTimer(task.Interval())
			for {
				go task.Run()
				<-t.C
				t.Reset(task.Interval())
			}
		}(task)
	}

}
