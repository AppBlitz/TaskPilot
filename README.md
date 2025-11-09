
# Task tracker 

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

# Task Pilot

It is an application that creates tasks using command lines, where the main features are the ability to update, create, and delete the different task lists you have.
Tasks are saved in a json file on the computer where they are executed.


# How build 

Clone the repository and run the following command:

```git
https://github.com/AppBlitz/TaskPilot.git
cd TaskPilot
```


Run the following command to build and run the project:

```bash
go build -o task-tracker
./task-tracker --help # To see the list of available commands

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```


https://github.com/charmbracelet/lipgloss
