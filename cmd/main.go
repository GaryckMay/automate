package main

import (
    "fmt"
    "github.com/tebeka/selenium"
    "github.com/tebeka/selenium/chrome"
    "log"
)

func main() {
    // Run Chrome browser
    ops := []selenium.ServiceOption{
    }
    //service, err := selenium.NewChromeDriverService("./driver/chromedriver126.exe", 4444,ops...)
    service, err := selenium.NewChromeDriverService("./driver/yandexdriver.exe", 4444,ops...)
    if err != nil {
        log.Fatal("Error:", err)
    }
    defer service.Stop()

    caps := selenium.Capabilities{}
    caps.AddChrome(chrome.Capabilities{Args: []string{
        "window-size=1920x1080",
        "--no-sandbox",
        "--disable-dev-shm-usage",
        "disable-gpu",
        // "--headless",  // comment out this line to see the browser
    }})

    driver, err := selenium.NewRemote(caps, "")
    if err != nil {
        log.Fatal("Error:", err)
    }

    err = driver.Get("https://www.google.com")
    if err != nil {
        log.Fatal("Error:", err)
    }

    html, err := driver.PageSource()
    if err != nil {
    log.Fatal("Error:", err)
    }
    fmt.Println(html)
}