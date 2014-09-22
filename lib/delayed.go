package lib

func ProcessDelayedJob(job func(interface{}), args interface{}) {
  go job(args)
}
