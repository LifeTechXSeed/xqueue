## XQueue 

- simple job queue write by golang

### TODO

- Main milestones:
  - Arbiter 
    - [ ] get job state
    - [ ] worker health check
    - [ ] enqueue instance job
    - [ ] kill job when it running
    - [ ] re-assign jobs when it's running worker dead

  - Scheduler
    - [ ] fetch nearby scheduled job
    - [ ] enqueue job when it's time

  - Worker agent
    - [ ] send current status to arbiter
    - [ ] dequeue job
    - [ ] run job
    - [ ] limit concurrent job

- Optional
  - [ ] job pipeline
    - [ ] running job step by step
    - [ ] global variable through pipeline
    - [ ] stop pipeline instantly 1 job error
  - [ ] specify worker for job
  - [ ] import/export job define 