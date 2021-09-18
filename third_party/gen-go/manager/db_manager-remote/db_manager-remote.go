// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/apache/thrift/lib/go/thrift"
        "manager"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  i64 CreateExperiment(ExperimentAddDto experiment)")
  fmt.Fprintln(os.Stderr, "  bool UpdateExperiment(ExperimentEditDto experiment)")
  fmt.Fprintln(os.Stderr, "  bool DeleteExperiment(i64 experimentId)")
  fmt.Fprintln(os.Stderr, "  i64 CreateTrial(i64 experimentId, i32 trialNo, string trialName)")
  fmt.Fprintln(os.Stderr, "  bool SuspendTrial(i64 id)")
  fmt.Fprintln(os.Stderr, "  bool ResumeTrial(i64 id)")
  fmt.Fprintln(os.Stderr, "  bool DeleteTrial(i64 id)")
  fmt.Fprintln(os.Stderr, "  bool UpdateTrialStatus(TrialEditDto trial)")
  fmt.Fprintln(os.Stderr, "  i64 CreateInfer(InferAddDto infer)")
  fmt.Fprintln(os.Stderr, "  bool SuspendInfer(i64 id)")
  fmt.Fprintln(os.Stderr, "  bool ResumeInfer(i64 id)")
  fmt.Fprintln(os.Stderr, "  bool DeleteInfer(i64 id)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := manager.NewDbManagerClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "CreateExperiment":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateExperiment requires 1 args")
      flag.Usage()
    }
    arg26 := flag.Arg(1)
    mbTrans27 := thrift.NewTMemoryBufferLen(len(arg26))
    defer mbTrans27.Close()
    _, err28 := mbTrans27.WriteString(arg26)
    if err28 != nil {
      Usage()
      return
    }
    factory29 := thrift.NewTJSONProtocolFactory()
    jsProt30 := factory29.GetProtocol(mbTrans27)
    argvalue0 := manager.NewExperimentAddDto()
    err31 := argvalue0.Read(jsProt30)
    if err31 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CreateExperiment(context.Background(), value0))
    fmt.Print("\n")
    break
  case "UpdateExperiment":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UpdateExperiment requires 1 args")
      flag.Usage()
    }
    arg32 := flag.Arg(1)
    mbTrans33 := thrift.NewTMemoryBufferLen(len(arg32))
    defer mbTrans33.Close()
    _, err34 := mbTrans33.WriteString(arg32)
    if err34 != nil {
      Usage()
      return
    }
    factory35 := thrift.NewTJSONProtocolFactory()
    jsProt36 := factory35.GetProtocol(mbTrans33)
    argvalue0 := manager.NewExperimentEditDto()
    err37 := argvalue0.Read(jsProt36)
    if err37 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.UpdateExperiment(context.Background(), value0))
    fmt.Print("\n")
    break
  case "DeleteExperiment":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "DeleteExperiment requires 1 args")
      flag.Usage()
    }
    argvalue0, err38 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err38 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.DeleteExperiment(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CreateTrial":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CreateTrial requires 3 args")
      flag.Usage()
    }
    argvalue0, err39 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err39 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err40 := (strconv.Atoi(flag.Arg(2)))
    if err40 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.CreateTrial(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "SuspendTrial":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SuspendTrial requires 1 args")
      flag.Usage()
    }
    argvalue0, err42 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err42 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SuspendTrial(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ResumeTrial":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ResumeTrial requires 1 args")
      flag.Usage()
    }
    argvalue0, err43 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err43 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ResumeTrial(context.Background(), value0))
    fmt.Print("\n")
    break
  case "DeleteTrial":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "DeleteTrial requires 1 args")
      flag.Usage()
    }
    argvalue0, err44 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err44 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.DeleteTrial(context.Background(), value0))
    fmt.Print("\n")
    break
  case "UpdateTrialStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UpdateTrialStatus requires 1 args")
      flag.Usage()
    }
    arg45 := flag.Arg(1)
    mbTrans46 := thrift.NewTMemoryBufferLen(len(arg45))
    defer mbTrans46.Close()
    _, err47 := mbTrans46.WriteString(arg45)
    if err47 != nil {
      Usage()
      return
    }
    factory48 := thrift.NewTJSONProtocolFactory()
    jsProt49 := factory48.GetProtocol(mbTrans46)
    argvalue0 := manager.NewTrialEditDto()
    err50 := argvalue0.Read(jsProt49)
    if err50 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.UpdateTrialStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CreateInfer":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateInfer requires 1 args")
      flag.Usage()
    }
    arg51 := flag.Arg(1)
    mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
    defer mbTrans52.Close()
    _, err53 := mbTrans52.WriteString(arg51)
    if err53 != nil {
      Usage()
      return
    }
    factory54 := thrift.NewTJSONProtocolFactory()
    jsProt55 := factory54.GetProtocol(mbTrans52)
    argvalue0 := manager.NewInferAddDto()
    err56 := argvalue0.Read(jsProt55)
    if err56 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CreateInfer(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SuspendInfer":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SuspendInfer requires 1 args")
      flag.Usage()
    }
    argvalue0, err57 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err57 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SuspendInfer(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ResumeInfer":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ResumeInfer requires 1 args")
      flag.Usage()
    }
    argvalue0, err58 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err58 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ResumeInfer(context.Background(), value0))
    fmt.Print("\n")
    break
  case "DeleteInfer":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "DeleteInfer requires 1 args")
      flag.Usage()
    }
    argvalue0, err59 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err59 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.DeleteInfer(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
