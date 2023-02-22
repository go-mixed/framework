package queue

//go:generate mockery --name=Queue
type Queue interface {
	Worker(args *Args) Worker
	// Register Register jobs
	Register(jobs []Job)
	// GetJobs Get all jobs
	GetJobs() []Job
	// Job Add a job to queue
	Job(job Job, args []Arg) Task
	// Chain Creates a chain of jobs to be processed one by one, passing
	Chain(jobs []Jobs) Task
}

type Worker interface {
	Run() error
}

type Args struct {
	// Specify connection
	Connection string
	// Specify queue
	Queue string
	// Concurrent num
	Concurrent int
}

type Arg struct {
	Type  string
	Value any
}
