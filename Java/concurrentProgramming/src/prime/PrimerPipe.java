package prime;

public class PrimerPipe extends Thread {
	private int prime;
	private PrimerPipe next;
	private int buffer = -1;
	public static long startTime = 0;

	public PrimerPipe(int prime) {
		super("Prime " + prime);
		this.prime = prime;
		this.start();
	}

	public void run() {
		int testNumber = 0;
		while (true) {
			testNumber = receive();
			if (testNumber == 0) {
				if (next != null) {
					next.send(testNumber);
				} else {
					long finishTime = System.nanoTime();
					System.out.format("0\t%d\t%,d\t%,d ns/op\n", PrimerPipeMain.maxInt, prime,
							(finishTime - startTime));
				}
				break;
			}
			if (testNumber % prime != 0) {
				if (next != null) {
					next.send(testNumber);
				} else {
					next = new PrimerPipe(testNumber);
				}
			}
		}
	}

	synchronized void send(int i) {
		try {
			while (buffer >= 0) {
				wait();
			}
			buffer = i;
			notify();
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
	}

	synchronized int receive() {
		int result = 0;
		try {
			while ((result = buffer) < 0) {
				wait();
			}
			buffer = -1;
			notify();
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
		return result;
	}
}