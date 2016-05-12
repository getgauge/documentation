---
title: Custom Screenshots
--- 

# Taking Custom Screenshots

* By default gauge captures the display screen on failure it this feature has been enabled.

* If you need to take CustomScreenshots (using webdriver for example) because you need only a part of the screen captured, this can be done by **implementing** the `ICustomScreenshotGrabber` (`IScreenGrabber` in C#) interface;

> Note: If multiple custom ScreenGrabber implementations are found in classpath then gauge will pick one randomly to capture the screen. 
This is because Gauge selects the first ScreenGrabber it finds, which in turn depends on the order of scanning of the libraries.

## Java

#### Example - using webdriver

````java
public class CustomScreenGrabber implements ICustomScreenshotGrabber {

    // Return a screenshot byte array
    public byte[] takeScreenshot() {
        WebDriver driver = DriverFactory.getDriver();
        return ((TakesScreenshot) driver).getScreenshotAs(OutputType.BYTES);
    }
}
````

## C# #

#### Example - using WebDriver

````csharp
public class CustomScreenGrabber : IScreenGrabber {

    // Return a screenshot byte array
    public byte[] TakeScreenshot() {
        var driver = DriverFactory.getDriver();
        return ((ITakesScreenshot) driver).GetScreenshot().AsByteArray;
    }
}
````

## Ruby

> Note: This feature is not implemented in Gauge Ruby. Please refer [this issue](https://github.com/getgauge/gauge-ruby/issues/13) for updates.
