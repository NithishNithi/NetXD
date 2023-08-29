package main 
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)
type Job struct{
    id int 
    randomno int 
}
type Result struct{
    job Job
    digitsum int
}
var jobs=make(chan Job,10)
var results=make(chan Result,10)
func digits(num int) int{
  sum:=0
  no:=num
  for no!=0{
    digit:=num%10
    sum=sum+digit
    no/=10
  }
  return sum
}
func worker(wg*sync.WaitGroup){
    for job:= range jobs{
        output:=Result{job,digits(job.randomno)}
        results<-output
    }
    wg.Done()
}
func createWorkerpool(workerno int){
    var wg sync.WaitGroup
    for i:=0;i<workerno;i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    close(results)
}
func allocate(jobno int){
    for i:=0 ;i<jobno;i++ {
        randomno:=rand.Intn(999)
        job:=Job{i,randomno}
        jobs<-job


    }
    close(jobs)
}
func result(done chan bool){
    for result:=range results{
        fmt.Printf("Job id %d, input random no %d,  sum of digits %d\n",result.job.id,result.job.randomno,result.digitsum)
    }
}
func main(){
startTime:=time.Now()
jobno:=100
go allocate(jobno)
done:=make(chan bool)
go result(done)
workerno:=10
createWorkerpool(workerno)
<-done
endTime:=time.Now()
diff:=endTime.Sub(startTime)
fmt.Println("total time taken ",diff.Seconds(),"seconds")


}