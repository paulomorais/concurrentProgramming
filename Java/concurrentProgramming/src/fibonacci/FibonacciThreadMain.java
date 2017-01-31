package fibonacci;

public class FibonacciThreadMain {
	
	public static int result = 0;
	static int[] executions = {10_000, 10_000, 10_000, 10_000, 5_000, 5_000, 5_000, 2000, 
			2_000, 2_000, 500, 500, 200, 200, 200, 100, 100, 100, 20, 
			20, 20, 10, 5};
	
	private static FibonacciThread fiboThread = null;
	
	public static void main(String[] args) throws Exception {
		//Loop through all elements of of the array executions
		for (int n = 0; n < executions.length; n++) {
			
			//Starts to count the time
			long startTime = System.nanoTime();
			
			//Loop through the value of each element of of the array, this loop represents how many times the f(n) will be executed to get the average time
			for (int j = 0; j < executions[n]; j++) {
				result = run(n);
			}
			long finishTime = System.nanoTime();
			
			//Print: <'n'>	<#Executions>	<average time to run each calculation>	<fibonacci(n)>
			System.out.format("%d\t%,d\t%,d ns/op\tf(%d)=%d\n", n, executions[n], (finishTime - startTime)/executions[n], n, result);
		}
	}
	
	public static int run(int n) throws Exception {
		//Create Thread to calculate Fibonacci of 'n' and starts automatically
		fiboThread = new FibonacciThread(n);
		
		//Wait for thread to end
		fiboThread.join();
		
		//Return f(n) 
		return fiboThread.result;
	}
}

/*


*/