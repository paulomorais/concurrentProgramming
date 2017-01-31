package matrix;

public class MatrixMultiplication extends Thread {

	private int[][] matrixA, matrixB, matrixC;
	private int indexNum, indexEnd;

	public MatrixMultiplication(int[][] matrixA, int[][] matrixB, int[][] matrixC, int indexNum, int indexEnd) {
		this.matrixA = matrixA;
		this.matrixB = matrixB;
		this.matrixC = matrixC;
		this.indexNum = indexNum;
		this.indexEnd = indexEnd;

	}

	public void run() {
		// Loops through the rows assigned to this thread
		for (int i = indexNum; i < indexEnd; i++) {

			// Loops through all columns of the matrix
			for (int j = 0; j < matrixC.length; j++) {

				// Loops through all indices of matrix C assigned to this thread
				for (int k = 0; k < matrixC.length; k++) {
					matrixC[i][j] += matrixA[i][k] * matrixB[k][j];
				}
			}
		}
	}
}
