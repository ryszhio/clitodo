# CLI based TODO list
This cli tool is to track your TODO list, which provides following features for now:
- add
- remove
- complete
- list

## How to use ?

### add
add sub command is used to simply add a new todo task.
**Usage**:
```
./clitodo add task-name
Example: ./clitodo add Complete learning cobra
```
### remove
remove sub command is used to remove any existing todo task.
**Usage**
```
./clitodo remove task-id
Example: ./clitodo remove 1 # Removes todo task-id 1
```

### list
list sub command will display current task as well as completed task.
**Usage**
```
./clitodo list
Examples:
    ./clitodo list # Will display current task
    ./clitodo list --completed # Will display completed task
```

### complete
complete sub command will simply set a todo task as completed.
**Usage**
```
./clitodo complete task-id
Example: ./clitodo complete 1 # Will set todo task-id 1 as completed
```

## Future Plans
- Make this app an interactive CLI

## Credits
- cobra package: https://github.com/spf13/cobra