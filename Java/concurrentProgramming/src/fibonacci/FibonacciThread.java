package fibonacci;

/**
 * @author Paulo This class calculates the fibonacci number of a value "n",
 *         where fibonacci(n) = fibonacci(n-1) + fibonacci(n-2). For every call
 *         of fibonacci(n) one new thread is started. For example: When the
 *         first thread starts to calculate fibonacci(3), this thread will open
 *         two others threads in order to calc fibonacci(2) and fibonacci(1) and
 *         wait for their return, so fibonacci(2) will also start other two
 *         threads to calculate fibonacci(1) again and fibonacci(0), after all
 *         threads are done the return value is assigned to the result value of
 *         this thread.
 */
public class FibonacciThread extends Thread {

	public int n = 0;
	public int result = 0;
	public FibonacciThread subThread_1;
	public FibonacciThread subThread_2;

	public FibonacciThread(int n) {
		this.n = n;
		this.start();
	}

	public void run() {
		if (n < 2) {
			result = n;
			return;
		}
		subThread_1 = new FibonacciThread(n - 1);
		subThread_2 = new FibonacciThread(n - 2);
		try {
			subThread_1.join();
			subThread_2.join();
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
		result = subThread_1.result + subThread_2.result;
	}
}