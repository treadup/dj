# dj
dj is a small command line utility for executing django manage.py
commands. Unlike manage.py you can execute dj commands from anywhere
in the project.

The dj command will map the django command

    dj <command>

to

    python manage.py <command>

## Install
To install dj use the following command.

    go install github.com/treadup/dj@latest
