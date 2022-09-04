package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func main() {
  type placeholder [5]string
  zero := placeholder{
    "000",
    "0 0",
    "0 0",
    "0 0",
    "000",
  }

  one := placeholder{
    "00 ",
    " 0 ",
    " 0 ",
    " 0 ",
    "000",
  }

  two := placeholder{
    "000",
    "  0",
    "000",
    "0  ",
    "000",
  }

  three := placeholder{
    "000",
    "  0",
    "000",
    "  0",
    "000",
  }
  four := placeholder{
    "0 0",
    "0 0",
    "000",
    "  0",
    "  0",
  }
  five := placeholder{
    "000",
    "0  ",
    "000",
    "  0",
    "000",
  }
  six := placeholder{
    "000",
    "0  ",
    "000",
    "0 0",
    "000",
  }
  seven := placeholder{
    "000",
    "  0",
    "  0",
    "  0",
    "  0",
  }
  eight := placeholder{
    "000",
    "0 0",
    "000",
    "0 0",
    "000",
  }
  nine := placeholder{
    "000",
    "0 0",
    "000",
    "  0",
    "000",
  }
  colon := placeholder{
    "   ",
    " 0 ",
    "   ",
    " 0 ",
    "   ",
  }
  digits := [...]placeholder{
    zero, one, two, three, four, five, six, seven, eight, nine,
  }
  CallClear()
  
  for {
    
    now := time.Now()
    hour, min, sec := now.Hour(), now.Minute(), now.Second()
  
    clock := [...]placeholder{
      digits[hour/10], digits[hour%10],
      colon,
      digits[min/10], digits[min%10],
      colon,
      digits[sec/10], digits[sec%10],
    }
  
    for line := range clock[0]{
      for digit := range clock{
        fmt.Print(clock[digit][line]+" ")
		fmt.Print(" ")
      }
      fmt.Println()
    }
    time.Sleep(time.Second)
    CallClear()
  }  
}