package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // APIを叩いて結果を格納する

    // 実行確認
    fmt.Printf("実行完了！！\n")

    // ファイル生成
    fp, err := os.Create("sitemap.xml")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer fp.Close()
    fp.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
        `<urlset xmlns="https://www.sitemaps.org/schemas/sitemap/0.9">` + "\n" +
        "<url>\n" +
        "<loc>https://www.example.com/</loc>\n" +
        "<lastmod>2020-04-01</lastmod>\n" +
        "<priority>0.8</priority>\n" +
        "</url>\n" +
        "</urlset>")

    // ファイル移動する
    err = os.Rename("sitemap.xml", "public/sitemap.xml")
    if err != nil {
        log.Fatal(err)
    }
}
