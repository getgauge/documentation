# Taking Custom Screenshots

* By default gauge captures the display screen on failure it this feature has been enabled.

* If you need to take CustomScreenshots (using webdriver for example) because you need only a part of the screen captured, this can be done by **implementing** the `ICustomScreenshotGrabber` (`IScreenGrabber` in C#) interface;

> Note: If multiple custom ScreenGrabber implementations are found in classpath then gauge will pick one randomly to capture the screen. 
This is because Gauge selects the first ScreenGrabber it finds, which in turn depends on the order of scanning of the libraries.


{% codetabs name="Java", type="java" -%}
// Using Webdriver
public class CustomScreenGrabber implements ICustomScreenshotGrabber {

    // Return a screenshot byte array
    public byte[] takeScreenshot() {
        WebDriver driver = DriverFactory.getDriver();
        return ((TakesScreenshot) driver).getScreenshotAs(OutputType.BYTES);
    }
}
{%- language name="C#", type="csharp" -%}
//Using Webdriver
public class CustomScreenGrabber : IScreenGrabber {

    // Return a screenshot byte array
    public byte[] TakeScreenshot() {
        var driver = DriverFactory.getDriver();
        return ((ITakesScreenshot) driver).GetScreenshot().AsByteArray;
    }
}
{%- language name="Ruby", type="ruby" -%}
# Using Webdriver
Gauge.configure do |config|
    # Return a screenshot byte array
    config.screengrabber =  -> {
        driver.save_screenshot('/tmp/screenshot.png')
        return File.binread("/tmp/screenshot.png")
    }
end
{%- endcodetabs %}