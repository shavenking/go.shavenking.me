1. Install Go
2. Install git, supervisor
3. `go get github.com/shavenking/go.shavenking.me`
4. `supervisorctl reread`
5. `supervisorctl update`

## Supervisor Configuration

```
[program:go_shavenking_me]
command=/home/shavenking/go/bin/go.shavenking.me
directory=/home/shavenking/go/src/github.com/shavenking/go.shavenking.me
autostart=true
autorestart=true
stderr_logfile=/var/log/go_shavenking_me.err.log
stdout_logfile=/var/log/go_shavenking_me.out.log
```
