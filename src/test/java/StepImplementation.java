import com.thoughtworks.gauge.*;

public class StepImplementation {
    @BeforeSuite
    public void beforeSuite() {
        System.out.println("BeforeSuite");
    }

    @AfterSuite
    public void afterSuite() {
        System.out.println("AfterSuite");
    }
    @BeforeSpec
    public void beforeSpec() {
        System.out.println("BeforeSpec");
    }
    @AfterSpec
    public void afterSpec() {
        System.out.println("AfterSpec");
    }
    @BeforeScenario
    public void beforeScenario() {
        System.out.println("BeforeScenario");
    }
    @AfterScenario
    public void afterScenario() {
        System.out.println("AfterScenario");
    }


    @Step("noop")
	public void noop(){
		System.out.println("ThreadName: " + Thread.currentThread().getName());
	}
}
