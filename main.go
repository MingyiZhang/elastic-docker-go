package main

import (
  "context"
  "fmt"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/olivere/elastic/v7"
)

func main() {
  url := os.Getenv("ELASTIC_URL")
  client, err := elastic.NewClient(
    elastic.SetURL(url),
  )
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Connected to %s", url)

  errc := make(chan error)

  go func() {
    err := showNodes(client)
    if err != nil {
      log.Printf("nodes info failed: %v", err)
    }
    t := time.NewTicker(10 * time.Second)
    defer t.Stop()
    for {
      select {
      case <-t.C:
        err := showNodes(client)
        if err != nil {
          log.Printf("nodes info failed: %v", err)
        }
      }
    }
  }()

  go func() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
    log.Printf("existing with signal %v", fmt.Sprint(<-c))
    errc <- nil
  }()

  if err := <-errc; err != nil {
    os.Exit(1)
  }
}

func showNodes(client *elastic.Client) error {
  ctx := context.Background()
  info, err := client.NodesInfo().Do(ctx)
  if err != nil {
    return err
  }
  log.Printf("Cluster %q with %d node(s)", info.ClusterName, len(info.Nodes))
  for id, node := range info.Nodes {
    log.Printf("- Node %s with IP %s", id, node.IP)
  }
  return nil
}
