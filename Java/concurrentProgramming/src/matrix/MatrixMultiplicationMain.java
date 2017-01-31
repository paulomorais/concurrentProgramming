package matrix;


import java.util.Random;

public class MatrixMultiplicationMain {
	
    private static MatrixMultiplication[] threads;
    private static int[][] MatrixA, MatrixB, MatrixC;
    
    static int[][] execution = {
    		{4, 1, 5000}, {4, 2, 5000}, {4, 4, 5000},
    		{8, 1, 5000}, {8, 2, 5000}, {8, 4, 5000}, {8, 8, 2000},
    		{16, 1, 5000}, {16, 2, 5000}, {16, 4, 5000}, {16, 8, 5000}, {16, 16, 2000},
    		{32, 1, 5000}, {32, 2, 5000}, {32, 4, 5000}, {32, 8, 5000}, {32, 16, 2000}, {32, 32, 1000},
    		{64,1,2000}, {64,2,2000}, {64,4,3000}, {64,8,3000}, {64,16,3000}, {64,32,3000}, {64,64,3000},
    		{128,1,500}, {128,2,500}, {128,4,500}, {128,8,500}, {128,16,500}, {128,32,500}, {128,64,500}, {128,128,500}, 
    		{256,1,100}, {256,2,100}, {256,4,100}, {256,8,100}, {256,16,100}, {256,32,100}, {256,64,100}, {256,128,100}, {256,256,100}, 
    		{512,1,10}, {512,2,10}, {512,4,10}, {512,8,10}, {512,16,10}, {512,32,10}, {512,64,10}, {512,128,10}, {512,256,10}, {512,512,10},
    		{1024,1,4}, {1024,2,4}, {1024,4,4}, {1024,8,4}, {1024,16,4}, {1024,32,4}, {1024,64,4}, {1024,128,4}, {1024,256,4}, {1024,512,4}, {1024,1024,4}
    	};
    
    public static void main(String[] args) throws InterruptedException{
    	for (int e = 0; e < execution.length; e++) {
   			if (execution[e][0] % execution[e][1] == 0) {
   				testUnit(execution[e][0], execution[e][1], execution[e][2]);
   			}
    	}
    }

    public static void testUnit(int size, int concurrentExecutions, int repeat) throws InterruptedException{
        long sumTime = 0;
        generateRandomMatrix(size);
        
        for(int repetition = 0; repetition < repeat; repetition++) {
            threads = new MatrixMultiplication[concurrentExecutions]; // Construction for thread array

			long start = System.currentTimeMillis();
            for(int conc = 0; conc < concurrentExecutions; conc++) { // Allocating Part of Matrix A and B
            	threads[conc] = new MatrixMultiplication(MatrixA, MatrixB, MatrixC, 
            			conc 		* (size / concurrentExecutions), 
            			(conc+1)	* (size / concurrentExecutions)); 
            }
            for(int conc = 0; conc < concurrentExecutions; conc++) {// Starting threads
            	threads[conc].start();
            }
            for(int conc = 0; conc < concurrentExecutions; conc++) {// Waiting for threads
           		threads[conc].join();
            }
            long stop = System.currentTimeMillis(); // Getting End Time
            sumTime += (stop - start);
            clearMatrixC(size);
        }
        System.out.format("%d\t%d\t%,d us/op\tN%dT%d\n", repeat, sumTime, (sumTime*1000)/repeat, size, concurrentExecutions);
        Thread.sleep(2000);
    }
    
    public static void clearMatrixC(int size){
    	for(int i = 0; i < size; i++)
            for(int j = 0; j < size; j++) 
                MatrixC[i][j] = 0;
    }
    
    public static void generateRandomMatrix(int size){
    	MatrixA = new int[size][size];
        MatrixB = new int[size][size];
        MatrixC = new int[size][size];
        
    	Random mRandom = new Random();
        for(int i = 0; i < size; i++) {
            for(int j = 0; j < size; j++) {
                MatrixA[i][j] = mRandom.nextInt(100);
                MatrixB[i][j] = mRandom.nextInt(100);
            }
        }
    }
}
/*
repeat	time/operation	N<matrixSize>T<threads>
2000	271,429 ns/op	N64T1
3000	187,323 ns/op	N64T2
3000	200,848 ns/op	N64T4
3000	310,995 ns/op	N64T8
3000	564,900 ns/op	N64T16
3000	1,101,974 ns/op	N64T32
3000	2,460,191 ns/op	N64T64
300	2,029,609 ns/op	N128T1
500	1,081,569 ns/op	N128T2
1000	886,907 ns/op	N128T4
500	861,405 ns/op	N128T8
500	1,132,291 ns/op	N128T16
500	1,639,115 ns/op	N128T32
500	2,877,150 ns/op	N128T64
20	21,541,661 ns/op	N256T1
50	10,823,350 ns/op	N256T2
50	9,065,639 ns/op	N256T4
50	7,071,313 ns/op	N256T8
50	7,462,526 ns/op	N256T16
100	8,013,056 ns/op	N256T32
50	8,552,126 ns/op	N256T64
2	243,743,267 ns/op	N512T1
3	127,704,586 ns/op	N512T2
5	97,749,952 ns/op	N512T4
5	74,992,800 ns/op	N512T8
5	78,268,137 ns/op	N512T16
5	73,098,276 ns/op	N512T32
5	73,665,838 ns/op	N512T64
1	6,218,319,912 ns/op	N1024T1
1	4,015,686,160 ns/op	N1024T2
1	2,321,049,116 ns/op	N1024T4
1	2,136,995,757 ns/op	N1024T8
1	2,118,328,536 ns/op	N1024T16
1	2,120,757,061 ns/op	N1024T32
1	2,140,047,857 ns/op	N1024T64



1	100	334,582 ns/op	64
1	100	1,834,742 ns/op	128
1	50	18 ms/op	256
1	20	216 ms/op	512
1	10	6,249 ms/op	1024
2	100	201,193 ns/op	64
2	100	990,783 ns/op	128
2	50	9 ms/op	256
2	20	111 ms/op	512
2	10	3,990 ms/op	1024
4	100	252,704 ns/op	64
4	100	746,852 ns/op	128
4	50	11 ms/op	256
4	20	98 ms/op	512
4	10	2,564 ms/op	1024
8	100	442,963 ns/op	64
8	100	1,030,104 ns/op	128
8	50	8 ms/op	256
8	20	90 ms/op	512
8	10	2,548 ms/op	1024
16	100	903,696 ns/op	64
16	100	1,425,798 ns/op	128
16	50	9 ms/op	256
16	20	90 ms/op	512
16	10	2,418 ms/op	1024
32	100	1,641,185 ns/op	64
32	100	2,082,050 ns/op	128
32	50	9 ms/op	256
32	20	142 ms/op	512
32	10	2,412 ms/op	1024
64	100	3,845,675 ns/op	64
64	100	3,842,846 ns/op	128
64	50	10 ms/op	256
64	20	91 ms/op	512
64	10	2,438 ms/op	1024
*/
