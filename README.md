# GTask

## Overview
Gtask is cli task manager.

## Build
```
go build main.go
```

## HowTo
### Confirm Task
```
% gtask p
-------------------------------------------------------------
|id |title                              |date           |c  |
-------------------------------------------------------------
|1  |Task                               |2017/02/05     |-  |
-------------------------------------------------------------
```

confirm all task(is completed)
```
% gtask print -c true
```


### Append Task
```
% gtask add -t [Task name] -d [in deadline days]
```

### Complete Task
```
% gtask finish -i [id]
```

### Update Task
```
% gtask update -i [id] -t [Task name] -d [in deadline days]
```

### Post Slack
```
% gtask post
```

### plans
- improve task view
- improve command interface

