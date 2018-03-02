package worker
import "sync"
type Worker interface{
	Task()
}

type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}
// 创建线程池
func New(maxGoroutines int)*Pool{
	p:= Pool{
		work:make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i:=0;i<maxGoroutines;i++{
		go func(){
			for w:=range p.work{
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}
//提交工作到线程池
func (p*Pool)Run(w Worker){
	p.work<-w
}
func (p*Pool)Shutdown(){
	close(p.work)
	p.wg.Wait()
}