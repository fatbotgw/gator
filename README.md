# gator
RSS aggreGATOR Project for boot.dev

# requirements
This is a Go program and uses Postgres for the database.

# config file
The repo contains an example config file.  You will need to place the config file
in your home directory and update the database address to reflect where your
database is located.  You can also update the current_user_name value, but it 
will be updated when you use the program's login command.

# commands
There are multiple commands, which are listed in the main.go file. Generally,
if there is a "handler" file the associated commands are in that file.

