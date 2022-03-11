This repository is a work in progress

# Go-Service
Use this repository as a template for creating a variety of Go Services.

## Background
I find myself using the same framework anytime I need to create a Go Service.  To save myself -- and potentially others -- some time, I'm putting the code here.

These services exist to be put into a package or a module so that it can perform some action regularly.  In particular, 'cron' tasks (like synchronization) use a lot of the same code project to project.

## Future Changes
In the future, I hope to add/update:
- A Wrapper for Viper or a similar library to:
    - Load Config from a file
    - Load Overrides from Env Vars
    - Combined multiple "Configs" into a single larger config
- A Generator, if possible for Configs and for some service features / frameworks

