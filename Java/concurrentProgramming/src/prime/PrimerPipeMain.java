package prime;

public class PrimerPipeMain {
	
	public static int maxInt = 10_000;
	
	public static void main(String[] args) {
		PrimerPipe.startTime = System.nanoTime();
		PrimerPipe first = new PrimerPipe(2);
		for (int n = 3; n <= maxInt; n++) {
			first.send(n);
		}
		first.send(0);
	}
}
/*
repeat	maxInt	avgTime/operation
500	1000	13 ms/op
200	2000	32 ms/op
100	3000	53 ms/op
100	4000	80 ms/op
100	5000	105 ms/op
10	10000	264 ms/op
2	50000	5,113 ms/op
1	100000	19,818 ms/op
 */

