# golang-forever
Used to run a script forever. On script failure it will restart the run automatically.
## Usage
1. Run a script forever
```bash
$ go run main.go start  command filename 
```
### Example
```bash
$ go run main.go start node app.js
2019/01/28 23:05:55 Waiting for command to finish...
2019/01/28 23:05:55 Process id is 3390
2019/01/28 23:06:03 Command finished with error, now restarting: signal: terminated
2019/01/28 23:06:03 Waiting for command to finish...
2019/01/28 23:06:03 Process id is 3392
```