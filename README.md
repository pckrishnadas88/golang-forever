# golang-forever
Used to run a script forever. On script failure it will restart the run automatically.
## Usage
1. Run a script forever
```bash
$ go run main.go start  command filename 
```
## Example
```bash
$ go run main.go start node app.js
```