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
% gtask p -c true
```


### Append Task
```
% gtask in -t [Task name] -d [in deadline days]
```

### Complete Task
```
% gtask fi -i [id]
```

### Update Task
```
% gtask u -i [id] -t [Task name] -d [in deadline days]
```

### plans
- post to slack
- improve task view

