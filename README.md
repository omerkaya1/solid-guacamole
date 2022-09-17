# Solid-Guacamole

## Disclaimer

This project is NOT intended as a way to abuse the service,
NOR it is intended to rob other people of their right to the same service.

# Installation
`go install github.com/omerkaya1/solid-guacamole@latest`
or
- download the [file](./solid-guacamole) (for mac)
- download build artifacts from [actions page](https://github.com/omerkaya1/solid-guacamole/actions)


## Usage (currently works for mac and linux)

The following steps are performed through terminal/shell of your computer.

1. Run the programme on the background: `solid-guacamole &`
NOTE: If the file was downloaded, you need to make it "runnable": `chmod u=x path/to/solid-guacamole && ./solid-guacamole &`
2. When the free slot is found, the programme will notify you with 
screen blinking and a sound once a second; to stop that, send SIGUSR1 to the process: `kill -s SIGUSR1 $(pgrep solid-guacamole)`
3. Once the booking is complete, kill the programme: `kill -s SIGKILL $(pgrep solid-guacamole)`