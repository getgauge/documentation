# Taking Custom Screenshots
* By default gauge captures the display screen on failure it this feature has been enabled.

* If you need to take CustomScreenshots (using webdriver for example) because you need only a part of the screen captured, this can be done by **implementing** the **ICustomScreenshotGrabber** interface;

Just Implement the interface in a class and gauge will use it to take capture screenshots.

````
Note: If multiple ICustomScreenshotGrabber implementations are found in classpath then gauge will pick one randomly to capture the screen.
````

#### Example - using webdriver

````java
public class CustomScreenshotFactory implements ICustomScreenshotGrabber {

    // Return a screenshot byte array
    public byte[] takeScreenshot() {
        WebDriver driver = DriverFactory.getDriver();
        return ((TakesScreenshot) driver).getScreenshotAs(OutputType.BYTES);
    }
}
````
