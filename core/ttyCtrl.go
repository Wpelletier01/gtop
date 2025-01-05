package core;


import (
    "os"
    "golang.org/x/sys/unix"
    "fmt"
)


func RawMode() (func(), error) {
    
    termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)

    if err != nil {
        return nil, fmt.Errorf("rawMode: error getting terminal flags: %w", err) 
    }

    copy := *termios

    // Turn off:
    //  - echoing
    //  - disable canonical mode 
    //  - ctrl-C and ctrl-Z signal
    //  - ctrl-S and ctrl-Q signal

    termios.Lflag = termios.Lflag &^ (unix.ECHO | unix.ICANON | unix.ISIG)
    termios.Lflag = termios.Lflag &^ (unix.ECHO | unix.ICANON | unix.ISIG | unix.IEXTEN)
	termios.Iflag = termios.Iflag &^ (unix.IXON | unix.ICRNL)
    termios.Oflag = termios.Oflag &^ (unix.OPOST)
    
     if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, termios); err != nil {
        return nil, fmt.Errorf("rawMode: error setting terminal flags: %w", err)
    }

    return func() {
		if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, &copy); err != nil {
			fmt.Fprintf(os.Stderr, "rawMode: error restoring originl console settings: %s", err)
		}
	}, nil


}



