asciinema-edit
--------------

Auxiliary tools for dealing with [asciinema](https://asciinema.org/docs/getting-started) casts.


NAME:
   asciinema-edit - edit recorded asciinema casts

INSTALL
   Using `go`, fetch the latest from `master`

     go get -u -v github.com/cirocosta/asciinema-edit

   Retrieving from GitHub releases

     curl -SOL ...

USAGE:
   asciinema-edit [global options] command [command options] [arguments...]

VERSION:
   dev

DESCRIPTION:
   asciinema-edit provides missing features from the "asciinema" tool
   when it comes to editing a cast that has already been recorded.

COMMANDS:
     cut  Removes a certain range of time frames.

   If no file name is specified as a positional argument a cast is
   expected to be served via stdin.

   Once the transformation has been performed, the resulting cast is
   either written to a file specified in the '--out' flag or to stdout
   (default).

EXAMPLES
   Remove frames from 12.2s to 16.3s from the cast passed in the commands
   stdin.

     cat 1234.cast | \
       asciinema-edit cut \
         --from=12.2 --to=15.3

   Remove the exact frame at timestamp 12.2 from the cast file named
   1234.cast.

     asciinema-edit cut \
       --from=12.2 --to=12.2 \
       1234.cast
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
