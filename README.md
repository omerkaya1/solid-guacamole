# Solid-Guacamole

## Disclaimer

This project is NOT intended as a way to abuse the service,
NOR it is intended to rob other people of their right to the same service.

# Installation
`go install github.com/omerkaya1/solid-guacamole@latest`

## Usage (currently works for mac and linux)

1. Run the programme on the background: `solid-guacamole &`
2. When the free slot is spotted, the programme will notify you with 
screen blinking and a sound once a second; to stop that, on linux/mac
send SIGUSR1 to the process: `kill -s SIGUSR1 $(pgrep solid-guacamole)`
3. Once the booking is complete, kill the programme: `kill -s SIGKILL $(pgrep solid-guacamole)`