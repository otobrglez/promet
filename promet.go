package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "fmt"
import (
  "runtime"
  "net/http"
  "io/ioutil"
)

//export WhatsMyName
func WhatsMyName() *C.char {
  s := "Oto Brglez"
  c_str := C.CString(s)
  defer C.free(unsafe.Pointer(c_str))
  return C.CString("Oto Brglez")
}

//export GoEnv
func GoEnv() *C.char {
  return C.CString(fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}

//export Events
func Events() (*C.char, *C.char) {
  events, error := getEncodedEvents()

  if error != nil {
    return C.CString(""), C.CString(error.Error())
  }

  return C.CString(string(events)), C.CString("")
}

func getEncodedEvents() ([]byte, error) {
  eventsUrl := "http://www.promet.si/rwproxy/RWProxy.ashx?method=get&remoteUrl=http%3A//promet/events_pp"

  request, err := http.NewRequest("GET", eventsUrl, nil)
  request.Header.Set("Content-Type", "application/json")

  if err != nil {
    return nil, err
  }

  response, err_2 := (&http.Client{}).Do(request)
  if err_2 != nil {
    return nil, err_2
  }

  data, err_3 := ioutil.ReadAll(response.Body)
  if err_3 != nil {
    return nil, err_3
  }

  return data, nil
}

func main() {
  fmt.Printf("promet on %s / %s\n", runtime.GOOS, runtime.GOARCH)

  events, _ := Events()
  fmt.Printf("%s", C.GoString(events))
}


